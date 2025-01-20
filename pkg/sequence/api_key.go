package sequence

import (
	"context"
	"encoding/base64"
	"net/http"
)

// APIKey is a struct that holds the client_id and client_secret for Sequence API
type APIKey struct {
	ClientID     string `json:"client_id" yaml:"client_id"`
	ClientSecret string `json:"client_secret" yaml:"client_secret"`
}

// Base64 returns the base64 encoded string of the APIKey
func (apiKey APIKey) Base64() string {
	auth := apiKey.ClientID + ":" + apiKey.ClientSecret
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// AuthorizationHeaderValue returns the string used as value in the Authorization header, already encoded for Basic authentication
func (apiKey APIKey) AuthorizationHeaderValue() string {
	return "Basic " + apiKey.Base64()
}

// WithAPIKey is used to configure a client with a global API key for Sequence API
func WithAPIKey(apiKey APIKey) ClientOption {
	encodedAuthHeaderValue := apiKey.AuthorizationHeaderValue()
	return WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", encodedAuthHeaderValue)
		return nil
	})
}
