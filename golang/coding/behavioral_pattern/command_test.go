package behavioral_pattern

import (
	"github.com/gookit/color"
	"testing"
)

// Command turns a request into a stand-alone object.
//
// Benefits:
// 1. Decouples the callee from the invoker.
// 2. provide support for undoable operations.

func TestCommand(t *testing.T) {
	userInfoClient := &UserInfoClient{}

	getUserInfoCommand := &GetUserInfoCommand{UserInfoClient: userInfoClient}
	updateUserInfoCommand := &UpdateUserInfoCommand{UserInfoClient: userInfoClient}

	stub := &Stub{}

	stub.SetCommand(getUserInfoCommand)
	stub.ExecuteCommand()

	stub.SetCommand(updateUserInfoCommand)
	stub.ExecuteCommand()
}

type Command interface {
	Execute()
}

type UserInfoClient struct {
}

func (u *UserInfoClient) GetUserInfo() {
	color.Blueln("Get user info")
}

func (u *UserInfoClient) UpdateUserInfo() {
	color.Blueln("Update user info")
}

type GetUserInfoCommand struct {
	UserInfoClient *UserInfoClient
}

func (g *GetUserInfoCommand) Execute() {
	g.UserInfoClient.GetUserInfo()
}

type UpdateUserInfoCommand struct {
	UserInfoClient *UserInfoClient
}

func (u *UpdateUserInfoCommand) Execute() {
	u.UserInfoClient.UpdateUserInfo()
}

type Stub struct {
	command Command
}

func (s *Stub) SetCommand(command Command) {
	s.command = command
}

func (s *Stub) ExecuteCommand() {
	s.command.Execute()
}
