package main

import "sort"

//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

//方法1，排序
func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}

	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

//方法2，计数
func groupAnagrams2(strs []string)[][]string{
	  mp := map[[26]int][]string{}
    for _, str := range strs {
        cnt := [26]int{}
        for _, b := range str {
            cnt[b-'a']++
        }
        mp[cnt] = append(mp[cnt], str)
    }
    ans := make([][]string, 0, len(mp))
    for _, v := range mp {
        ans = append(ans, v)
    }
    return ans
}