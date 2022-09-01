package models

type ProjectReport struct {
	ProjectToken ProjectToken
	Participants []Participant
	TotalFunded  int64
}
