package stats

const topTokenMaxSize = 50

type TopTokenArray []TopToken

type TopToken struct {
	Token string
	Count uint
}

func (a *TopTokenArray) insert(token string, count uint) {
	ta := *a // allow accessing token array without indirection everywhere
	insertAt := -1
	currentIndex := -1

	if len(ta) == 0 {
		ta = append(ta, TopToken{token, count})
		*a = ta
		return
	}

	for i, t := range ta {
		if insertAt == -1 && count > t.Count {
			insertAt = i
		}

		if currentIndex == -1 && token == t.Token {
			currentIndex = i
		}

		if currentIndex != -1 && insertAt != -1 {
			break
		}
	}

	if currentIndex >= 0 {
		if insertAt < 0 {
			return
		}
		if insertAt < currentIndex {
			ta[currentIndex].Token, ta[insertAt].Token =
				ta[insertAt].Token, ta[currentIndex].Token
			ta[insertAt].Count = count
		} else {
			ta[currentIndex].Count = count
		}
	} else if len(ta) < topTokenMaxSize {
		ta = append(ta, TopToken{token, count})
		*a = ta
	} else if insertAt >= 0 {
		ta[insertAt].Token = token
		ta[insertAt].Count = count
	}
}
