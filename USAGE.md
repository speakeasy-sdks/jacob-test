<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	jacobtest "github.com/speakeasy-sdks/jacob-test/v3"
	"github.com/speakeasy-sdks/jacob-test/v3/pkg/models/shared"
	"log"
)

func main() {
	s := jacobtest.New()

	ctx := context.Background()
	res, err := s.Pets.CreatePets(ctx, shared.Pet{
		ID:   596804,
		Name: "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->