package bitgrail

const (
	testingApiKey    = "YOUR API KEY"
	testingApiSecret = "YOUR API SECRET"
)

func newAuthClient() *Client {
	return NewClient().Auth(testingApiKey, testingApiSecret)
}
