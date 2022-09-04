package lexer

// accept only digit "1"
// 2 states - init, end
func simpleState(input string) func() (string, bool) {
	state := "init"
	indx := 0

	return func() (string, bool) {
		switch state {
		case "init":
			char := input[indx]
			if char == '1' {
				state = "accept"
				indx++
				return state, true
			}
			// stuck, reject
			state = "reject"
			return state, false

		case "accept":
			if indx >= len(input) {
				return "accept", false
			}
		}
		return "reject", false
	}
}

// any number of 1s followed by a single 0
func anyNumberOfOnes(input string) func() (string, bool) {
	state := "A"
	idx := 0
	return func() (string, bool) {
		switch state {
		case "A":
			if idx >= len(input) {
				return "reject", false
			} else if input[idx] == '1' {
				state = "A"
			} else if input[idx] == '0' {
				state = "B"
			}
			idx++
			return state, true
		case "B":
			if idx >= len(input) {
				return "accept", false
			}
		}
		return "reject", false
	}
}
