package dto

import "school-service/domain/enums"

type SchoolRequest struct {
	Name        string            `json:"name"`
	Title       string            `json:"title"`
	SchoolLevel enums.SchoolLevel `json:"schoolLevel"`
	SchoolType  enums.SchoolType  `json:"schoolType"`
}
