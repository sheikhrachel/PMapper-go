package errUtils

import "errors"

var (
	ErrArnEmpty         = errors.New("resource must have an ARN")
	ErrSourceEmpty      = errors.New("edge must have a source node object")
	ErrDestinationEmpty = errors.New("edge must have a destination node object")
	ErrReasonEmpty      = errors.New("edge must have a reason string")
	ErrShortReasonEmpty = errors.New("edge must have a short reason string")
	ErrPolicyDocNil     = errors.New("policy must have a non-nil policy document")
	ErrAppDataNotSet    = errors.New("%APPDATA% was unexpectedly not set")
)
