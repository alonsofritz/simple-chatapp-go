package user

import (
	"context"
	"strconv"
	"time"

	"github.com/alonsofritz/simple-chatapp-go/server/util"
)

type userService struct {
	UserRepository
	timeout time.Duration
}

func NewUserService(userRepository UserRepository) UserService {
	return &userService{
		userRepository,
		time.Duration(2) * time.Second,
	}
}

func (s *userService) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	r, err := s.UserRepository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}

	return res, nil
}
