package dao

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/advancego/week02/model"
)

type Database struct {

}

func (db *Database) QueryUser(id int) (user *model.User,err error){
		u:=&model.User{}
		return u ,errors.Wrap(sql.ErrNoRows,"user do not exists in db")
}

