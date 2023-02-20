package app

import (
	"Epam_final/pkg/db/models"
	"Epam_final/pkg/front"
)

func makeUserRequestRes(row *models.UserRequest) *front.UserRequestRes {
	return &front.UserRequestRes{
		ID:     row.ID,
		AccID:  row.AccID,
		Status: row.Status,
	}
}
