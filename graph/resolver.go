package graph

import "github.com/blackmax1886/tas9-api/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	tasks    []*model.Task
	subtasks []*model.Subtask
}
