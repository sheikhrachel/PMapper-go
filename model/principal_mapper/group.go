package principal_mapper

import "github.com/sheikhrachel/PMapper-go/api_common/errUtils"

// Group represents an IAM group.
type Group struct {
	// Arn is the ARN of the group
	Arn string `json:"arn"`
	// AttachedPolicies is a list of policies attached to the group
	AttachedPolicies []Policy `json:"attached_policies"`
}

// NewGroup creates a new group object.
func NewGroup(arn string, attachedPolicies []Policy) (group *Group, err error) {
	if err = validateGroup(arn); errUtils.HandleError(err) {
		return nil, err
	}
	return &Group{
		Arn:              arn,
		AttachedPolicies: attachedPolicies,
	}, nil
}

// Map returns a map representation of the group.
func (g *Group) Map() map[string]interface{} {
	return map[string]interface{}{
		"arn":               g.Arn,
		"attached_policies": g.AttachedPolicies,
	}
}

// validateGroup ensures that the group has all required fields.
func validateGroup(arn string) error {
	if arn == "" {
		return errUtils.ErrArnEmpty
	}
	return nil
}
