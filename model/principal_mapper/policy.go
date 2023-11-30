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

// NewPolicy creates a new policy object.
func NewPolicy(arn, name string, policyDoc map[string]interface{}) (policy *Policy, err error) {
	if err = validatePolicy(arn, policyDoc); errUtils.HandleError(err) {
		return nil, err
	}
	return &Policy{
		Arn:       arn,
		Name:      name,
		PolicyDoc: policyDoc,
	}, nil
}

// Map returns a map representation of the policy.
func (p *Policy) Map() map[string]interface{} {
	return map[string]interface{}{
		"arn":        p.Arn,
		"name":       p.Name,
		"policy_doc": p.PolicyDoc,
	}
}

// validatePolicy ensures that the policy has all required fields.
func validatePolicy(arn string, policyDoc map[string]interface{}) error {
	if arn == "" {
		return errUtils.ErrArnEmpty
	}
	if policyDoc == nil {
		return errUtils.ErrPolicyDocNil
	}
	return nil
}
