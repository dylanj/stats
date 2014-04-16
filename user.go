package stats

const NameLength = 32
const HostmaskLength = 128

type User struct {
	Name     []byte
	Hostmask []byte
	Messages []*Message
}

func NewUser(name []byte, hostmask []byte) *User {
	user := User{
		Name:     make([]byte, NameLength),
		Hostmask: make([]byte, HostmaskLength),
		Messages: make([]*Message, 50),
	}

	copy(user.Name, name)
	copy(user.Hostmask, hostmask)

	return &user
}
