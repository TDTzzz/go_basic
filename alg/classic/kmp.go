package main

import "log"

func main() {
	s := "bbc abcdab abcdabcdabde"
	p := "abcdabd"
	res := strStr(s, p)
	log.Println(res)
}

//获得next数组
//即：前缀后缀的重合数量
func getNext(p string) []int {
	pLen := len(p)
	next := make([]int, pLen, pLen)
	next[0] = -1
	next[1] = 0
	i := 0
	j := 1
	for j < pLen-1 { //因为next[pLen-1]由s[i] == s[pLen-2]算出
		if i == -1 || p[i] == p[j] { //-1代表了起始位不匹配，i=0,s[0]!=s[j]=>i=next[0]=-1
			i++
			j++
			next[j] = i
		} else {
			i = next[i]
		}
		log.Println(i, j, next)

	}
	log.Println(next)
	return next
}

//优化next数组
func getNextOptimize(p string) []int {
	pLen := len(p)
	next := make([]int, pLen, pLen)
	next[0] = -1
	next[1] = 0
	i := 0
	j := 1
	for j < pLen-1 { //因为next[pLen-1]由s[i] == s[pLen-2]算出
		if i == -1 || p[i] == p[j] { //-1代表了起始位不匹配，i=0,s[0]!=s[j]=>i=next[0]=-1
			i++
			j++
			if p[i] != p[j] { //因为出现在j位置不匹配的话会跳到next[j]=i位置去匹配,p[i] == p[j]肯定又是不匹配（优化核心点）
				next[j] = i
			} else {
				next[j] = next[i]
			}

			log.Println(i, j, next)
		} else {
			i = next[i]
		}
	}
	log.Println(next)
	return next
}

func KmpSearch(s, p string) int {
	i, j := 0, 0
	pLen := len(p)
	sLen := len(s)
	//next := getNext(p)
	next := getNextOptimize(p)
	for i < sLen && j < pLen {
		if j == -1 || s[i] == p[j] { //s[i]!=s[0]=>j=next[0]=-1,第0位不匹配所以i++，j++;j=0
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == pLen {
		return i - j
	} else {
		return -1
	}
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 || len(haystack) < len(needle) {
		return -1
	}
	if len(needle) == 1 { // 子串长度=0 时单独判断
		i := 0
		for ; i < len(haystack); i++ {
			if haystack[i] == needle[0] {
				break
			}
		}
		if i < len(haystack) {
			return i
		} else {
			return -1
		}
	}

	return KmpSearch(haystack, needle)
}
