package principal_mapper

import (
	"fmt"

	"github.com/sheikhrachel/PMapper-go/api_common/errUtils"
)

// Edge represents an edge in the organization graph.
type Edge struct {
	// Source is the source of the edge
	Source string `json:"source"`
	// Destination is the destination of the edge
	Destination string `json:"destination"`
	// Reason is the reason for the edge
	Reason string `json:"reason"`
	// ShortReason is a short reason for the edge
	ShortReason string `json:"short_reason"`
}

// NewEdge creates a new edge object.
func NewEdge(source, destination, reason, shortReason string) (edge *Edge, err error) {
	if err = validateEdge(source, destination, reason, shortReason); errUtils.HandleError(err) {
		return nil, err
	}
	return &Edge{
		Source:      source,
		Destination: destination,
		Reason:      reason,
		ShortReason: shortReason,
	}, nil
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

// validateEdge ensures that the edge has all required fields.
func validateEdge(source, destination, reason, shortReason string) (err error) {
	if source == "" {
		err = errUtils.ErrSourceEmpty
	}
	if destination == "" {
		err = errUtils.ErrDestinationEmpty
	}
	if reason == "" {
		err = errUtils.ErrReasonEmpty
	}
	if shortReason == "" {
		err = errUtils.ErrShortReasonEmpty
	}
	return err
}
