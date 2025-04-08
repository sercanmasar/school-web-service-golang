package dto

import "school-service/domain/enums"

type SchoolRequest struct {
	Name        string
	Title       string
	SchoolLevel enums.SchoolLevel
	SchoolType  enums.SchoolType
}
