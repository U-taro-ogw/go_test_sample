package handlers

import (
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestSignup(t *testing.T) {

	db, mock, err := createDbMock()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)
	//redisCon := authenticationDb.RedisConnect()

	UserHandler{Db: db}.Signup(c)

	T.Run("test", func(t *testing.T) {


	})
}

func createDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb, err := gorm.Open("mysql", db)
	if err != nil {
		return nil, nil, err
	}
	return gdb, mock, nil
}