package principal_mapper

// OrgNode represents an AWS OU.
type OrgNode struct {
	// ID is the ID of the node / AWS OU
	ID string `json:"ou_id"`
	// Name is the name of the node / AWS OU
	Name string `json:"ou_name"`
	// Accounts is a list of accounts in the node / AWS OU
	Accounts []OrgAccount `json:"accounts"`
	// ChildNodes is a list of child nodes
	ChildNodes []OrgNode `json:"child_nodes"`
	// SCPs is a list of SCPs applied to the account
	SCPs []Policy `json:"scps"`
	// Tags is a map of tags applied to the account
	Tags map[string]string `json:"tags"`
}

// NewOrgNode creates a new organization node object.
func NewOrgNode(id, name string, accounts []OrgAccount, childNodes []OrgNode, scps []Policy, tags map[string]string) (orgNode *OrgNode, err error) {
	orgNode = &OrgNode{
		ID:         id,
		Name:       name,
		Accounts:   accounts,
		ChildNodes: childNodes,
		SCPs:       scps,
		Tags:       tags,
	}
	if orgNode.Tags == nil {
		orgNode.Tags = make(map[string]string)
	}
	return orgNode, nil
}

// Map returns a map representation of the org node.
func (o *OrgNode) Map() map[string]interface{} {
	return map[string]interface{}{
		"ou_id":       o.ID,
		"ou_name":     o.Name,
		"accounts":    o.Accounts,
		"child_nodes": o.ChildNodes,
		"scps":        o.SCPs,
		"tags":        o.Tags,
	}
}
