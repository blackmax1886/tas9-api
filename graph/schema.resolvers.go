package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
	"strconv"

	"github.com/blackmax1886/tas9-api/graph/entity"
	"github.com/blackmax1886/tas9-api/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := entity.User{
		Name:  input.Name,
		Email: input.Email,
	}
	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return entity.NewUserFromEntity(&user), nil
}

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	panic("not impremented")
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	var u entity.User
	if err := r.DB.Find(&u, userID).Error; err != nil {
		return nil, err
	}

	return &model.User{
		ID:    fmt.Sprintf("%d", u.ID),
		Name:  u.Name,
		Email: u.Email,
	}, nil
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	panic("not impremented")
}

// Subtasks is the resolver for the subtasks field.
func (r *queryResolver) Subtasks(ctx context.Context) ([]*model.Subtask, error) {
	panic("not impremented")
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
