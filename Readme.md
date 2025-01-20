# sequencehq-go-sdk

Automatically generated SDK for https://www.sequencehq.com/

## Usage && Examples

```go
import sequence "github.com/attentiontech/sequencehq-go-sdk/pkg/sequence"

...

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
```

**For production**, use `sequence.NewProductionEuropeClient`.

See `pkg/example_test.go`

## Updating

```
scripts/generate.sh
```

## Test

```
go test ./...
```
