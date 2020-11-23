package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"foodworks.ml/m/internal/auth"
	"foodworks.ml/m/internal/generated/ent"
	"foodworks.ml/m/internal/generated/ent/customer"
	"foodworks.ml/m/internal/generated/ent/imagepath"
	"foodworks.ml/m/internal/generated/ent/order"
	"foodworks.ml/m/internal/generated/ent/product"
	"foodworks.ml/m/internal/generated/ent/rating"
	"foodworks.ml/m/internal/generated/ent/restaurant"
	"foodworks.ml/m/internal/generated/ent/restaurantowner"
	"foodworks.ml/m/internal/generated/ent/tag"
	generated "foodworks.ml/m/internal/generated/graphql"
	"foodworks.ml/m/internal/generated/graphql/model"
	gabs "github.com/Jeffail/gabs/v2"
)

func (r *customerResolver) Address(ctx context.Context, obj *ent.Customer) (*ent.Address, error) {
	address, err := r.EntClient.Customer.QueryAddress(obj).First(ctx)
	return address, ent.MaskNotFound(err)
}

func (r *customerResolver) RatedProducts(ctx context.Context, obj *ent.Customer) ([]*ent.Rating, error) {
	ratings, err := r.EntClient.Customer.QueryRatings(obj).All(ctx)
	return ratings, ent.MaskNotFound(err)
}

