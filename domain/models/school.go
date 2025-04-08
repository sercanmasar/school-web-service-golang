package models

import "school-service/domain/enums"

type School struct {
	Id          int64
	Name        string
	Title       string
	SchoolLevel enums.SchoolLevel
	SchoolType  enums.SchoolType
}
