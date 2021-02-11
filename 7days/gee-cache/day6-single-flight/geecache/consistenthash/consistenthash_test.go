package consistenthash

import (
	"strconv"
	"testing"
)

func TestHashing(t *testing.T) {
	hash := New(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	//对应的虚拟节点应该为
	//2 4 6 12 14 16 22 24 26
	hash.Add("6", "4", "2")

	//把上面的虚拟节点想象成环，可以推测数据为如下map
	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}

	//添加一个节点，也就是 8 18 28
	hash.Add("8")

	//由于虚拟节点里有28 离27更近，所以映射到了8上
	testCases["27"] = "8"

	for k, v := range testCases {
		if hash.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s", k, v)
		}
	}
}

