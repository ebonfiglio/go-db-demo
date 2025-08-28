package menu

import (
	"fmt"
	"go-db-demo/internal/domain"
)

func OrganizationMenu(organizationService domain.OrganizationService) {
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
			createOrganizationCommand(organizationService)
			fmt.Println("Success!")
		case "2":
			updateOrganizationCommand(organizationService)
			fmt.Println("Success!")
		case "3":
			getOrganizationCommand(organizationService)
		case "4":
			fmt.Println("Deleting Org...")
			deleteOrganizationCommand(organizationService)
		case "5":
			listOrganizationsCommand(organizationService)
		case "6":
			return
		}
	}
}

func createOrganizationCommand(organizationService domain.OrganizationService) {
	newOrgValues := getEntityInput()
	org, err := domain.JsonToOrganization(newOrgValues)
	if err != nil {
		fmt.Println(err)
	}
	org, err = organizationService.CreateOrganization(org)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Organziation ID: ", org.ID)
}

func listOrganizationsCommand(organizationService domain.OrganizationService) {
	organizations, err := organizationService.GetAllOrganizations()
	if err != nil {
		fmt.Println(err)
	}
	for _, o := range organizations {
		fmt.Println(o.ID, o.Name)
	}
}

func getOrganizationCommand(organizationService domain.OrganizationService) {
	id := getId()
	if id == 0 {
		return
	}
	organization, err := organizationService.GetOrganization(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(organization.ID, organization.Name)
}

func updateOrganizationCommand(organizationService domain.OrganizationService) {
	newOrgValues := getEntityInput()
	org, err := domain.JsonToOrganization(newOrgValues)
	if err != nil {
		fmt.Println(err)
	}
	org, err = organizationService.UpdateOrganization(org)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Organziation ID: ", org.ID)
	fmt.Println("Organziation Name: ", org.Name)
}

func deleteOrganizationCommand(organizationService domain.OrganizationService) {
	id := getId()
	if id == 0 {
		return
	}

	_, err := organizationService.DeleteOrganization(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Organization deleted!")
}
