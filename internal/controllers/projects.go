package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sharif-starter-backend/internal/models"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("user").(*models.Token)
	project := &models.Project{}
	json.NewDecoder(r.Body).Decode(project)
	project.UserEmail = token.Email
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
	db.Where("user_email = ?", token.Email).Find(&projects)
	json.NewEncoder(w).Encode(projects)
}

func GetProjectTokenInfo(w http.ResponseWriter, r *http.Request) {
	projectToken := &models.ProjectToken{}
	json.NewDecoder(r.Body).Decode(projectToken)
	db.Where("project_id = ?", projectToken.ProjectId).First(projectToken)
	json.NewEncoder(w).Encode(projectToken)
}

func AddProjectTokenDistribution(w http.ResponseWriter, r *http.Request) {
	projectToken := &models.ProjectToken{}
	json.NewDecoder(r.Body).Decode(projectToken)

	var project models.Project
	db.Where("id = ?", projectToken.ProjectId).First(&project)
	project.Status = models.ReadyForSale
	db.Save(&project)

	createdProjectToken := db.Create(&projectToken)
	var errMessage = createdProjectToken.Error

	if errMessage != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdProjectToken)
}

func CancelProject(w http.ResponseWriter, r *http.Request) {
	project := &models.Project{}
	json.NewDecoder(r.Body).Decode(project)
	db.Where("id = ?", project.ID).First(project)

	project.Status = models.Canceled
	db.Save(&project)

	json.NewEncoder(w).Encode(project)
}

func ReleaseProjectToPublic(w http.ResponseWriter, r *http.Request) {
	project := &models.Project{}
	json.NewDecoder(r.Body).Decode(project)
	db.Where("id = ?", project.ID).First(project)

	project.Status = models.InProgress
	db.Save(&project)

	json.NewEncoder(w).Encode(project)
}

func FinishProjectSale(w http.ResponseWriter, r *http.Request) {
	project := &models.Project{}
	json.NewDecoder(r.Body).Decode(project)
	db.Where("id = ?", project.ID).First(project)

	project.Status = models.Finished
	db.Save(&project)

	json.NewEncoder(w).Encode(project)
}
