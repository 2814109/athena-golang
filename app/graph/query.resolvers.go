package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"
	"my_gql_server/graph/model"
	"my_gql_server/infrastructures/repositories"
	"my_gql_server/my_models"

	lop "github.com/samber/lo/parallel"
)

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context, userID int) ([]*models.Todo, error) {
	todos, err := repositories.FindAllTodoByUserId(ctx, userID)

	if err != nil {
		return nil, err
	}

	result := lop.Map(todos, func(todo *models.Todo, _ int) *models.Todo {
		return todo

	})

	return result, nil
}

// Articles is the resolver for the articles field.
func (r *queryResolver) Articles(ctx context.Context) ([]*model.Article, error) {
	panic(fmt.Errorf("not implemented: Articles - articles"))
}

// Items is the resolver for the items field.
func (r *queryResolver) Items(ctx context.Context, userID string) ([]*model.Item, error) {
	panic(fmt.Errorf("not implemented: Items - items"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
