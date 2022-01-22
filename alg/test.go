package main

//练手

func main() {

}

func quickSort(data []int, left int, right int) []int {
	//1.找基准数

	if left < right {

		pivot := (left + right) / 2
		value := data[pivot]
		i, j := left, right

		for {
			for value < data[i] {
				i++
			}
			for value > data[j] {
				j--
			}

			if i >= j {
				break
			}
			data[i], data[j] = data[j], data[i]
		}

		quickSort(data, left, i-1)
		quickSort(data, j+1, right)
	}
}
