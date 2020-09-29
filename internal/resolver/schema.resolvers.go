package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"foodworks.ml/m/internal/auth"
	"foodworks.ml/m/internal/generated/ent"
	generated "foodworks.ml/m/internal/generated/graphql"
	"foodworks.ml/m/internal/generated/graphql/model"
)

func (r *addressResolver) StreetLine(ctx context.Context, obj *ent.Address) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *customerResolver) Address(ctx context.Context, obj *ent.Customer) (*ent.Address, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCustomerProfile(ctx context.Context, input model.RegisterCustomerInput) (int, error) {
	currentUser := auth.ForContext(ctx)

	newAddress, err := r.Client.Address.
		Create().
		SetLatitude(input.Address.Latitude).
		SetLongitude(input.Address.Longitude).
		SetStreet(input.Address.StreetLine).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	newCustomer, err := r.Client.Customer.
		Create().
		SetName(input.Name).
		SetEmail(currentUser.Email).
		SetKratosID(currentUser.Id).
		SetPhone(input.Phone).
		AddAddress(newAddress).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	return newCustomer.ID, nil
}

func (r *queryResolver) GetCurrentCustomer(ctx context.Context) (*ent.Customer, error) {
	panic(fmt.Errorf("not implemented"))
}

// Address returns generated.AddressResolver implementation.
func (r *Resolver) Address() generated.AddressResolver { return &addressResolver{r} }

// Customer returns generated.CustomerResolver implementation.
func (r *Resolver) Customer() generated.CustomerResolver { return &customerResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type addressResolver struct{ *Resolver }
type customerResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
