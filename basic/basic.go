package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {

	testRange2()
}

func testRange() {
	arr := [2]int{1, 2}
	res := []*int{}
	res2 := []*int{}
	res3 := []*int{}

	for _, v := range arr {
		res = append(res, &v)
	}
	//会发现数组里的2指针是一样的
	//原因是因为  for-range其实是语法糖，内部调用还是for循环，初始化会拷贝带遍历的列表（如array，slice，map），
	//然后每次遍历的v都是对同一个元素的遍历赋值。 也就是说如果直接对v取地址，最终只会拿到一个地址，而对应的值就是最后遍历的那个元素所附给v的值。
	log.Println(res)

	//改进方法一般2种
	//局部拷贝后再存指针
	for _, v := range arr {
		v := v
		res2 = append(res2, &v)
	}

	//直接索引获取原本的值
	for k := range arr {
		res3 = append(res3, &arr[k])
	}
	log.Println(res2)
	log.Println(res3)
}

func testRange2() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	log.Println(v)

	var m = map[int]int{1: 1, 2: 2, 3: 3}
	//only del key once, and not del the current iteration key
	var o sync.Once
	for i := range m {
		o.Do(func() {
			for _, key := range []int{1, 2, 3} {
				log.Println(i, "-", key)
				if key != i {
					log.Println(i, "----", key)
					fmt.Printf("when iteration key %d, del key %d\n", i, key)
					delete(m, key)
					break
				}
			}
		})
		fmt.Printf("%d%d ", i, m[i])
	}
	log.Println(m)
}
