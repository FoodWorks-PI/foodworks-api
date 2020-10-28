package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"foodworks.ml/m/internal/auth"
	"foodworks.ml/m/internal/generated/ent"
	"foodworks.ml/m/internal/generated/ent/address"
	"foodworks.ml/m/internal/generated/ent/customer"
	"foodworks.ml/m/internal/generated/ent/product"
	"foodworks.ml/m/internal/generated/ent/rating"
	"foodworks.ml/m/internal/generated/ent/restaurant"
	"foodworks.ml/m/internal/generated/ent/restaurantowner"
	"foodworks.ml/m/internal/generated/ent/tag"
	generated "foodworks.ml/m/internal/generated/graphql"
	"foodworks.ml/m/internal/generated/graphql/model"
	gabs "github.com/Jeffail/gabs/v2"
	"github.com/facebook/ent/dialect/sql"
	"github.com/rs/zerolog/log"
)

func (r *customerResolver) Address(ctx context.Context, obj *ent.Customer) (*ent.Address, error) {
	address, err := r.EntClient.Customer.QueryAddress(obj).First(ctx)
	return address, ent.MaskNotFound(err)
}

func (r *customerResolver) RatedProducts(ctx context.Context, obj *ent.Customer) ([]*ent.Rating, error) {
	ratings, err := r.EntClient.Customer.QueryRatings(obj).All(ctx)
	return ratings, ent.MaskNotFound(err)
}

func (r *mutationResolver) CreateCustomerProfile(ctx context.Context, input model.RegisterCustomerInput) (int, error) {
	currentUser := auth.ForContext(ctx)

	newAddress, err := r.EntClient.Address.
		Create().
		SetLatitude(input.Address.Latitude).
		SetLongitude(input.Address.Longitude).
		SetStreetLine(input.Address.StreetLine).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	newCustomer, err := r.EntClient.Customer.
		Create().
		SetName(input.Name).
		SetEmail(currentUser.Email).
		SetKratosID(currentUser.ID).
		SetPhone(input.Phone).
		SetAddress(newAddress).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	return newCustomer.ID, nil
}

func (r *mutationResolver) UpdateCustomerProfile(ctx context.Context, input model.UpdateCustomerInput) (int, error) {
	kratosSessionUser := auth.ForContext(ctx)

	currentUser, err := r.EntClient.Customer.
		Query().
		Where(customer.KratosID(kratosSessionUser.ID)).
		First(ctx)

	if err != nil {
		return -1, err
	}

	_, err = currentUser.
		Update().
		SetName(input.Name).
		SetPhone(input.Phone).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	return currentUser.ID, nil
}

func (r *mutationResolver) UpdateCustomerAddress(ctx context.Context, input model.RegisterAddressInput) (int, error) {
	kratosSessionUser := auth.ForContext(ctx)

	currentUser, err := r.EntClient.Customer.
		Query().
		Where(customer.KratosID(kratosSessionUser.ID)).
		WithAddress().
		First(ctx)

	if err != nil {
		return -1, err
	}

	updatedAddress, err := currentUser.Edges.Address.
		Update().
		SetLatitude(input.Latitude).
		SetLongitude(input.Longitude).
		SetStreetLine(input.StreetLine).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	_, err = currentUser.
		Update().
		SetAddress(updatedAddress).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	return currentUser.ID, nil
}

func (r *mutationResolver) DeleteCustomerProfile(ctx context.Context) (int, error) {
	kratosSessionUser := auth.ForContext(ctx)

	currentUser, err := r.EntClient.Customer.
		Query().
		Where(customer.KratosID(kratosSessionUser.ID)).
		First(ctx)

	if err != nil {
		return -1, err
	}

	err = r.EntClient.Customer.DeleteOne(currentUser).Exec(ctx)

	if err != nil {
		return -1, err
	}

	return currentUser.ID, nil
}

func (r *mutationResolver) CreateRestaurantOwnerProfile(ctx context.Context, input model.RegisterRestaurantOwnerInput) (int, error) {
	currentUser := auth.ForContext(ctx)

	newBankingData, err := r.EntClient.BankingData.
		Create().
		SetBankAccount(input.Banking.BankAccount).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	newAddress, err := r.EntClient.Address.
		Create().
		SetLatitude(input.Restaurant.Address.Latitude).
		SetLongitude(input.Restaurant.Address.Longitude).
		SetStreetLine(input.Restaurant.Address.StreetLine).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	tagEntities, err := GetOrCreateTagId(r.Resolver, input.Restaurant.Tags, ctx)

	newRestaurant, err := r.EntClient.Restaurant.
		Create().
		SetName(input.Restaurant.Name).
		SetDescription(input.Restaurant.Description).
		SetAddress(newAddress).
		AddTags(tagEntities...).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	newRestaurantOwner, err := r.EntClient.RestaurantOwner.
		Create().
		SetName(input.Name).
		SetEmail(currentUser.Email).
		SetKratosID(currentUser.ID).
		SetPhone(input.Phone).
		SetBankingData(newBankingData).
		SetRestaurant(newRestaurant).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	return newRestaurantOwner.ID, nil
}

