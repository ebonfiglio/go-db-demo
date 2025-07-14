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
			"Get All Jobs",
			"Back",
		})

		switch choice {
		case "1":
			createJobCommand(dbConn)
			fmt.Println("Success!")
		case "2":
			fmt.Println("Updating Job...")
		case "3":
			getJobCommand(dbConn)
		case "4":
			fmt.Println("Deleting Job...")
		case "5":
			getAllJobsCommand(dbConn)
		case "6":
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

func getAllJobsCommand(dbConn *sqlx.DB) {
	jobs, err := service.GetAllJobs(dbConn)
	if err != nil {
		fmt.Println(err)
	}

	for _, j := range jobs {
		fmt.Println(j.ID, j.Name, j.OrganizationID)
	}
}

func getJobCommand(dbConn *sqlx.DB) {
	id := getId()
	if id == 0 {
		return
	}
	job, err := service.GetJob(id, *dbConn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(job.ID, job.Name, job.OrganizationID)

}
