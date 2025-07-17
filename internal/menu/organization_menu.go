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
			updateOrganizationCommand(dbConn)
			fmt.Println("Success!")
		case "3":
			getOrganizationCommand(dbConn)
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
	newOrgValues := getEntityInput()
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

func getOrganizationCommand(dbConn *sqlx.DB) {
	id := getId()
	if id == 0 {
		return
	}
	organization, err := service.GetOrganization(id, dbConn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(organization.ID, organization.Name)
}

func updateOrganizationCommand(dbConn *sqlx.DB) {
	newOrgValues := getEntityInput()
	org, err := domain.JsonToOrganization(newOrgValues)
	if err != nil {
		fmt.Println(err)
	}
	org, err = service.UpdateOrganization(org, dbConn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Organziation ID: ", org.ID)
	fmt.Println("Organziation Name: ", org.Name)
}
