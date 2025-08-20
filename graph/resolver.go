package graph

import "github.com/elyseeMB/relay-compiler/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Articles []*model.Article
	Authors  []*model.Author
}
