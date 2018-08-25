package handlers

import (
	"net/http"

	"encoding/json"
	"middware-go/models"
)

//make a slice to hold our jobs data

var jobs []models.Job

//A handler to fetch all the jobs
func GetJobs(w http.ResponseWriter, r *http.Request) {
	///add some job to the slice
	jobs = append(jobs, models.Job{ID: 1, Name: "Accounting"})
	jobs = append(jobs, models.Job{ID: 2, Name: "Programming"})
	json.NewEncoder(w).Encode(jobs)
}
