package repositories

import (
	"fmt"
	"time"

	"github.com/rhidayat1980/fitness-api/models"
	"github.com/rhidayat1980/fitness-api/storage"
)

func CreateUser(user models.User) (models.User, error) {
	fmt.Println("Creating user...")
	db := storage.GetDB()
	sqlStatement := `
		INSERT INTO users (name, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return models.User{}, err
	}
	user.ID = id
	return user, nil
}

func UpdateUser(user models.User, id int) (models.User, error) {
	db := storage.GetDB()
	sqlStatement := `
		UPDATE users
		SET name = $2, email = $3, password = $4, updated_at = $5
		WHERE id = $1
		RETURNING id`
	err := db.QueryRow(sqlStatement, id, user.Name, user.Email, user.Password, time.Now()).Scan(&id)
	if err != nil {
		return models.User{}, err
	}
	user.ID = id
	return user, nil
}
