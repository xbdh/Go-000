package main

import (
	"database/sql"
	"fmt"
	"github.com/advancego/week02/controller"
	"github.com/advancego/week02/model"
	"github.com/pkg/errors"
)

func main()  {
	_,err:=controller.QueryUserById(1)
	if errors.Is(err,sql.ErrNoRows){
		fmt.Println(model.User{Name: "mock"})
		fmt.Printf("%T \n %v\n",errors.Cause(err),errors.Cause(err))
		fmt.Printf("%+v",err)
	}
}