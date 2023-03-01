package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"errors"
	"fmt"

	"github.com/blackmax1886/tas9-api/graph/model"
	ulid "github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := model.User{
		ID:    ulid.Make().String(),
		Name:  input.Name,
		Email: input.Email,
	}
	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	task := model.Task{
		ID:       ulid.Make().String(),
		Name:     input.Name,
		Content:  input.Content,
		Done:     false,
		Archived: false,
		UserID:   &input.UserID,
	}
	if err := r.DB.Create(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

// CreateSubTask is the resolver for the createSubTask field.
func (r *mutationResolver) CreateSubTask(ctx context.Context, input model.NewSubtask) (*model.Subtask, error) {
	subtask := model.Subtask{
		ID:       ulid.Make().String(),
		Name:     input.Name,
		Content:  input.Content,
		Done:     false,
		Archived: false,
		TaskID:   &input.TaskID,
	}
	if err := r.DB.Create(&subtask).Error; err != nil {
		return nil, err
	}
	return &subtask, nil
}

// LinkAccount is the resolver for the linkAccount field.
func (r *mutationResolver) LinkAccount(ctx context.Context, input *model.Account) (*model.User, error) {
	var user model.User
	switch input.Provider {
	case "google":
		err := r.DB.Find(&user, "id = ?", input.UserID).Update("google_id", input.ProviderAccountID).Error
		if err != nil {
			return nil, err
		}
		return &user, nil
	default:
		panic("This provider is not supported")
	}
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	if err := r.DB.Find(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context, userID *string) ([]*model.Task, error) {
	var tasks []*model.Task
	if err := r.DB.Find(&tasks, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// Subtasks is the resolver for the subtasks field.
func (r *queryResolver) Subtasks(ctx context.Context, taskID *string) ([]*model.Subtask, error) {
	var subtasks []*model.Subtask
	if err := r.DB.Find(&subtasks, "task_id = ?", taskID).Error; err != nil {
		return nil, err
	}
	return subtasks, nil
}

// UserByEmail is the resolver for the userByEmail field.
func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UserByAccount is the resolver for the userByAccount field.
func (r *queryResolver) UserByAccount(ctx context.Context, partialAccount model.PartialAccount) (*model.User, error) {
	var user model.User
	switch partialAccount.Provider {
	case "google":
		err := r.DB.Where("google_id = ?", partialAccount.ProviderAccountID).First(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}
		return &user, nil
	default:
		return nil, fmt.Errorf("This provider is not supported")
	}
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
