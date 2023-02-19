package proxy

import (
	"log"
	"time"
)

type IUser interface {
	Login(username, password string) error
}

// User 用户
type User struct {
}

// Login 用户登录
func (u *User) Login(username, password string) error {
	// TODO: 不实现细节
	return nil
}

// UserProxy 代理类
type UserProxy struct {
	user *User
}

func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{
		user: user,
	}
}

func (p *UserProxy) Login(username, password string) error {
	// TODO: before 这里可能会有一些统计的逻辑
	start := time.Now()

	//这里是原有的业务逻辑
	if err := p.user.Login(username, password); err != nil {
		return err
	}
	log.Printf("user login cost time: %s", time.Now().Sub(start))
	return nil
}
