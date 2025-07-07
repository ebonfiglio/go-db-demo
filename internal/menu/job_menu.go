package menu

import (
	"fmt"
	"go-db-demo/internal/domain"
	"go-db-demo/internal/service"

	"github.com/jmoiron/sqlx"
)

func JobMenu(dbConn *sqlx.DB) {
	for {
		choice := DisplayMenuOptions([]string{
			"Create Job",
			"Update Job",
			"Lookup Job",
			"Delete Job",
			"Back",
		})

		switch choice {
		case "1":
			createJobCommand(dbConn)
			fmt.Println("Success!")
		case "2":
			fmt.Println("Updating Job...")
		case "3":
			fmt.Println("Looking up Job...")
		case "4":
			fmt.Println("Deleting Job...")
		case "5":
			return
		}
	}
}

func createJobCommand(dbConn *sqlx.DB) {
	newJobValues := getNewEntityInput()
	job, err := domain.JsonToJob(newJobValues)
	if err != nil {
		fmt.Println(err)
	}
	job, err = service.CreateJob(job, dbConn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Job ID: ", job.ID)
}
