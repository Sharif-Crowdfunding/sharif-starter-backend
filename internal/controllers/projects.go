package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sharif-starter-backend/internal/models"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	project := &models.Project{}
	json.NewDecoder(r.Body).Decode(project)
	fmt.Println(project)
	createdProject := db.Create(project)
	var errMessage = createdProject.Error

	if errMessage != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdProject)
}
func GetUserProjects(w http.ResponseWriter, r *http.Request) {
	project := &models.Project{}
	json.NewDecoder(r.Body).Decode(project)
	fmt.Println(project)
	createdProject := db.Create(project)
	var errMessage = createdProject.Error

	if errMessage != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdProject)
}
