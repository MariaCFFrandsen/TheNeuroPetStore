package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"theNeuroPetStore/api/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "docker"
	dbname   = "postgres"
)

func ConnectToDatabase() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

func GetPets(db *sql.DB) (*models.Pets, error) {
	rows, err := db.Query("SELECT * FROM pets")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	var pets []*models.Pet
	strings, err := rows.Columns()
	fmt.Println(strings[0])
	for rows.Next() {
		var pet *models.Pet
		if err = rows.Scan(&pet.ID, &pet.Name, &pet.ImageURL, &pet.Available, &pet.Qualities); err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}
	result := &models.Pets{
		pets[0],
		pets[1],
		pets[2],
		pets[3],
	}
	fmt.Println(pets[0].Name)

	return result, err
}
