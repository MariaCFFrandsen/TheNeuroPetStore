package handlers

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"theNeuroPetStore/api/restapi/operations/pets"
	"theNeuroPetStore/internal/sql"
)

type petsList struct {
	view *sql.View
}

func NewPetsListHandler(view *sql.View) pets.PetsListHandler {
	return &petsList{
		view: view,
	}
}

func (impl *petsList) Handle(pets.PetsListParams) middleware.Responder {
	context := context.Background()
	getPets, _ := impl.view.GetPets(context)
	return pets.NewPetsListOK().WithPayload(getPets)
}
