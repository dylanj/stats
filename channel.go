package stats

type Channel struct {
	name     []byte
	topic    []byte
	joins    int
	parts    int
	users    []User
	messages []Message
}
