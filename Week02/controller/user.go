package controller

import (
	"github.com/advancego/week02/model"
	"github.com/advancego/week02/service"
	"github.com/pkg/errors"
)


func QueryUserById(id int) (*model.User, error ){
	s:=service.UserService{}
	u,err:=s.QueryUserById(id)
	if err!=nil{
		return &model.User{}, errors.Wrap(err,"controller wrong")
	}
	return u,err
}


