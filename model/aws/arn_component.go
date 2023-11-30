package aws

type ArnComponent int

const (
	Partition ArnComponent = iota + 1
	Service
	Region
	AccountID
	ResourceType
)

// Int returns the integer representation of the ArnComponent.
func (a ArnComponent) Int() int {
	return int(a)
}