func (r *mutationResolver) UpdateRestaurantOwnerProfile(ctx context.Context, input model.UpdateRestaurantOwnerInput) (int, error) {
	kratosSessionUser := auth.ForContext(ctx)

	currentUser, err := r.EntClient.RestaurantOwner.
		Query().
		Where(restaurantowner.KratosID(kratosSessionUser.ID)).
		First(ctx)

	if err != nil {
		return -1, err
	}

	_, err = currentUser.
		Update().
		SetName(input.Name).
		SetPhone(input.Phone).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	return currentUser.ID, nil
}

func (r *mutationResolver) UpdateRestaurantOwnerBankingData(ctx context.Context, input model.RegisterBankingInput) (int, error) {
	kratosSessionUser := auth.ForContext(ctx)

	currentUser, err := r.EntClient.RestaurantOwner.
		Query().
		Where(restaurantowner.KratosID(kratosSessionUser.ID)).
		WithBankingData().
		First(ctx)

	if err != nil {
		return -1, err
	}

	updatedBankingData, err := currentUser.Edges.BankingData.
		Update().
		SetBankAccount(input.BankAccount).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	_, err = currentUser.
		Update().
		SetBankingData(updatedBankingData).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	return currentUser.ID, nil
}

func (r *mutationResolver) DeleteRestaurantOwnerProfile(ctx context.Context) (int, error) {
	kratosSessionUser := auth.ForContext(ctx)

	currentUser, err := r.EntClient.RestaurantOwner.
		Query().
		Where(restaurantowner.KratosID(kratosSessionUser.ID)).
		First(ctx)

	if err != nil {
		return -1, err
	}

	err = r.EntClient.RestaurantOwner.DeleteOne(currentUser).Exec(ctx)

	if err != nil {
		return -1, err
	}

	return currentUser.ID, nil
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.RegisterProductInput) (int, error) {
	//currentUser := auth.ForContext(ctx)

	//TODO: FALTA TAGS
	//Falta edge con restaurant
	restaurantRef, err := r.EntClient.Restaurant.
		Query().
		Where(restaurant.ID(input.RestaurantID)).
		First(ctx)

	if err != nil {
		return -1, err
	}
	tagEntities, err := GetOrCreateTagId(r.Resolver, input.Tags, ctx)

	newProduct, err := r.EntClient.Product.
		Create().
		SetName(input.Name).
		SetDescription(input.Description).
		SetCost(input.Cost).
		SetIsActive(input.Active).
		AddRestaurant(restaurantRef).
		AddTags(tagEntities...).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	return newProduct.ID, nil
}

func (r *mutationResolver) UpdateProduct(ctx context.Context, input model.UpdateProductInput) (int, error) {
	tagEntities, err := GetOrCreateTagId(r.Resolver, input.Tags, ctx)

	if err != nil {
		return -1, err
	}

	p, err := r.EntClient.Product.
		UpdateOneID(input.ID).
		SetName(input.Name).
		SetDescription(input.Description).
		SetCost(input.Cost).
		SetIsActive(input.Active).
		ClearTags().
		AddTags(tagEntities...).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	return p.ID, nil
}

func (r *mutationResolver) ToggleProductStatus(ctx context.Context, input int) (bool, error) {
	product, err := r.EntClient.Product.
		Query().
		Where(product.ID(input)).
		First(ctx)

	if err != nil {
		return false, err
	}

	updatedProduct, err := r.EntClient.Product.
		UpdateOneID(input).
		SetIsActive(!product.IsActive).
		Save(ctx)

	if err != nil {
		return false, err
	}

	return updatedProduct.IsActive, nil
}

func (r *mutationResolver) UploadProductPhoto(ctx context.Context, input model.UploadImageInput) ([]string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProductPhoto(ctx context.Context, input []string) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, input int) (int, error) {
	err := r.EntClient.Product.DeleteOneID(input).Exec(ctx)

	if err != nil {
		return -1, err
	}

	return input, nil
}

