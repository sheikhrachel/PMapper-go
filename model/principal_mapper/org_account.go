package principal_mapper

// OrgAccount represents an AWS account.
type OrgAccount struct {
	// AccountID is the ID of the aws account
	AccountID string `json:"account_id"`
	// SCPs is a list of SCPs applied to the account
	SCPs []Policy `json:"scps"`
	// Tags is a map of tags applied to the account
	Tags map[string]string `json:"tags"`
}

// NewOrgAccount creates a new organization account object.
func NewOrgAccount(accountID string, scps []Policy, tags map[string]string) (orgAccount *OrgAccount, err error) {
	orgAccount = &OrgAccount{
		AccountID: accountID,
		SCPs:      scps,
		Tags:      tags,
	}
	if orgAccount.Tags == nil {
		orgAccount.Tags = make(map[string]string)
	}
	return orgAccount, nil
}

// Map returns a map representation of the org account.
func (o *OrgAccount) Map() map[string]interface{} {
	return map[string]interface{}{
		"account_id": o.AccountID,
		"scps":       o.SCPs,
		"tags":       o.Tags,
	}
}
