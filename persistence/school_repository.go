package persistence

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"school-service/domain/enums"
	"school-service/domain/models"
)

type ISchoolRepository interface {
	CreateSchool(school models.School) (bool, error)
	GetSchoolById(id int64) (models.School, error)
}

type SchoolRepository struct {
	connectionPool *pgxpool.Pool
}

func NewSchoolRepository(pool *pgxpool.Pool) ISchoolRepository {
	return &SchoolRepository{pool}
}

func (schoolRepository *SchoolRepository) CreateSchool(school models.School) (bool, error) {
	ctx := context.Background()
	createSchoolSql := `insert into school(name,title,schoolLevel,schoolType) values($1,$2,$3,$4)`
	_, err := schoolRepository.connectionPool.Exec(ctx, createSchoolSql, school.Name, school.Title, school.SchoolLevel, school.SchoolType)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (schoolRepository *SchoolRepository) GetSchoolById(schoolId int64) (models.School, error) {
	ctx := context.Background()
	getSchoolByIdSql := `select * from school where id = $1`
	queryResult, err := schoolRepository.connectionPool.Query(ctx, getSchoolByIdSql, schoolId)

	if err != nil {
		return models.School{}, errors.New(fmt.Sprintf("School not found by id %v", schoolId))
	}

	var id int64
	var name string
	var title string
	var schoolLevel int32
	var schoolType int32

	err = queryResult.Scan(&id, &name, &title, &schoolLevel, &schoolType)
	if err != nil {
		return models.School{}, err
	}

	return models.School{
		Id:          id,
		Name:        name,
		Title:       title,
		SchoolLevel: enums.SchoolLevel(schoolLevel),
		SchoolType:  enums.SchoolType(schoolType),
	}, nil
}
