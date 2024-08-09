<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/log10-io/log10go"
	"log"
)

func main() {
	s := log10go.New(
		log10go.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Sessions.Create(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->