func (r *mutationResolver) UpdateRestaurant(ctx context.Context, input model.RegisterRestaurantInput) (int, error) {
	kratosSessionUser := auth.ForContext(ctx)

	currentUser, err := r.EntClient.RestaurantOwner.
		Query().
		Where(restaurantowner.KratosID(kratosSessionUser.ID)).
		WithBankingData().
		First(ctx)

	if err != nil {
		return -1, err
	}

	updatedAddress, err := currentUser.Edges.Restaurant.Edges.Address.
		Update().
		SetLatitude(input.Address.Latitude).
		SetLongitude(input.Address.Longitude).
		SetStreetLine(input.Address.StreetLine).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	tagEntities, err := GetOrCreateTagId(r.Resolver, input.Tags, ctx)
	if err != nil {
		return -1, err
	}
	updatedRestaurant, err := currentUser.Edges.Restaurant.
		Update().
		SetName(input.Name).
		SetDescription(input.Description).
		SetAddress(updatedAddress).
		ClearTags().
		AddTags(tagEntities...).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	_, err = currentUser.
		Update().
		SetRestaurant(updatedRestaurant).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	return currentUser.ID, nil
}

func (r *mutationResolver) UploadRestaurantPhoto(ctx context.Context, input model.UploadImageInput) ([]string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteRestuarantPhoto(ctx context.Context, input model.DeleteImageInput) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteRestaurant(ctx context.Context, input int) (int, error) {
	err := r.EntClient.Restaurant.DeleteOneID(input).Exec(ctx)

	if err != nil {
		return -1, err
	}

	return input, nil
}

func (r *mutationResolver) CreateProductRating(ctx context.Context, input model.RegisterRatingInput) (int, error) {
	kratosUser := auth.ForContext(ctx)

	currentCustomer, err := r.EntClient.Customer.
		Query().
		Where(customer.KratosID(kratosUser.ID)).
		First(ctx)
	if err != nil {
		return -1, nil
	}
	rating, err := r.EntClient.Rating.Create().
		SetProductID(input.ProductID).
		SetCustomer(currentCustomer).
		SetRating(input.Rating).
		SetComment(*input.Comment).
		Save(ctx)

	if err != nil {
		return -1, nil
	}

	return rating.ID, nil
}

func (r *mutationResolver) UpdateProductRating(ctx context.Context, input model.UpdateRatingInput) (int, error) {
	_, err := r.EntClient.Rating.UpdateOneID(input.ID).
		SetComment(*input.Comment).
		SetRating(input.Rating).
		Save(ctx)
	if err != nil {
		return -1, err
	}
	return input.ID, nil
}

func (r *mutationResolver) DeleteRating(ctx context.Context, input int) (int, error) {
	err := r.EntClient.Rating.DeleteOneID(input)
	if err != nil {
		return -1, nil
	}
	return input, nil
}

func (r *mutationResolver) UploadPhotoDemo(ctx context.Context, input model.UploadImageInput) ([]string, error) {
	// TODO: Why do we need to dereference to a variable?
	fileHandler := *r.FileHandler
	paths := make([]string, len(input.Files))
	for i, image := range input.Files {
		path, err := fileHandler.Upload(image)
		if err != nil {
			return nil, err
		}
		paths[i] = path
	}

	return paths, nil
}

func (r *mutationResolver) DeletePhotoDemo(ctx context.Context, input model.DeleteImageInput) ([]string, error) {
	fileHandler := *r.FileHandler
	for _, image := range input.FileNames {
		err := fileHandler.Delete(image)
		if err != nil {
			return nil, err
		}
	}

	return input.FileNames, nil
}

func (r *productResolver) Tags(ctx context.Context, obj *ent.Product) ([]string, error) {
	tags, err := r.EntClient.Product.QueryTags(obj).Select(tag.FieldName).Strings(ctx)
	return tags, ent.MaskNotFound(err)
}

func (r *productResolver) AverageRating(ctx context.Context, obj *ent.Product) (float64, error) {
	var avg float64
	err := r.EntClient.Product.QueryRatings(obj).
		GroupBy(rating.EdgeProduct).
		Aggregate(ent.Mean(rating.FieldRating)).
		Scan(ctx, &avg)
	return avg, ent.MaskNotFound(err)
}

func (r *productResolver) Ratings(ctx context.Context, obj *ent.Product) ([]*ent.Rating, error) {
	ratings, err := r.EntClient.Product.QueryRatings(obj).All(ctx)
	return ratings, ent.MaskNotFound(err)
}

func (r *productResolver) Restaurant(ctx context.Context, obj *ent.Product) (*ent.Restaurant, error) {
	restaurant, err := r.EntClient.Product.QueryRestaurant(obj).First(ctx)
	return restaurant, ent.MaskNotFound(err)
}

func (r *queryResolver) GetCurrentCustomer(ctx context.Context) (*ent.Customer, error) {
	kratosUser := auth.ForContext(ctx)

	currentCustomer, err := r.EntClient.Customer.
		Query().
		Where(customer.KratosID(kratosUser.ID)).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return currentCustomer, nil
}

func (r *queryResolver) GetCurrentRestaurantOwner(ctx context.Context) (*ent.RestaurantOwner, error) {
	kratosUser := auth.ForContext(ctx)

	currentRestaurantOwner, err := r.EntClient.RestaurantOwner.
		Query().
		Where(restaurantowner.KratosID(kratosUser.ID)).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return currentRestaurantOwner, nil
}

func (r *queryResolver) GetRestaurantByID(ctx context.Context, input int) (*ent.Restaurant, error) {
	restaurant, err := r.EntClient.Restaurant.
		Query().
		Where(restaurant.ID(input)).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return restaurant, nil
}

func (r *queryResolver) GetClosestRestaurants(ctx context.Context, input *int) ([]*model.RestaurantSearchResult, error) {
	// TODO: Add comments
	kratosUser := auth.ForContext(ctx)
	currentUser, err := r.EntClient.Customer.
		Query().
		Where(customer.KratosID(kratosUser.ID)).
		WithAddress().
		First(ctx)

	log.Info().Msg(currentUser.Name)
	if err != nil {
		return nil, err
	}
	columnsTmp := sql.Table(restaurant.Table).Columns(restaurant.Columns...)
	var myRegex = "s`\\."
	var re = regexp.MustCompile(myRegex)
	for i, v := range columnsTmp {
		vNew := re.ReplaceAllString(v, "`.")
		columnsTmp[i] = fmt.Sprintf(`%s AS "%s" `, v, vNew)
	}
	columns := fmt.Sprintf(strings.Join(columnsTmp, ","))
	query := `select geom from customers join addresses a on a.id = customers.customer_address where kratos_id = $1`
	var geom string
	err = r.DBClient.Get(&geom, query, kratosUser.ID)
	if err != nil {
		return nil, err
	}
	distance := "distance"
	limit := 20
	query = fmt.Sprintf(
		`SELECT %s,
		round((ST_DISTANCESPHERE(A.geom,'%s')/1000)::numeric,2) as %s FROM %s JOIN %s A on A.%s = %s ORDER BY A.geom <-> '%s' LIMIT %d`,
		columns, geom, distance, restaurant.Table,
		address.Table, address.FieldID,
		sql.Table(restaurant.Table).C(restaurant.AddressColumn), geom, limit,
	)
	query = strings.ReplaceAll(query, "`", "")
	log.Info().Msg(query)
	//var restaurants []struct{
	//	ent.Restaurant
	//	Distance string
	//}
	var restaurants []*model.RestaurantSearchResult
	err = r.DBClient.Select(&restaurants, query)
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (r *queryResolver) GetProductsByRestaurantID(ctx context.Context, input model.ProductsFilterByRestaurantInput) ([]*ent.Product, error) {
	//No jala
	restaurant, err := r.EntClient.Restaurant.
		Query().
		Where(restaurant.ID(input.RestaurantID)).
		WithProducts().
		First(ctx)

	if err != nil {
		return nil, err
	}

	return restaurant.Edges.Products, nil
}

func (r *queryResolver) GetProductByID(ctx context.Context, input int) (*ent.Product, error) {
	res, err := r.EntClient.Product.Get(ctx, input)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *queryResolver) GetProductsByAllFields(ctx context.Context, input model.ProductsByAllFieldsInput) ([]*ent.Product, error) {
	var products []ent.Product
	//restaurants, err:= r.EntClient.Restaurant.
	kratosUser := auth.ForContext(ctx)
	address, err := r.EntClient.Customer.
		Query().
		Where(customer.KratosID(kratosUser.ID)).
		QueryAddress().
		First(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(address.Geom)

	//var restaurants []ent.Restaurant
	//err = r.EntClient.Restaurant.
	//	Query().
	//	//QueryAddress().
	//	//WithAddress().
	//	//Where().
	//	Where(OrderByDistanceP(restaurant.Table, *address.Geom)).
	//	//GroupBy(restaurant.FieldID).
	//	//Aggregate(SelectDistanceP("distance")).
	//	Select(fmt.Sprintf(`st_distancesphere(geom, '%s') as "distance"`, *address.Geom), GetColumns(restaurant.Table, restaurant.Columns)...).
	//	UnsafeScan(ctx, &restaurants)
	//fmt.Println(restaurants)
	//if err != nil {
	//	return nil, err
	//}

	err = r.EntClient.Product.
		Query().
		Where(OrderByDistanceP(product.Table, *address.Geom)).
		Select(fmt.Sprintf(`st_distancesphere(geom, '%s') as "distance"`, *address.Geom), GetColumns(product.Table, product.Columns)...).
		UnsafeScan(ctx, &products)
	fmt.Println(products)
	if err != nil {
		return nil, err
	}
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AutoCompleteTag(ctx context.Context, input string) ([]string, error) {
	jsonObj := gabs.New()
	_, _ = jsonObj.SetP(1, "suggest.suggestion.completion.fuzzy.fuzziness")
	_, _ = jsonObj.SetP("autocomplete", "suggest.suggestion.completion.field")
	_, _ = jsonObj.SetP(input, "suggest.suggestion.prefix")
	body := jsonObj.String()
	results, err := r.ElasticClient.Search(
		r.ElasticClient.Search.WithBody(strings.NewReader(body)),
	)
	if err != nil {
		return nil, err
	}
	defer results.Body.Close()
	parsed, err := gabs.ParseJSONBuffer(results.Body)
	if err != nil {
		return nil, err
	}
	tags := make([]string, 0)
	for _, child := range parsed.Path("suggest.suggestion.0.options").Children() {
		source := child.S("_source")
		name := source.S("name").Data().(string)
		tags = append(tags, name)
		// fmt.Printf("key: %v, value: %v\n", key, child.Data().(float64))
	}
	return tags, nil
}

func (r *ratingResolver) Product(ctx context.Context, obj *ent.Rating) (*ent.Product, error) {
	product, err := r.EntClient.Rating.QueryProduct(obj).First(ctx)
	return product, ent.MaskNotFound(err)
}

func (r *ratingResolver) Customer(ctx context.Context, obj *ent.Rating) (*ent.Customer, error) {
	customer, err := r.EntClient.Rating.QueryCustomer(obj).First(ctx)
	return customer, ent.MaskNotFound(err)
}

func (r *restaurantResolver) Address(ctx context.Context, obj *ent.Restaurant) (*ent.Address, error) {
	address, err := r.EntClient.Restaurant.QueryAddress(obj).First(ctx)
	return address, ent.MaskNotFound(err)
}

func (r *restaurantResolver) Tags(ctx context.Context, obj *ent.Restaurant) ([]string, error) {
	tags, err := r.EntClient.Restaurant.QueryTags(obj).Select(tag.FieldName).Strings(ctx)
	return tags, ent.MaskNotFound(err)
}

func (r *restaurantResolver) Products(ctx context.Context, obj *ent.Restaurant) ([]*ent.Product, error) {
	products, err := r.EntClient.Restaurant.QueryProducts(obj).All(ctx)
	return products, ent.MaskNotFound(err)
}

func (r *restaurantResolver) RestaurantOwner(ctx context.Context, obj *ent.Restaurant) (*ent.RestaurantOwner, error) {
	restaurantOwner, err := r.EntClient.Restaurant.QueryOwner(obj).First(ctx)
	return restaurantOwner, ent.MaskNotFound(err)
}

func (r *restaurantOwnerResolver) Banking(ctx context.Context, obj *ent.RestaurantOwner) (*ent.BankingData, error) {
	banking, err := r.EntClient.RestaurantOwner.QueryBankingData(obj).First(ctx)
	return banking, ent.MaskNotFound(err)
}

func (r *restaurantOwnerResolver) Restaurant(ctx context.Context, obj *ent.RestaurantOwner) (*ent.Restaurant, error) {
	restaurant, err := r.EntClient.RestaurantOwner.QueryRestaurant(obj).First(ctx)
	return restaurant, ent.MaskNotFound(err)
}

// Customer returns generated.CustomerResolver implementation.
func (r *Resolver) Customer() generated.CustomerResolver { return &customerResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Rating returns generated.RatingResolver implementation.
func (r *Resolver) Rating() generated.RatingResolver { return &ratingResolver{r} }

// Restaurant returns generated.RestaurantResolver implementation.
func (r *Resolver) Restaurant() generated.RestaurantResolver { return &restaurantResolver{r} }

// RestaurantOwner returns generated.RestaurantOwnerResolver implementation.
func (r *Resolver) RestaurantOwner() generated.RestaurantOwnerResolver {
	return &restaurantOwnerResolver{r}
}

type customerResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type ratingResolver struct{ *Resolver }
type restaurantResolver struct{ *Resolver }
type restaurantOwnerResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *productResolver) Distance(ctx context.Context, obj *ent.Product) (float64, error) {
	panic(fmt.Errorf("not implemented"))
}
