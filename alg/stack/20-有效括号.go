package main

import "log"

func main() {
	s := "()[]{}"
	res := isValid(s)
	log.Println(res)
}

func isValid(s string) bool {
	l := len(s)
	if l%2 != 0 {
		return false
	}

	pairs := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stk := []byte{}

	for i := 0; i < l; i++ {
		tmp := s[i]
		if _, ok := pairs[tmp]; ok {
			stk = append(stk, tmp)
		} else {
			if len(stk) == 0 || pairs[stk[len(stk)-1]] != tmp {
				return false
			}
			stk = stk[:len(stk)-1]
		}
	}
	return len(stk) == 0
}
