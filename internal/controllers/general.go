package controllers

import (
	"encoding/json"
	"net/http"
	"sharif-starter-backend/internal/models"
	"sharif-starter-backend/pkg/mail"
)

func JoinWaitingList(w http.ResponseWriter, r *http.Request) {
	member := &models.NewsLetterMember{}
	json.NewDecoder(r.Body).Decode(member)

	finded := &models.NewsLetterMember{}
	db.Where(member).First(finded)
	if finded.Email != "" {
		err := ErrorResponse{
			Err: "ایمیل قبلا ثبت شده است.",
		}
		json.NewEncoder(w).Encode(err)
	}

	created := db.Create(member)
	go mail.SendEmail(member.Email, "JOIN")
	json.NewEncoder(w).Encode(created)
}
