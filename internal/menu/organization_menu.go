package menu

import (
	"fmt"
	"go-db-demo/internal/domain"
	"go-db-demo/internal/service"

	"github.com/jmoiron/sqlx"
)

func OrganizationMenu(dbConn *sqlx.DB) {
	for {
		choice := DisplayMenuOptions([]string{
			"Create Org",
			"Update Org",
			"Lookup Org",
			"Delete Org",
			"List Orgs",
			"Back",
		})

		switch choice {
		case "1":
			createOrganizationCommand(dbConn)
			fmt.Println("Success!")
		case "2":
			fmt.Println("Updating Org...")
		case "3":
			fmt.Println("Looking up Org...")
		case "4":
			fmt.Println("Deleting Org...")
		case "5":
			listOrganizationsCommand(dbConn)
		case "6":
			return
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

func listOrganizationsCommand(dbConn *sqlx.DB) {
	organizations, err := service.GetAllOrganizations(dbConn)
	if err != nil {
		fmt.Println(err)
	}
	for _, o := range organizations {
		fmt.Println(o.ID, o.Name)
	}
}
