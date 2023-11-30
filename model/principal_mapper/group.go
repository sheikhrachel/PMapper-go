package principal_mapper

import "github.com/sheikhrachel/PMapper-go/api_common/errUtils"

// Group represents an IAM group.
type Group struct {
	// Arn is the ARN of the group
	Arn string `json:"arn"`
	// AttachedPolicies is a list of policies attached to the group
	AttachedPolicies []Policy `json:"attached_policies"`
}

// Map returns a map representation of the group.
func (g *Group) Map() map[string]interface{} {
	return map[string]interface{}{
		"arn":               g.Arn,
		"attached_policies": g.AttachedPolicies,
	}
}

// Validate ensures that the group has all required fields.
func (g *Group) Validate() error {
	if g.Arn == "" {
		return errUtils.ErrArnEmpty
	}
	return nil
}
