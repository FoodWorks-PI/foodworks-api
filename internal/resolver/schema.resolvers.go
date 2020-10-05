package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"foodworks.ml/m/internal/auth"
	"foodworks.ml/m/internal/generated/ent"
	"foodworks.ml/m/internal/generated/ent/customer"
	"foodworks.ml/m/internal/generated/ent/restaurantowner"
	generated "foodworks.ml/m/internal/generated/graphql"
	"foodworks.ml/m/internal/generated/graphql/model"
)

func (r *customerResolver) Address(ctx context.Context, obj *ent.Customer) (*ent.Address, error) {
	address, err := r.Client.Customer.QueryAddress(obj).First(ctx)
	return address, ent.MaskNotFound(err)
}

func (r *mutationResolver) CreateCustomerProfile(ctx context.Context, input model.RegisterCustomerInput) (int, error) {
	currentUser := auth.ForContext(ctx)

	newAddress, err := r.Client.Address.
		Create().
		SetLatitude(input.Address.Latitude).
		SetLongitude(input.Address.Longitude).
		SetStreetLine(input.Address.StreetLine).
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
		SetAddress(newAddress).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	return newCustomer.ID, nil
}

func (r *mutationResolver) CreateRestaurantOwnerProfile(ctx context.Context, input model.RegisterRestaurantOwnerInput) (int, error) {
	currentUser := auth.ForContext(ctx)

	newBankingData, err := r.Client.BankingData.
		Create().
		SetBankAccount(input.Banking.BankAccount).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	newRestaurantOwner, err := r.Client.RestaurantOwner.
		Create().
		SetName(input.Name).
		SetEmail(currentUser.Email).
		SetKratosID(currentUser.Id).
		SetPhone(input.Phone).
		SetBankingData(newBankingData).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	return newRestaurantOwner.ID, nil
}

func (r *queryResolver) GetCurrentCustomer(ctx context.Context) (*ent.Customer, error) {
	kratosUser := auth.ForContext(ctx)

	currentCustomer, err := r.Client.Customer.
		Query().
		Where(customer.KratosID(kratosUser.Id)).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return currentCustomer, nil
}

func (r *queryResolver) GetCurrentRestaurantOwner(ctx context.Context) (*ent.RestaurantOwner, error) {
	kratosUser := auth.ForContext(ctx)

	currentRestaurantOwner, err := r.Client.RestaurantOwner.
		Query().
		Where(restaurantowner.KratosID(kratosUser.Id)).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return currentRestaurantOwner, nil
}

func (r *restaurantOwnerResolver) Banking(ctx context.Context, obj *ent.RestaurantOwner) (*ent.BankingData, error) {
	banking, err := r.Client.RestaurantOwner.QueryBankingData(obj).First(ctx)
	return banking, ent.MaskNotFound(err)
}

// Customer returns generated.CustomerResolver implementation.
func (r *Resolver) Customer() generated.CustomerResolver { return &customerResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// RestaurantOwner returns generated.RestaurantOwnerResolver implementation.
func (r *Resolver) RestaurantOwner() generated.RestaurantOwnerResolver {
	return &restaurantOwnerResolver{r}
}

type customerResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type restaurantOwnerResolver struct{ *Resolver }
