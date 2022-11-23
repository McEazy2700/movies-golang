package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/mceazy2700/movies-golang/graph/generated"
	"github.com/mceazy2700/movies-golang/graph/model"
)

// ID is the resolver for the id field.
func (r *movieResolver) ID(ctx context.Context, obj *model.Movie) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Views is the resolver for the views field.
func (r *movieResolver) Views(ctx context.Context, obj *model.Movie) (int, error) {
	panic(fmt.Errorf("not implemented: Views - views"))
}

// UpdateViewsCount is the resolver for the updateViewsCount field.
func (r *mutationResolver) UpdateViewsCount(ctx context.Context, movieID string) (*model.MovieSuccess, error) {
	panic(fmt.Errorf("not implemented: UpdateViewsCount - updateViewsCount"))
}

// PopularMovies is the resolver for the popularMovies field.
func (r *queryResolver) PopularMovies(ctx context.Context) ([]*model.Movie, error) {
	panic(fmt.Errorf("not implemented: PopularMovies - popularMovies"))
}

// Movie returns generated.MovieResolver implementation.
func (r *Resolver) Movie() generated.MovieResolver { return &movieResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type movieResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