func (r *customerResolver) Orders(ctx context.Context, obj *ent.Customer) ([]*ent.Order, error) {
	kratosUser := auth.ForContext(ctx)

	orders, err := r.EntClient.Customer.
		Query().
		Where(customer.KratosID(kratosUser.ID)).
		QueryOrders().
		All(ctx)

	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *customerResolver) PaymentMethod(ctx context.Context, obj *ent.Customer) (*ent.PaymentMethod, error) {
	product, err := obj.QueryPaymentMethod().First(ctx)
	return product, ent.MaskNotFound(err)
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
		SetLastName(input.LastName).
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
		SetLastName(input.LastName).
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

func (r *mutationResolver) UpdateCustomerPaymentMethod(ctx context.Context, input *model.PaymentMethodInput) (int, error) {
	kratosSessionUser := auth.ForContext(ctx)

	currentUser, err := r.EntClient.Customer.
		Query().
		Where(customer.KratosID(kratosSessionUser.ID)).
		WithPaymentMethod().
		First(ctx)

	if err != nil {
		return -1, err
	}
	var paymentMethod *ent.PaymentMethod
	if len(currentUser.Edges.PaymentMethod) == 0 {
		paymentMethod, err = r.EntClient.PaymentMethod.Create().
			SetData(input.Data).
			Save(ctx)
		if err != nil {
			return -1, err
		}
		_, err = currentUser.Update().AddPaymentMethod(paymentMethod).Save(ctx)
	} else {
		paymentMethod = currentUser.Edges.PaymentMethod[0]
		_, err = paymentMethod.Update().
			SetData(input.Data).
			Save(ctx)
	}
	if err != nil {
		return -1, err
	}
	return paymentMethod.ID, nil
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

	if err != nil {
		return -1, err
	}

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
		SetLastName(input.LastName).
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
		SetLastName(input.LastName).
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

	if err != nil {
		return -1, err
	}

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

func (r *mutationResolver) UploadProductPhoto(ctx context.Context, input model.UploadProductImageInput) (string, error) {
	fileHandler := *r.FileHandler
	path, err := fileHandler.Upload(&input.File)
	if err != nil {
		return "", err
	}
	_, err = r.EntClient.ImagePath.Create().SetProductID(input.ID).SetPath(path).Save(ctx)
	if err != nil {
		return "", err
	}
	return path, nil
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

	currentUserRestaurant, err := r.EntClient.RestaurantOwner.
		Query().
		Where(restaurantowner.KratosID(kratosSessionUser.ID)).
		QueryRestaurant().
		WithAddress().
		First(ctx)

	if err != nil {
		return -1, err
	}

	_, err = currentUserRestaurant.Edges.Address.
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

	_, err = currentUserRestaurant.
		Update().
		SetName(input.Name).
		SetDescription(input.Description).
		ClearTags().
		AddTags(tagEntities...).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	return currentUserRestaurant.ID, nil
}

func (r *mutationResolver) UploadRestaurantPhoto(ctx context.Context, input model.UploadRestaurantImageInput) (string, error) {
	fileHandler := *r.FileHandler
	path, err := fileHandler.Upload(&input.File)
	if err != nil {
		return "", err
	}
	_, err = r.EntClient.ImagePath.Create().SetRestaurantID(input.ID).SetPath(path).Save(ctx)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (r *mutationResolver) DeleteRestaurant(ctx context.Context, input int) (int, error) {
	err := r.EntClient.Restaurant.DeleteOneID(input).Exec(ctx)

	if err != nil {
		return -1, err
	}

	return input, nil
}

func (r *mutationResolver) DeletePhoto(ctx context.Context, input model.DeleteImageInput) (int, error) {
	fileHandler := *r.FileHandler
	err := fileHandler.Delete(input.FilePath)
	if err != nil {
		return -1, err
	}
	_, err = r.EntClient.ImagePath.Delete().Where(imagepath.PathEQ(input.FilePath)).Exec(ctx)
	if err != nil {
		return -1, err
	}
	return -1, nil
}

func (r *mutationResolver) CreateProductRating(ctx context.Context, input model.RegisterRatingInput) (int, error) {
	kratosUser := auth.ForContext(ctx)

	currentCustomer, err := r.EntClient.Customer.
		Query().
		Where(customer.KratosID(kratosUser.ID)).
		First(ctx)
	if err != nil {
		return -1, err
	}
	rating, err := r.EntClient.Rating.Create().
		SetProductID(input.ProductID).
		SetCustomer(currentCustomer).
		SetRating(input.Rating).
		SetComment(*input.Comment).
		Save(ctx)

	if err != nil {
		return -1, err
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
	err := r.EntClient.Rating.DeleteOneID(input).Exec(ctx)
	if err != nil {
		return -1, err
	}
	return input, nil
}

func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.RegisterOrderInput) (int, error) {
	kratosUser := auth.ForContext(ctx)

	currentCustomer, err := r.EntClient.Customer.
		Query().
		Where(customer.KratosID(kratosUser.ID)).
		First(ctx)

	if err != nil {
		return -1, err
	}

	order, err := r.EntClient.Order.Create().
		SetQuantity(input.Quantity).
		SetOrderState(model.OrderStatePendingPayment.String()).
		SetUpdatedAt(time.Now().UTC()).
		AddCustomer(currentCustomer).
		AddProductIDs(input.ProductID).
		Save(ctx)

	if err != nil {
		return -1, err
	}

	return order.ID, nil
}

func (r *mutationResolver) UpdateOrder(ctx context.Context, input *model.UpdateOrderInput) (int, error) {
	kratosUser := auth.ForContext(ctx)

	currentCustomer, err := r.EntClient.Customer.
		Query().
		Where(customer.KratosID(kratosUser.ID)).
		First(ctx)

	if ent.IsNotFound(err) {
		currentOwner, err := r.EntClient.RestaurantOwner.
			Query().
			Where(restaurantowner.KratosID(kratosUser.ID)).
			First(ctx)
		if err != nil {
			return -1, err
		}
		return r.updateOrderAsRestaurantOwner(ctx, input, currentOwner)
	}

	if err != nil {
		return -1, err
	}

	_, err = r.EntClient.Order.Update().
		Where(
			order.HasCustomerWith(customer.ID(currentCustomer.ID)),
			order.IDEQ(input.OrderID),
		).
		SetOrderState(input.OrderState.String()).
		SetUpdatedAt(time.Now().UTC()).
		Save(ctx)

	if err != nil {
		return -1, err
	}
	return input.OrderID, nil
}

func (r *orderResolver) Product(ctx context.Context, obj *ent.Order) (*ent.Product, error) {
	product, err := obj.QueryProduct().First(ctx)
	return product, ent.MaskNotFound(err)
}

func (r *orderResolver) Customer(ctx context.Context, obj *ent.Order) (*ent.Customer, error) {
	customer, err := obj.QueryCustomer().First(ctx)
	return customer, ent.MaskNotFound(err)
}

func (r *orderResolver) OrderState(ctx context.Context, obj *ent.Order) (model.OrderState, error) {
	var model model.OrderState
	err := model.UnmarshalGQL(obj.OrderState)
	return model, err
}

func (r *orderResolver) UpdatedAt(ctx context.Context, obj *ent.Order) (int64, error) {
	return obj.UpdatedAt.Unix(), nil
}

func (r *productResolver) Tags(ctx context.Context, obj *ent.Product) ([]string, error) {
	tags, err := r.EntClient.Product.QueryTags(obj).Select(tag.FieldName).Strings(ctx)
	return tags, ent.MaskNotFound(err)
}

func (r *productResolver) Active(ctx context.Context, obj *ent.Product) (bool, error) {
	return obj.IsActive, nil
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

func (r *productResolver) Image(ctx context.Context, obj *ent.Product) (string, error) {
	image, err := r.EntClient.Product.QueryImages(obj).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return "", nil
		}
		return "", err
	}
	return image.Path, ent.MaskNotFound(err)
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

func (r *queryResolver) GetClosestRestaurants(ctx context.Context) ([]*ent.Restaurant, error) {
	// TODO: Add comments
	kratosUser := auth.ForContext(ctx)
	address, err := r.EntClient.Customer.
		Query().
		Where(customer.KratosID(kratosUser.ID)).
		QueryAddress().
		First(ctx)

	if err != nil {
		return nil, err
	}

	var restaurants []*ent.Restaurant
	err = r.EntClient.Restaurant.
		Query().
		Where(OrderByDistanceP(restaurant.Table, *address.Geom)).
		Select(fmt.Sprintf(`round((st_distancesphere(geom, '%s')/1000)::numeric,2) as "distance"`, *address.Geom), GetColumns(restaurant.Table, restaurant.Columns)...).
		UnsafeScan(ctx, &restaurants)
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

func (r *queryResolver) GetProductByTag(ctx context.Context, input string) ([]*ent.Product, error) {
	products, err := r.EntClient.Tag.Query().Where(tag.Name(input)).QueryProduct().All(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *queryResolver) SearchProductsAndRestaurants(ctx context.Context, input model.ProductsByAllFieldsInput) (*model.GlobalSearchResult, error) {
	var result model.GlobalSearchResult
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

	var restaurants []*ent.Restaurant
	err = r.EntClient.Restaurant.
		Query().
		Where(WhereTextMatch(restaurant.Table, input.SearchString)).
		Where(OrderByDistanceP(restaurant.Table, *address.Geom)).
		Select(fmt.Sprintf(`round((st_distancesphere(geom, '%s')/1000)::numeric,2) as "distance"`, *address.Geom), GetColumns(restaurant.Table, restaurant.Columns)...).
		UnsafeScan(ctx, &restaurants)
	if err != nil {
		return nil, err
	}

	var products []*ent.Product
	err = r.EntClient.Product.
		Query().
		Where(WhereTextMatch(product.Table, input.SearchString)).
		Where(OrderByDistanceP(product.Table, *address.Geom)).
		Select(fmt.Sprintf(`round((st_distancesphere(geom, '%s')/1000)::numeric,2) as "distance"`, *address.Geom), GetColumns(product.Table, product.Columns)...).
		UnsafeScan(ctx, &products)
	if err != nil {
		return nil, err
	}
	result.Products = RemoveDuplicateProducts(products)
	result.Restaurants = RemoveDuplicateRestaurant(restaurants)
	return &result, nil
}

func (r *queryResolver) GetFeed(ctx context.Context) ([]*model.FeedItem, error) {
	kratosUser := auth.ForContext(ctx)

	currentUser, err := r.EntClient.Customer.
		Query().
		Where(customer.KratosID(kratosUser.ID)).
		First(ctx)

	if err != nil {
		return nil, err
	}

	feeds := make([]*model.FeedItem, 2)

	rec, err := r.Recommender.GetUserRecommendations(currentUser.ID)

	if err != nil {
		return nil, err
	}
	var recProducts []*ent.Product
	if len(rec) > 0 {
		recProducts, err = r.EntClient.Product.Query().Where(product.IDIn(rec...)).All(ctx)
	} else {
		recProducts, err = r.EntClient.Product.Query().Limit(5).All(ctx)
	}
	if err != nil {
		return nil, err
	}
	feedItems := make([]model.FeedCard, len(recProducts))
	for i, recProduct := range recProducts {
		feedItems[i] = model.FeedCard(recProduct)
	}
	var recommended = model.FeedItem{Name: "Recommended Products", Cards: feedItems}
	feeds[0] = &recommended

	restaurants, err := r.GetClosestRestaurants(ctx)
	if err != nil {
		return nil, err
	}
	feedRestaurantItems := make([]model.FeedCard, len(restaurants))
	for i, recRestaurant := range restaurants {
		feedRestaurantItems[i] = model.FeedCard(recRestaurant)
	}
	var closestRestaurants = model.FeedItem{Name: "Closest restaurants", Cards: feedRestaurantItems}
	feeds[1] = &closestRestaurants
	return feeds, nil
}

func (r *queryResolver) GetCustomerOrders(ctx context.Context) ([]*ent.Order, error) {
	kratosUser := auth.ForContext(ctx)

	orders, err := r.EntClient.Customer.
		Query().
		Where(customer.KratosID(kratosUser.ID)).
		QueryOrders().
		All(ctx)

	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *queryResolver) GetRestaurantOrders(ctx context.Context) ([]*ent.Order, error) {
	kratosUser := auth.ForContext(ctx)

	orders, err := r.EntClient.RestaurantOwner.
		Query().
		Where(restaurantowner.KratosID(kratosUser.ID)).
		QueryRestaurant().
		QueryProducts().
		QueryOrders().
		All(ctx)

	if err != nil {
		return nil, err
	}
	return orders, nil
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

func (r *restaurantResolver) Orders(ctx context.Context, obj *ent.Restaurant) ([]*ent.Order, error) {
	kratosUser := auth.ForContext(ctx)

	orders, err := r.EntClient.RestaurantOwner.
		Query().
		Where(restaurantowner.KratosID(kratosUser.ID)).
		QueryRestaurant().
		QueryProducts().
		QueryOrders().
		All(ctx)

	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *restaurantResolver) RestaurantOwner(ctx context.Context, obj *ent.Restaurant) (*ent.RestaurantOwner, error) {
	restaurantOwner, err := r.EntClient.Restaurant.QueryOwner(obj).First(ctx)
	return restaurantOwner, ent.MaskNotFound(err)
}

func (r *restaurantResolver) Image(ctx context.Context, obj *ent.Restaurant) (string, error) {
	image, err := r.EntClient.Restaurant.QueryImages(obj).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return "", nil
		}
		return "", err
	}
	return image.Path, ent.MaskNotFound(err)
}

func (r *restaurantOwnerResolver) Banking(ctx context.Context, obj *ent.RestaurantOwner) (*ent.BankingData, error) {
	banking, err := r.EntClient.RestaurantOwner.QueryBankingData(obj).First(ctx)
	return banking, ent.MaskNotFound(err)
}

func (r *restaurantOwnerResolver) Restaurant(ctx context.Context, obj *ent.RestaurantOwner) (*ent.Restaurant, error) {
	restaurant, err := r.EntClient.RestaurantOwner.QueryRestaurant(obj).First(ctx)
	return restaurant, ent.MaskNotFound(err)
}

func (r *subscriptionResolver) OnOrderStateChange(ctx context.Context) (<-chan int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Customer returns generated.CustomerResolver implementation.
func (r *Resolver) Customer() generated.CustomerResolver { return &customerResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Order returns generated.OrderResolver implementation.
func (r *Resolver) Order() generated.OrderResolver { return &orderResolver{r} }

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

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type customerResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type orderResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type ratingResolver struct{ *Resolver }
type restaurantResolver struct{ *Resolver }
type restaurantOwnerResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) DeleteProductPhoto(ctx context.Context, input model.DeleteImageInput) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) DeleteRestaurantPhoto(ctx context.Context, input model.DeleteImageInput) (int, error) {
	fileHandler := *r.FileHandler
	err := fileHandler.Delete(input.FilePath)
	if err != nil {
		return -1, err
	}
	_, err = r.EntClient.ImagePath.Delete().Where(imagepath.PathEQ(input.FilePath)).Exec(ctx)
	if err != nil {
		return -1, err
	}
	return -1, nil
}
func (r *mutationResolver) DeleteRestuarantPhoto(ctx context.Context, input model.DeleteImageInput) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
