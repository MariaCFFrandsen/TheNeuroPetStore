package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"theNeuroPetStore/api/models"
	"theNeuroPetStore/api/restapi/operations/pets"
)

type petsList struct{}

func NewPetsListHandler() pets.PetsListHandler {
	return &petsList{}
}

func (impl *petsList) Handle(pets.PetsListParams) middleware.Responder {
	responseVal := &models.Pets{
		&models.Pet{
			ID:       1,
			ImageURL: "https://images.pexels.com/photos/45201/kitty-cat-kitten-pet-45201.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2",
			Name:     "Cuddles",
			Species:  "Cat",
			Qualities: []string{
				"cuddly",
				"sweet",
				"indoors",
			},
		},
	}
	return pets.NewPetsListOK().WithPayload(*responseVal)
}
