package sequence_test

import (
	"context"

	sequence "github.com/attentiontech/sequencehq-go-sdk/pkg/sequence"
)

func ExampleNewSandboxClient() {
	client, _ := sequence.NewSandboxClient(sequence.WithAPIKey(sequence.APIKey{
		ClientID:     "myclient",
		ClientSecret: "mysecret"},
	))
	ctx := context.Background()
	_, _ = client.PostSeatEventsWithResponse(ctx, &sequence.PostSeatEventsParams{}, sequence.PostSeatEventsJSONRequestBody{
		EventTimestamp: "2024-01-07T00:00:00Z",
		CustomerAlias:  "clientalias",
		SeatType:       "test",
		Total:          2,
	})
}
