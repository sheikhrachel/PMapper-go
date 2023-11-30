package principal_mapper

import "github.com/sheikhrachel/PMapper-go/api_common/errUtils"

// Policy represents an IAM policy.
type Policy struct {
	// Arn is the ARN of the policy
	Arn string `json:"arn"`
	// Name is the name of the policy
	Name string `json:"name"`
	// PolicyDoc is a map containing the contents of the IAM Policy contents
	PolicyDoc map[string]interface{} `json:"policy_doc"`
}

// Map returns a map representation of the policy.
func (p *Policy) Map() map[string]interface{} {
	return map[string]interface{}{
		"arn":        p.Arn,
		"name":       p.Name,
		"policy_doc": p.PolicyDoc,
	}
}

// Validate ensures that the policy has all required fields.
func (p *Policy) Validate() (err error) {
	if p.Arn == "" {
		err = errUtils.ErrArnEmpty
	}
	if p.PolicyDoc == nil {
		err = errUtils.ErrPolicyDocNil
	}
	return err
}
