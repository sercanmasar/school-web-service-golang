package dto

import "school-service/domain/enums"

type SchoolResponse struct {
	Id          int64
	Name        string
	Title       string
	SchoolLevel enums.SchoolLevel
	SchoolType  enums.SchoolType
}
