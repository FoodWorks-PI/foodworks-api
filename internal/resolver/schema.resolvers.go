package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"foodworks.ml/m/internal/generated/ent"
	"foodworks.ml/m/internal/generated/ent/user"
	generated "foodworks.ml/m/internal/generated/graphql"
	"foodworks.ml/m/internal/generated/graphql/model"
	"github.com/99designs/gqlgen/graphql"
)

func (r *mutationResolver) SetUserIdentity(ctx context.Context, input model.IdentityInput) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productResolver) Name(ctx context.Context, obj *ent.Product) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productResolver) Description(ctx context.Context, obj *ent.Product) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCustomerByEmail(ctx context.Context, email string) (*model.Customer, error) {
	user, err := r.Client.User.
		Query().
		Where(user.EmailEQ(email)).
		First(ctx)
	if err != nil {
		graphql.AddErrorf(ctx, "Error %d", err)
		return nil, err
	}
	pIdentity := &model.Identity{
		Email: user.Email,
	}
	customer := &model.Customer{
		Identity: pIdentity,
	}
	return customer, nil
}

func (r *queryResolver) GetRestaurantOffers(ctx context.Context, restaurantID int) ([]*ent.Restaurant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *restaurantResolver) Name(ctx context.Context, obj *ent.Restaurant) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *restaurantResolver) Address(ctx context.Context, obj *ent.Restaurant) ([]string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Restaurant returns generated.RestaurantResolver implementation.
func (r *Resolver) Restaurant() generated.RestaurantResolver { return &restaurantResolver{r} }

type mutationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type restaurantResolver struct{ *Resolver }
