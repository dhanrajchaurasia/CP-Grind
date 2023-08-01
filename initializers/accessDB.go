package initializers

import (
	"errors"
	"log"

	"github.com/dhanrajchaurasia/CP-GRIND/helpers"
	"github.com/dhanrajchaurasia/CP-GRIND/models"
)

func CreateNewUser(user models.User) error {
	var Users []models.User
	DB.Raw("SELECT * FROM users WHERE username = $1", user.Username).Scan(&Users)
	if len(Users) > 0 {
		return errors.New("Username Already Exists!")
	}
	DB.Raw("SELECT * FROM users WHERE email = $1", user.Email).Scan(&Users)
	if len(Users) > 0 {
		return errors.New("Email Already Exists!")
	}
	user.Password, err = helpers.Encrypt(user.Password)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if user.AuthToken, err = helpers.CreateToken(user.Username); err != nil {
		log.Fatal(err)
		return err
	}
	return DB.Create(&user).Error
}

func IsUserPresent(username string, password string) (string, error) {
	var Users []models.User
	DB.Raw("SELECT * FROM users WHERE username = $1", username).Scan(&Users)
	if len(Users) != 1 {
		return "", errors.New("User doesn't exist!")
	}
	pass, err := helpers.Decrypt(Users[0].Password)
	if err != nil {
		return "", err
	}
	if pass != password {
		return "", errors.New("Password doesn't match!")
	}
	return Users[0].AuthToken, nil
}

func IsValidToken(token string) bool {
	var tokens []models.User
	DB.Raw("SELECT auth_token FROM users WHERE auth_token = $1", token).Scan(&tokens)
	return len(tokens) > 0
}
