package main

import (
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgraph/gql"
)

func prettyPrint(v interface{}) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return err
}

func main() {
	query := `{
  q(func: uid(0x5)) {
    name
    posts {
      title
    }
  }
}`
	g := gql.Request{
		Str: query,
	}
	result, err := gql.Parse(g)
	if err != nil {
		fmt.Printf("Err during parsing: %v\n", err)
	}
	fmt.Printf("Query:\n%s\n", query)
	fmt.Println("---")
	fmt.Println("Parsed:")
	prettyPrint(result.Query[0])
}
