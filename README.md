# Seed Loader

This is used to load a json seed for your tests or any parsing of plain json files

See usage below

```go
package myprogram

import (
	"github.com/borderlesshq/seedloader"
	"log"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	// Note, ./seeds is a directory where you've decided to store plain json files
	seeder, err := seedloader.NewSeedLoader("myprogram", "seeds")

	if err != nil {
		log.Fatalf("failed to initialize seed loader %v", err)
	}

	var user User
	if err := seeder.ParseSeed("user.json", &user); err != nil {
		log.Fatalf("failed to parse seed %v", err)
	}

	// output: { Borderless HQ Inc. }
	log.Printf("output: %v", user)
}
```