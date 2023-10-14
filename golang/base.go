package golang

type User struct{}

func (u User) toString() {
	println("user")
}

func (u *User) toString2() {
	println("user2")
}

type Manager struct {
	User
}

func (m Manager) toString() {
	println("manager")
}

func (m *Manager) toString2() {
	println("manager2")
}

type StringType1 interface {
	toString()
	toString2()
}

type StringType2 interface {
	StringType1
	toString3()
}
