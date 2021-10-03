package test

import (
	"fmt"
	"go-app/library/boot"
	"go-app/library/db"
	"go-app/library/util/json"
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

type User struct {
	Id       string `json:"id" xorm:"id"`
	Username string `json:"username" xorm:"username"`
	Password string `json:"password" xorm:"password"`
	Nickname string `json:"nickname" xorm:"nickname"`
	CreateAt int64  `json:"create_at" xorm:"create_at"`
	UpdateAt int64  `json:"update_at" xorm:"update_at"`
}

func TestSelect(t *testing.T) {
	initDB()

	e := db.GetDB()
	var users []*User
	err := e.Table("user").Find(&users)
	assert.Nil(t, err)
	fmt.Println("size:", len(users))
	json.PrintJsonPretty(users)
}
