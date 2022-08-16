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
	token := r.Context().Value("user").(*models.Token)
	var projects []models.Project
	db.Find(&projects).Where("email = ?", token.Email)
	json.NewEncoder(w).Encode(projects)
}

func AddProjectTokenDistribution(w http.ResponseWriter, r *http.Request) {
	projectToken := &models.ProjectToken{}
	json.NewDecoder(r.Body).Decode(projectToken)

	var project models.Project
	db.First(&project).Where("id = ?", projectToken.ProjectId)
	project.Status = models.ReadyForSale
	db.Save(&project)

	createdProjectToken := db.Create(&projectToken)
	var errMessage = createdProjectToken.Error

	if errMessage != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdProjectToken)
}
