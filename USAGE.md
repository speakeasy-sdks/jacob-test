<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	jacobtest "github.com/speakeasy-sdks/jacob-test"
)

func main() {
    s := jacobtest.New()

    ctx := context.Background()
    res, err := s.Pets.CreatePets(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->