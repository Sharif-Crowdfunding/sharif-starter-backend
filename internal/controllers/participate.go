package controllers

import (
	"encoding/json"
	"net/http"
	"sharif-starter-backend/internal/models"
)

func BuyToken(w http.ResponseWriter, r *http.Request) {

}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	var projects []models.Project
	db.Find(&projects)
	for _, project := range projects {
		project.Token = &models.ProjectToken{}
		db.Where("project_id = ?", project.ID).First(project.Token)
	}
	json.NewEncoder(w).Encode(projects)
}
func GetSaleProjectTokenInfo(w http.ResponseWriter, r *http.Request) {
	projectToken := &models.ProjectToken{}
	json.NewDecoder(r.Body).Decode(projectToken)
	db.Where("project_id = ?", projectToken.ProjectId).First(projectToken)
	json.NewEncoder(w).Encode(projectToken)
}
