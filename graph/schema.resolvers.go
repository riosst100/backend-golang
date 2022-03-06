package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"inovasi-aktif-go/graph/generated"
	"inovasi-aktif-go/graph/model"
	"inovasi-aktif-go/internal/auth"
	"inovasi-aktif-go/internal/repository"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) Register(ctx context.Context, input model.NewUser) (string, error) {
	return repository.CreateUser(input)
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginUser) (string, error) {
	return repository.AuthenticateUser(input)
}

func (r *queryResolver) UserList(ctx context.Context) ([]*model.User, error) {
	return repository.UserList()
}

func (r *queryResolver) BusinessList(ctx context.Context) ([]*model.Business, error) {
	return repository.BusinessList()
}

func (r *queryResolver) ProductList(ctx context.Context) ([]*model.Product, error) {
	return repository.ProductList()
}

func (r *queryResolver) UserByID(ctx context.Context, id string) (*model.User, error) {
	return repository.UserByID(id)
}

func (r *queryResolver) BusinessByID(ctx context.Context, id string) (*model.Business, error) {
	return repository.BusinessByID(id)
}

func (r *queryResolver) ProductByID(ctx context.Context, id string) (*model.Product, error) {
	return repository.ProductByID(id)
}

func (r *queryResolver) OrderByID(ctx context.Context, id string) (*model.Order, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, gqlerror.Errorf("Mohon daftar/masuk terlebih dahulu")
	}

	return repository.OrderByID(id)
}

func (r *queryResolver) BusinessByUserID(ctx context.Context, userID string) ([]*model.Business, error) {
	return repository.BusinessByUserID(userID)
}

func (r *queryResolver) ProductByBusinessID(ctx context.Context, businessID string) ([]*model.Product, error) {
	return repository.ProductByBusinessID(businessID)
}

func (r *queryResolver) DesaByKecamatanID(ctx context.Context, kecamatanID string) ([]*model.Desa, error) {
	return repository.DesaByKecamatanID(kecamatanID)
}

func (r *queryResolver) KecamatanList(ctx context.Context) ([]*model.Kecamatan, error) {
	return repository.KecamatanList()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
