package awsUtils

/*
`The following are the general formats for ARNs. The specific formats depend on the resource.
To use an ARN, replace the italicized text with the resource-specific information.
Be aware that the ARNs for some resources omit the Region, the account ID, or both the Region and the account ID.`

```
arn:partition:service:region:account-id:resource-id
arn:partition:service:region:account-id:resource-type/resource-id
arn:partition:service:region:account-id:resource-type:resource-id
```

Documentation: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
*/

import (
	"strings"

	"github.com/sheikhrachel/PMapper-go/model/aws"
)

// GetArnComponent returns the specified component of an ARN
func GetArnComponent(arn string, component aws.ArnComponent) string {
	if !IsArnValid(arn) {
		return ""
	}
	return strings.Split(arn, ":")[component.Int()]
}

func IsArnValid(arn string) (valid bool) {
	arnSlice := strings.Split(arn, ":")
	return len(arnSlice) >= 6 && arnSlice[0] == "arn"
}
