package test

import (
	"go-app/library/boot"
	"go-app/library/db"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func initDB() {
	boot.Register(&db.DBStarter{})
	testBootRun()
}

func TestDB(t *testing.T) {
	initDB()

	err := db.GetDB().Ping()
	assert.Nil(t, err)
}
