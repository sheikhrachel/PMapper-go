package errUtils

import "errors"

var (
	ErrArnEmpty                      = errors.New("resource must have an ARN")
	ErrSourceEmpty                   = errors.New("edge must have a source node object")
	ErrDestinationEmpty              = errors.New("edge must have a destination node object")
	ErrReasonEmpty                   = errors.New("edge must have a reason string")
	ErrShortReasonEmpty              = errors.New("edge must have a short reason string")
	ErrPolicyDocNil                  = errors.New("policy must have a non-nil policy document")
	ErrNodeIDValueEmpty              = errors.New("node must have a non-empty ID value")
	ErrNodeUserNonNilTrustPolicy     = errors.New("IAM users do not have trust policies, pass nil for the parameter trust_policy")
	ErrNodeRoleNilTrustPolicy        = errors.New("IAM roles have trust policies, which must be passed as a map in trust_policy")
	ErrNodeUserNonNilInstanceProfile = errors.New("IAM users do not have instance profiles. Pass nil for the parameter instance_profile")
	ErrNodeNilArnForUserOrRole       = errors.New("ARN must be non-empty and start with either 'user/' or 'role/'")
	ErrAppDataNotSet                 = errors.New("%APPDATA% was unexpectedly not set")
)
