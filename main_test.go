package go_orm

import (
	"github.com/stretchr/testify/suite"
	"github.com/jinzhu/gorm"
	"go_orm/model"
	"database/sql"
	"github.com/stretchr/testify/require"
	"regexp"
	"fmt"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository Repository
	person     *model.Person
}




func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.repository = CreateRepository(s.DB)
}

func (s *Suite) Test_repository_Get() {
	var (
		id   = "test_id"
		name = "test-name"
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "person" WHERE (id = $1)`)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
		AddRow(id, name))

	res, err := s.repository.Get(id)

	require.NoError(s.T(), err)
	fmt.Println(res)
	// require.Nil(s.T(), deep.Equal(&model.Person{ID: id, Name: name}, res))
}

func (s *Suite) Test_repository_Create() {
	var (
		id   = "test_id"
		name = "test-name"
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "person" ("id","name") 
       VALUES ($1,$2) RETURNING "person"."id"`)).
		WithArgs(id, name).
		WillReturnRows(
		sqlmock.NewRows([]string{"id"}).AddRow(id))

	err := s.repository.Create(id, name)

	require.NoError(s.T(), err)
}

func TestRepo_Get(t *testing.T) {
	var (
		id   = "test_id"
		name = "test-name"
	)

	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	gorm.Open("mysql", db)


	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `people` WHERE (id = ?)")).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
		AddRow(id, name))

	DB, _ := gorm.Open("mysql", db)

	repo := CreateRepository(DB)
	ret, err := repo.Get("test_id")
	// ret, err := db.Exec("SELECT * FROM `person` *")
	if err != nil{
		t.Error(err)
	}
	fmt.Println(ret)
}

func TestRepo_Get1(t *testing.T) {
	var (
		id   = "test_id"
		name = "test-name"
	)

	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	gorm.Open("mysql", db)


	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `people` WHERE (id = ?)")).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
		AddRow(id, name))

	DB, _ := gorm.Open("mysql", db)

	repo := CreateRepository(DB)
	ret, err := repo.Get("test_id")
	// ret, err := db.Exec("SELECT * FROM `person` *")
	if err != nil{
		t.Error(err)
	}
	fmt.Println(ret)
}

func TestRepo_Create(t *testing.T) {
	var (
		id   = "test_id"
		name = "test-name"
	)

	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	gorm.Open("mysql", db)


	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `people` \\(`id`,`name`\\) VALUES \\(\\?\\,\\?\\)").
		WithArgs(id, name).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	// mock.ExpectQuery(regexp.QuoteMeta(
	// 	`INSERT INTO "person" ("id","name")
     //   VALUES ($1,$2) RETURNING "person"."id"`)).
	// 	WithArgs(id, name).
	// 	WillReturnRows(
	// 	sqlmock.NewRows([]string{"id"}).AddRow(id))

	DB, _ := gorm.Open("mysql", db)

	repo := CreateRepository(DB)
	ret := repo.Create("test_id", "test-name")
	// ret, err := db.Exec("SELECT * FROM `person` *")
	if ret != nil{
		t.Error(ret)
	}
	fmt.Println(ret)
}