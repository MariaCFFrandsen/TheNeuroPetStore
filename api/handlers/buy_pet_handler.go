package handlers

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"theNeuroPetStore/api/restapi/operations/pets"
	"theNeuroPetStore/internal/sql"
)

type buyPet struct {
	view *sql.View
}

func NewBuyPet(view *sql.View) pets.BuyPetHandler {
	return &buyPet{
		view: view,
	}
}

func (b *buyPet) Handle(params pets.BuyPetParams) middleware.Responder {
	ctx := context.Background()
	err := b.view.BuyPet(ctx, int32(params.ID))
	if err != nil {
		return pets.NewBuyPetInternalServerError()
	}
	return pets.NewPetsListOK()
}
