package sequence

// APIURL is the URL of the Sequence API
type APIURL string

const (
	// ProductionEuropeAPIURL is the URL of the Sequence API in production
	ProductionEuropeAPIURL = APIURL("https://eu.sequencehq.com/api")
	// SandboxAPIURL is the URL of the Sequence API in sandbox
	SandboxAPIURL = APIURL("https://sandbox.sequencehq.com/api")
)

// NewProductionEuropeClient creates a new client for the Sequence API in production in Europe
func NewProductionEuropeClient(opts ...ClientOption) (*ClientWithResponses, error) {
	return NewClientWithResponses(string(ProductionEuropeAPIURL), opts...)
}

// NewSandboxClient creates a new client for the Sequence API in sandbox
func NewSandboxClient(opts ...ClientOption) (*ClientWithResponses, error) {
	return NewClientWithResponses(string(SandboxAPIURL), opts...)
}
