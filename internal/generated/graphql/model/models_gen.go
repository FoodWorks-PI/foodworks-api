// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"foodworks.ml/m/internal/generated/ent"
	"github.com/99designs/gqlgen/graphql"
)

type DeleteImageInput struct {
	FileNames []string `json:"fileNames"`
}

type ProductsByAllFieldsInput struct {
	SearchString        string                     `json:"searchString"`
	ProductFilterConfig *ProductsFilterConfigInput `json:"productFilterConfig"`
}

type ProductsFilterByRestaurantInput struct {
	RestaurantID        int                        `json:"restaurantID"`
	ProductFilterConfig *ProductsFilterConfigInput `json:"productFilterConfig"`
}

type ProductsFilterByTagInput struct {
	Tag                 int                        `json:"tag"`
	ProductFilterConfig *ProductsFilterConfigInput `json:"productFilterConfig"`
}

type ProductsFilterConfigInput struct {
	IncludeInactive *bool `json:"includeInactive"`
}

type RegisterAddressInput struct {
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	StreetLine string  `json:"streetLine"`
}

type RegisterBankingInput struct {
	BankAccount string `json:"bankAccount"`
}

type RegisterCustomerInput struct {
	Name     string                `json:"name"`
	LastName string                `json:"lastName"`
	Phone    string                `json:"phone"`
	Address  *RegisterAddressInput `json:"address"`
}

type RegisterProductInput struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Tags         []string `json:"tags"`
	Cost         int      `json:"cost"`
	Active       bool     `json:"active"`
	RestaurantID int      `json:"restaurantID"`
}

type RegisterRatingInput struct {
	ProductID int     `json:"productID"`
	Rating    int     `json:"rating"`
	Comment   *string `json:"comment"`
}

type RegisterRestaurantInput struct {
	Name        string                `json:"name"`
	Address     *RegisterAddressInput `json:"address"`
	Description string                `json:"description"`
	Tags        []string              `json:"tags"`
}

type RegisterRestaurantOwnerInput struct {
	Name       string                   `json:"name"`
	LastName   string                   `json:"lastName"`
	Phone      string                   `json:"phone"`
	Banking    *RegisterBankingInput    `json:"banking"`
	Restaurant *RegisterRestaurantInput `json:"restaurant"`
}

type RestaurantSearchResult struct {
	Restaurant *ent.Restaurant `json:"restaurant"`
	Distance   float64         `json:"distance"`
}

type UpdateCustomerInput struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Phone    string `json:"phone"`
}

type UpdateProductInput struct {
	ProductID   int      `json:"productID"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Cost        int      `json:"cost"`
	Active      bool     `json:"active"`
}

type UpdateRestaurantOwnerInput struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Phone    string `json:"phone"`
}

type UploadImageInput struct {
	Files []*graphql.Upload `json:"files"`
}
