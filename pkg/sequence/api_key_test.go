package sequence_test

import (
	"testing"

	"github.com/attentiontech/sequencehq-go-sdk/pkg/sequence"
)

func TestAPIKey(t *testing.T) {
	key := sequence.APIKey{
		ClientID:     "myclient",
		ClientSecret: "mysecret",
	}
	if key.Base64() != "bXljbGllbnQ6bXlzZWNyZXQ=" {
		t.Error("Base64() did not return the expected value")
	}
	if key.AuthorizationHeaderValue() != "Basic bXljbGllbnQ6bXlzZWNyZXQ=" {
		t.Error("AuthorizationHeaderValue() did not return the expected value")
	}
}
