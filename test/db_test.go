package test

import (
	"fmt"
	"go-app/library/boot"
	"go-app/library/db"
	"go-app/library/util/json"
	"testing"
	"time"

	"xorm.io/builder"

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
	Id       int    `json:"id" xorm:"id pk autoincr"`
	Username string `json:"username" xorm:"username "`
	Password string `json:"password" xorm:"password"`
	Nickname string `json:"nickname" xorm:"nickname"`
	CreateAt int64  `json:"create_at" xorm:"create_at"`
	UpdateAt int64  `json:"update_at" xorm:"update_at"`
}

func TestInsert(t *testing.T) {
	initDB()

	user := &User{
		Username: "tom",
		Password: "111111",
		Nickname: "tom",
		CreateAt: time.Now().UnixNano() / 1e6,
		UpdateAt: time.Now().UnixNano() / 1e6,
	}
	affected, err := db.GetDB().Table("user").Insert(user)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), affected)
	json.PrintJsonPretty(user)
}

func TestDelete(t *testing.T) {
	initDB()

	affected, err := db.GetDB().Table("user").Where(builder.Eq{"username": "tom"}).Delete(&User{})
	assert.Nil(t, err)
	fmt.Println("affected:", affected)
}

func TestUpdate(t *testing.T) {
	initDB()

	update := map[string]interface{}{"update_at": time.Now().UnixNano() / 1e6}
	affected, err := db.GetDB().Table("user").Where(builder.Eq{"username": "admin"}).Update(update)
	assert.Nil(t, err)
	fmt.Println("affected:", affected)
}

func TestSelect(t *testing.T) {
	initDB()

	var users []*User
	err := db.GetDB().Table("user").Find(&users)
	assert.Nil(t, err)
	fmt.Println("size:", len(users))
	json.PrintJsonPretty(users)
}
