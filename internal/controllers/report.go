package controllers

import (
	"encoding/json"
	"net/http"
	"sharif-starter-backend/internal/models"
)

func GetProjectReport(w http.ResponseWriter, r *http.Request) {
	projectToken := &models.ProjectToken{}
	json.NewDecoder(r.Body).Decode(projectToken)
	db.Where("project_id = ?", projectToken.ProjectId).First(projectToken)

	var participants []models.Participant
	db.Where("project_token_id = ?", projectToken.ProjectId).Find(&participants)

	var totalFunded int64
	for _, participant := range participants {
		totalFunded += participant.PurchasedTokens * projectToken.PricePerTokenByGwei
	}
	report := &models.ProjectReport{
		ProjectToken: *projectToken,
		Participants: participants,
		TotalFunded:  totalFunded,
	}
	json.NewEncoder(w).Encode(report)
}
