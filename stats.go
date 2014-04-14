package stats

type Stats struct {
	users    []User
	channels []Channel
}

func (s *Stats) NewStats() *Stats {
	return &Stats{
	// users: make([]User),
	// channels: make([]Channels),
	}
}

func (s *Stats) AddChannel(channel Channel) {
	append(s.channels, channel)
}

func (s *Stats) FindChannel(name []byte) Channel {
	for _, channel := range s.channels {
		fmt.Printf("channel: %s\n", channel.topic)

		if channel.name == name {
			return channel
		}
	}

	return nil
}

func (s *Stats) HasChannel(channel []byte) bool {
	channel := s.FindChannel(channel)

	return channel != nil
}
