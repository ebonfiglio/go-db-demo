package main

import (
	"bufio"
	"fmt"
	"go-db-demo/internal/db"
	"go-db-demo/internal/domain"
	"go-db-demo/internal/service"
	"os"
	"sort"
	"strings"

	"github.com/jmoiron/sqlx"
)

var commands = map[string]string{
	"1": "Create Org",
	"2": "Update Org",
	"3": "Lookup Org",
	"4": "Delete Org",
	"5": "Create User",
	"6": "Update User",
	"7": "Lookup User",
	"8": "Delete User",
	"9": "Exit",
}

func main() {

	dbConn := db.Connect()
	defer dbConn.Close()

	var input string

	fmt.Println("Welcome to the User Management System")
	for {
		listCommands(&input)

		switch input {
		case "1":
			createOrganizationCommand(dbConn)
		case "5":
			createUserCommand(dbConn)
		case "6":
			fmt.Println("Enter the users id:")
		case "7":
			fmt.Println("Enter the users id:")
		case "8":
			fmt.Println("Enter the users id:")
		case "9":
			fmt.Println("Goodbye")
			return
		default:
			listCommands(&input)
		}
	}
}

func createOrganizationCommand(dbConn *sqlx.DB) {
	newOrgValues := getNewEntityInput()
	org, err := domain.JsonToOrganization(newOrgValues)
	if err != nil {
		fmt.Println(err)
	}
	org, err = service.CreateOrganization(org, dbConn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Organziation ID: ", org.ID)
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

func getNewEntityInput() string {
	fmt.Println("Enter new entity as JSON:")
	reader := bufio.NewReader(os.Stdin)
	json, _ := reader.ReadString('\n')

	json = strings.TrimSpace(json)
	return json
}

func listCommands(input *string) {
	fmt.Println("Select a command")

	keys := make([]string, 0, len(commands))
	for k := range commands {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, ":", commands[k])
	}
	fmt.Scanln(input)
}
