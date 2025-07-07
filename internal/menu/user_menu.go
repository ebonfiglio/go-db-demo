package menu

import (
	"fmt"
	"go-db-demo/internal/domain"
	"go-db-demo/internal/service"

	"github.com/jmoiron/sqlx"
)

func UserMenu(dbConn *sqlx.DB) {
	for {
		choice := DisplayMenuOptions([]string{
			"Create User",
			"Update User",
			"Lookup User",
			"Delete User",
			"Back",
		})

		switch choice {
		case "1":
			createUserCommand(dbConn)
			fmt.Println("Success!")
		case "2":
			fmt.Println("Updating User...")
		case "3":
			fmt.Println("Looking up User...")
		case "4":
			fmt.Println("Deleting User...")
		case "5":
			return
		}
	}
}

func createUserCommand(dbConn *sqlx.DB) {
	newUserValues := getNewEntityInput()
	user, err := domain.JsonToUser(newUserValues)
	if err != nil {
		fmt.Println(err)
	}
	user, err = service.CreateUser(user, dbConn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("User ID: ", user.ID)
}
