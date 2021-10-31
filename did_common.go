package main

import (
  "fmt"
	"did"
)

func main() {
	did := did.New("did:example:123")
	fmt.Printf("%#v", did)
}
