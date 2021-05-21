package service

import (
	"context"
	"errors"
	"github.com/ssuxue/mellow-server/common/util"
	"github.com/ssuxue/mellow-server/domain/model"
	"github.com/ssuxue/mellow-server/domain/repository"
	"github.com/ssuxue/mellow-server/proto"
	"log"
	"time"
)

type UserServer struct {
	Service repository.IUserRepository
}

func (u *UserServer) Login(_ context.Context, req *proto.LoginRequest) (*proto.UserResponse, error) {
	user := &model.User{}
	user = u.Service.FindByUsername(req.Email)

	if user.IsEmpty() {
		return nil, errors.New("user doesn't exist")
	}

	isOk, _ := util.ValidatePassword(req.Password, user.Password)
	if !isOk {
		return nil, errors.New("username or password is wrong")
	}

	res := &proto.UserResponse{}

	res.Email = user.Username
	res.Id = int64(user.ID)
	res.AccessToken = "" // TODO
	res.RefreshAfter = 10
	res.AccessExpire = 60
	return res, nil
}

func (u *UserServer) Register(_ context.Context, req *proto.RegisterRequest) (*proto.UserResponse, error) {
	user := &model.User{}
	user = u.Service.FindByUsername(req.Email)
	if !user.IsEmpty() {
		return nil, errors.New("user is already exist")
	}

	user.Username = req.Email
	password, err := util.GeneratePassword(req.Password)
	if err != nil {
		log.Println(err)
	}
	user.Password = string(password)
	user.Phone = req.Username
	user.Birthday = time.Now()

	if res := u.Service.Insert(user); res == 0 {
		return nil, errors.New("failed to add user")
	}

	res := &proto.UserResponse{}
	res.Email = req.Email
	res.AccessToken = ""

	return res, nil
}

func (u *UserServer) Userinfo(_ context.Context, req *proto.UserinfoRequest) (*proto.UserResponse, error) {
	user := u.Service.FindById(int(req.Uid))

	if user.IsEmpty() {
		return nil, errors.New("failed to get user information")
	}

	res := &proto.UserResponse{}
	res.Email = user.Username
	res.Id = req.Uid
	return res, nil
}
