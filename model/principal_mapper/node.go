package principal_mapper

import (
	"strings"

	"github.com/sheikhrachel/PMapper-go/api_common/errUtils"
)

// Node represents a node in the organization graph.
type Node struct {
	// ARN is the ARN of the node
	ARN string `json:"arn"`
	// IDValue is the ID value of the node
	IDValue string `json:"id_value"`
	// AttachedPolicies is a list of policies attached to the node, attached or inline
	AttachedPolicies []Policy `json:"attached_policies"`
	// GroupMemberships is a list of groups the node is a member of
	GroupMemberships []Group `json:"group_memberships"`
	// TrustPolicy is the trust policy of the node, if the node represents an IAM Role
	TrustPolicy map[string]interface{} `json:"trust_policy"`
	// InstanceProfile is the instance profile of the node, if the node represents an IAM Role
	InstanceProfile []string `json:"instance_profile"`
	// ActivePassword is true if the node has an active password, false otherwise, if the node represents an IAM User
	ActivePassword bool `json:"active_password"`
	// AccessKeys is the number of access keys the node has, if the node represents an IAM User
	AccessKeysCount int `json:"num_access_keys"`
	// IsAdmin is true if the node is an admin, false otherwise
	IsAdmin bool `json:"is_admin"`
	// PermissionsBoundary is the permissions boundary of the node
	PermissionsBoundary *Policy `json:"permissions_boundary"`
	// HasMFA is true if the node has MFA enabled, false otherwise
	HasMFA bool `json:"has_mfa"`
	// Tags is a map of tags applied to the node
	Tags map[string]string `json:"tags"`
	// Cache is a map of in-memory cached values for the node
	Cache map[string]interface{} `json:"-"`
}

// NewNode creates a new node object.
func NewNode(
	arn, idValue string,
	attachedPolicies []Policy,
	groupMemberships []Group,
	trustPolicy map[string]interface{},
	instanceProfile []string,
	numAccessKeys int,
	activePassword, isAdmin, hasMFA bool,
	permissionsBoundary *Policy,
	tags map[string]string,
) (node *Node, err error) {
	if err = validateNode(arn, idValue, trustPolicy, instanceProfile); errUtils.HandleError(err) {
		return nil, err
	}
	return &Node{
		ARN:                 arn,
		IDValue:             idValue,
		AttachedPolicies:    attachedPolicies,
		GroupMemberships:    groupMemberships,
		TrustPolicy:         trustPolicy,
		InstanceProfile:     instanceProfile,
		ActivePassword:      activePassword,
		AccessKeysCount:     numAccessKeys,
		IsAdmin:             isAdmin,
		PermissionsBoundary: permissionsBoundary,
		HasMFA:              hasMFA,
		Tags:                tags,
		Cache:               make(map[string]interface{}),
	}, nil
}

const (
	userPrefix = "user/"
	rolePrefix = "role/"
)

// validateNode ensures that the node has all required fields.
func validateNode(
	arn, idValue string,
	trustPolicy map[string]interface{},
	instanceProfile []string,
) error {
	if arn == "" || !(strings.HasPrefix(arn, userPrefix) || strings.HasPrefix(arn, rolePrefix)) {
		return errUtils.ErrNodeNilArnForUserOrRole
	}
	if idValue == "" {
		return errUtils.ErrNodeIDValueEmpty
	}
	if strings.HasPrefix(arn, userPrefix) && trustPolicy != nil {
		return errUtils.ErrNodeUserNonNilTrustPolicy
	}
	if strings.HasPrefix(arn, rolePrefix) && trustPolicy == nil {
		return errUtils.ErrNodeRoleNilTrustPolicy
	}
	if strings.HasPrefix(arn, userPrefix) && instanceProfile != nil {
		return errUtils.ErrNodeUserNonNilInstanceProfile
	}
	return nil
}
