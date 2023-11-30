package awsUtils

type AWSClient struct {
}

func Init() (awsClient *AWSClient) {
	awsClient = &AWSClient{}
	return awsClient
}
