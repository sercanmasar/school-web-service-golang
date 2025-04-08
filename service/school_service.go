package service

import (
	"errors"
	"fmt"
	"school-service/common/dto"
	"school-service/domain/models"
	"school-service/persistence"
)

type ISchoolService interface {
	CreateSchool(schoolRequest dto.SchoolRequest) (bool, error)
	GetById(schoolId int64) (dto.SchoolResponse, error)
}

type SchoolService struct {
	schoolRepository persistence.ISchoolRepository
}

func NewSchoolService(schoolRepository persistence.ISchoolRepository) ISchoolService {
	return &SchoolService{schoolRepository: schoolRepository}
}

func (schoolService *SchoolService) GetById(schoolId int64) (dto.SchoolResponse, error) {
	school, err := schoolService.schoolRepository.GetSchoolById(schoolId)
	if err != nil {
		return dto.SchoolResponse{}, err
	}
	return dto.SchoolResponse{
		Id:          school.Id,
		Name:        school.Name,
		Title:       school.Title,
		SchoolLevel: school.SchoolLevel,
		SchoolType:  school.SchoolType,
	}, errors.New(fmt.Sprintf("School not found by id: %s", schoolId))
}

func (schoolService *SchoolService) CreateSchool(schoolRequest dto.SchoolRequest) (bool, error) {
	school := models.School{Name: schoolRequest.Name, Title: schoolRequest.Title, SchoolLevel: schoolRequest.SchoolLevel, SchoolType: schoolRequest.SchoolType}
	_, err := schoolService.schoolRepository.CreateSchool(school)
	if err != nil {
		return false, err
	}
	return true, nil
}
