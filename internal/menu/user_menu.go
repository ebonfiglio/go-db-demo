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
			"Get all Users",
			"Back",
		})

		switch choice {
		case "1":
			createUserCommand(dbConn)
			fmt.Println("Success!")
		case "2":
			updateUserCommand(dbConn)
			fmt.Println("Success!")
		case "3":
			getUserCommand(dbConn)
		case "4":
			fmt.Println("Deleting User...")
		case "5":
			getAllUsersCommand(dbConn)
		case "6":
			return
		}
	}
}

func createUserCommand(dbConn *sqlx.DB) {
	newUserValues := getEntityInput()
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

func getAllUsersCommand(dbConn *sqlx.DB) {
	users, err := service.GetAllUsers(dbConn)
	if err != nil {
		fmt.Println(err)
	}

	for _, u := range users {
		fmt.Println(u.ID, u.Name, u.JobID, u.OrganizationID)
	}
}

func getUserCommand(dbConn *sqlx.DB) {
	id := getId()
	if id == 0 {
		return
	}
	user, err := service.GetUser(id, *dbConn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user.ID, user.Name, user.JobID, user.OrganizationID)

}

func updateUserCommand(dbConn *sqlx.DB) {
	newUserValues := getEntityInput()
	user, err := domain.JsonToUser(newUserValues)
	if err != nil {
		fmt.Println(err)
	}
	user, err = service.UpdateUser(user, dbConn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("User ID: ", user.ID)
	fmt.Println("User Name: ", user.Name)
	fmt.Println("User Job ID: ", user.JobID)
	fmt.Println("User Org ID: ", user.OrganizationID)
}
