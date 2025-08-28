package main

import (
	"fmt"
	"go-db-demo/internal/db"
	"go-db-demo/internal/menu"
	"go-db-demo/internal/service"
)

var commands = map[string]string{
	"1":  "Create Org",
	"2":  "Update Org",
	"3":  "Lookup Org",
	"4":  "Delete Org",
	"5":  "Create Org",
	"6":  "Update Org",
	"7":  "Lookup Org",
	"8":  "Delete Org",
	"9":  "Create User",
	"10": "Update User",
	"11": "Lookup User",
	"12": "Delete User",
	"13": "Exit",
}

func main() {

	dbConn := db.Connect()
	defer dbConn.Close()

	userRepo := db.NewUserRepository(dbConn)
	userService := service.NewUserService(userRepo)

	orgRepo := db.NewOrganizationRepository(dbConn)
	orgService := service.NewOrganizationService(orgRepo)

	fmt.Println("Welcome to the Management System")

	for {
		choice := menu.DisplayMenuOptions([]string{"Organizations", "Jobs", "Users", "Exit"})

		switch choice {
		case "1":
			menu.OrganizationMenu(orgService)
		case "2":
			menu.JobMenu(dbConn)
		case "3":
			menu.UserMenu(userService)
		case "4":
			fmt.Println("Goodbye!")
			return
		}
	}
}
