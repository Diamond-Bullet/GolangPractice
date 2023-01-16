package golang

type User struct{}

func (u User) toString() {
	println("user")
}

func (u *User) toString2() {
	println("user user")
}

type Manager struct {
	User
}

func (m Manager) toString() {
	println("manager")
}

func (m *Manager) toString2() {
	println("manager manager")
}

type MultiString interface {
	toString()
	toString2()
}

type MMMMMultiString interface {
	MultiString
	toString3()
}

func Add(x, y int) int {
	return x + y
}
