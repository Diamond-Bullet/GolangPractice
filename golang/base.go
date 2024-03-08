package golang

import (
	"github.com/sirupsen/logrus"
)

type User struct{}

func (u User) ToString() {
	println("user")
}

func (u *User) ToStringPtr() {
	println("user2")
}

type Manager struct {
	User
}

func (m Manager) ToString() {
	println("manager")
}

func (m *Manager) ToStringPtr() {
	println("manager ptr")
}

type StringType1 interface {
	ToString()
	ToStringPtr()
}

type StringType2 interface {
	StringType1
	ToString2()
}

var logger = logrus.New()

func init() {
	logger.SetNoLock()
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{})
}