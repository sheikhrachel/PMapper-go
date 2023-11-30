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

// Map returns a map representation of the org account.
func (o *OrgAccount) Map() map[string]interface{} {
	if o.Tags == nil {
		o.Tags = make(map[string]string)
	}
	return map[string]interface{}{
		"account_id": o.AccountID,
		"scps":       o.SCPs,
		"tags":       o.Tags,
	}
}
