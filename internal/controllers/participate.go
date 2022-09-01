package controllers

import (
	"encoding/json"
	"net/http"
	"sharif-starter-backend/internal/models"
)

func BuyToken(w http.ResponseWriter, r *http.Request) {
	participant := &models.Participant{}
	json.NewDecoder(r.Body).Decode(participant)

	projectToken := &models.ProjectToken{}
	db.Where("id = ?", participant.ProjectTokenId).First(projectToken)
	participant.PurchasedTokens = int64(participant.PurchasedTokens / projectToken.PricePerTokenByGwei)
	savedParticipant := db.Save(participant)
	json.NewEncoder(w).Encode(savedParticipant)
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

func GetProjectById(w http.ResponseWriter, r *http.Request) {
	project := &models.Project{}
	json.NewDecoder(r.Body).Decode(project)
	db.Where("id = ?", project.ID).First(project)
	json.NewEncoder(w).Encode(project)
}
func GetSaleProjectTokenInfo(w http.ResponseWriter, r *http.Request) {
	projectToken := &models.ProjectToken{}
	json.NewDecoder(r.Body).Decode(projectToken)
	db.Where("project_id = ?", projectToken.ProjectId).First(projectToken)
	json.NewEncoder(w).Encode(projectToken)
}
