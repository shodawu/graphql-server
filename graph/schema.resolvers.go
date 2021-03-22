package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-users/graph/generated"
	"go-users/graph/model"
)

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var modUsers []*model.User
	er := r.DB.Model(&modUsers).Find(&modUsers).Error

	return modUsers, er
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
