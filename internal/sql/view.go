package sql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"theNeuroPetStore/api/models"
	"theNeuroPetStore/internal/sql/generated"
)

type View struct {
	conn    *sql.DB
	queries *generated.Queries
}

func NewView(conn *sql.DB) *View {
	return &View{
		conn:    conn,
		queries: generated.New(conn),
	}
}

func (view *View) GetPets(ctx context.Context) (models.Pets, error) {
	pets, err := view.queries.GetPets(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}

	result := ConvertToSwaggerModel(pets)
	return result, err
}

func ConvertToSwaggerModel(pets []generated.Pet) models.Pets {
	var swaggerpets models.Pets
	for _, pet := range pets {
		fmt.Println(pet.ID)
		swaggerpet := models.Pet{}
		swaggerpet.ID = int64(pet.ID)
		swaggerpet.Name = pet.Name
		swaggerpet.ImageURL = pet.Imageurl
		swaggerpet.Qualities = splitByComma(pet.Qualities)
		swaggerpet.Species = pet.Species
		swaggerpet.Available = pet.Available
		swaggerpets = append(swaggerpets, &swaggerpet)
	}
	return swaggerpets
}

func splitByComma(s string) []string {
	return strings.Split(s, ",")
}
