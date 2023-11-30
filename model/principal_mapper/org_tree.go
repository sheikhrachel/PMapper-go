package principal_mapper

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/sheikhrachel/PMapper-go/api_common/errUtils"
)

// OrgTree represents an AWS organization, sorted into a hierarchy of AWS accounts
type OrgTree struct {
	// ID is the ID of the AWS organization
	ID string `json:"org_id"`
	// ManagementAccountID is the ID of the management account
	ManagementAccountID string `json:"management_account_id"`
	// RootOUs is a list of root OUs in the organization
	RootOUs []OrgNode `json:"root_ous"`
	// AllSCPs is a list of all SCPs in the organization
	AllSCPs []Policy `json:"all_scps"`
	// Accounts is a list of all accounts in the organization
	Accounts []string `json:"accounts"`
	// EdgeList is a list of all edges in the organization
	EdgeList []Edge `json:"edge_list"`
	// Metadata is a map of metadata about the organization
	Metadata map[string]string `json:"metadata"`
}

// NewOrgTree creates a new organization tree object.
func NewOrgTree(
	id, managementAccountID string,
	rootOUs []OrgNode,
	allSCPs []Policy,
	accounts []string,
	edgeList []Edge,
	metadata map[string]string,
) (orgTree *OrgTree, err error) {
	orgTree = &OrgTree{
		ID:                  id,
		ManagementAccountID: managementAccountID,
		RootOUs:             rootOUs,
		AllSCPs:             allSCPs,
		Accounts:            accounts,
		EdgeList:            edgeList,
		Metadata:            metadata,
	}
	if orgTree.Metadata == nil {
		orgTree.Metadata = make(map[string]string)
	}
	return orgTree, nil
}

/*
Map returns a map representation of the org tree.

	Used for serialization to disk. SCPs and metadata are excluded
	since `SaveOrgToDisk` does those in a separate file.
*/
func (o *OrgTree) Map() map[string]interface{} {
	return map[string]interface{}{
		"org_id":                o.ID,
		"management_account_id": o.ManagementAccountID,
		"root_ous":              o.RootOUs,
		"accounts":              o.Accounts,
		"edge_list":             o.EdgeList,
	}
}

// SaveOrganizationToDisk stores the Organization object as JSON data to disk in a standard layout.
func (o *OrgTree) SaveOrganizationToDisk(path string) (err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, 0700); errUtils.HandleError(err) {
			return err
		}
	}
	var (
		metadataFilePath = filepath.Join(path, "metadata.json")
		scpsFilePath     = filepath.Join(path, "scps.json")
		orgDataFilePath  = filepath.Join(path, "org_data.json")
		scpData          = make([]map[string]interface{}, len(o.AllSCPs))
		orgData          = o.Map()
	)
	if err = saveJSONToFile(o.Metadata, metadataFilePath); errUtils.HandleError(err) {
		return err
	}
	for i, scp := range o.AllSCPs {
		scpData[i] = scp.Map()
	}
	if err = saveJSONToFile(scpData, scpsFilePath); errUtils.HandleError(err) {
		return err
	}
	if err = saveJSONToFile(orgData, orgDataFilePath); errUtils.HandleError(err) {
		return err
	}
	return nil
}

// Helper function to save data as JSON to a file
func saveJSONToFile(data interface{}, filepath string) error {
	fileData, err := json.MarshalIndent(data, "", "    ")
	if errUtils.HandleError(err) {
		return err
	}
	return os.WriteFile(filepath, fileData, 0600)
}

// TODO (rsheikh): add o.CreateFromDir function
