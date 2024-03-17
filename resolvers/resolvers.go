package resolvers

import (
	"github.com/graphql-go/graphql"
)

func TodoResolver(p graphql.ResolveParams) (interface{}, error) {
	// Implement your resolver here
	// For now, returning a dummy todo
	id := p.Args["id"].(string)
	return Todo{ID: id, Text: "Implement me", Done: false}, nil
}
