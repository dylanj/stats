package stats

type TokenCounter struct {
	All   map[string]uint
	Top   TopTokenArray
	Count uint
}

// NewTokens initializes the Tokens map.
func NewTokenCounter() TokenCounter {
	return TokenCounter{
		All: make(map[string]uint),
		Top: make([]TopToken, 0, topTokenMaxSize),
	}
}

func (tc *TokenCounter) addToken(token string) {
	tc.All[token]++
	count := tc.All[token]

	tc.Top.insert(token, count)
	tc.Count++
}
