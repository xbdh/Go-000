package service

import (
	"github.com/advancego/week02/dao"
	"github.com/advancego/week02/model"
	"github.com/pkg/errors"
)

type UserService struct {

}

func(s *UserService) QueryUserById(id int) (*model.User, error ){
		db:=dao.Database{}
		u,err:=db.QueryUser(id)
		if err!=nil{
			return &model.User{}, errors.Wrap(err,"service wrong")
		}
		return u,err
}