package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	// "fmt"
	"project-management/configs"
	"project-management/graph/model"
)
var (
	db = configs.MongoConnect()
)

// CreateOwner is the resolver for the createOwner field.
func (r *mutationResolver) CreateOwner(ctx context.Context, input *model.NewOwner) (*model.Owner, error) {
	owner , err := db.CreateOwner(input)
	if err != nil {
		return nil, err
	}
	return owner, nil
}

// CreateProject is the resolver for the createProject field.
func (r *mutationResolver) CreateProject(ctx context.Context, input *model.NewProject) (*model.Project, error) {
	project, err := db.CreateProject(input)
	if err != nil {
		return nil, err
	}
	return project, nil
}

// Owners is the resolver for the owners field.
func (r *queryResolver) Owners(ctx context.Context) ([]*model.Owner, error) {
	owners, err := db.GetOwners()
	if err != nil {
		return nil, err
	}
	return owners, nil
}

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	projects, err := db.GetProjects()
	if err != nil {
		return nil, err
	}
	return projects, nil
}

// Owner is the resolver for the owner field.
func (r *queryResolver) Owner(ctx context.Context, input *model.FetchOwnerInput) (*model.Owner, error) {
	owner , err := db.SingleOwner(input.ID)
	if err!=nil{
		return nil ,err
	}
	return owner, err
}

// Project is the resolver for the project field.
func (r *queryResolver) Project(ctx context.Context, input *model.FetchProject) (*model.Project, error) {
	project , err := db.SingleProject(input.ID)
	if err!=nil {
		return nil , err
	}
	return project ,err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
// func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
// 	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
// }
// func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
// 	panic(fmt.Errorf("not implemented: Todos - todos"))
// }