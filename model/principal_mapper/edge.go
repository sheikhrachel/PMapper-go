package principal_mapper

import (
	"fmt"

	"github.com/sheikhrachel/PMapper-go/api_common/errUtils"
)

type Edge struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Reason      string `json:"reason"`
	ShortReason string `json:"short_reason"`
}

const describeTmpl = "{%+v} {%+v} {%+v}"

// Describe returns a string representation of the edge.
func (e *Edge) Describe() string {
	return fmt.Sprintf(describeTmpl, e.Source, e.Destination, e.Reason)
}

// Map returns a map representation of the edge.
func (e *Edge) Map() map[string]string {
	return map[string]string{
		"source":       e.Source,
		"destination":  e.Destination,
		"reason":       e.Reason,
		"short_reason": e.ShortReason,
	}
}

// Validate ensures that the edge has all required fields.
func (e *Edge) Validate() (err error) {
	if e.Source == "" {
		err = errUtils.ErrSourceEmpty
	}
	if e.Destination == "" {
		err = errUtils.ErrDestinationEmpty
	}
	if e.Reason == "" {
		err = errUtils.ErrReasonEmpty
	}
	if e.ShortReason == "" {
		err = errUtils.ErrShortReasonEmpty
	}
	return err
}
