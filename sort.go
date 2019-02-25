package sort

import "fmt"

func main() {
	i := []int{10, 2, 7, 3, 0, 6, 4, 9}
	selectSort(i)
	fmt.Println(i)
}

func bubbleSort(s []int) {
	length := len(s)
	for i := 0; i < length-1; i++ {
		for j := 0; j < i; j++ {
			if s[j] > s[j+1] {
				s[j+1], s[j] = s[j], s[j+1]
			}
		}
	}
}

func selectSort(s []int) {
	length := len(s)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if s[j] < s[min] {
				min = j
			}
		}
		if min != i {
			s[i], s[min] = s[min], s[i]
		}
	}
}

func fastSort(s []int) []int {
	length := len(s)
	if length < 1 {
		return s
	}
	stand := s[0]
	max := []int{}
	min := []int{}
	for i, _ := range s {
		fmt.Println(s[i])
		if s[i] < stand {
			min = append(min, s[i])
		} else if s[i] > stand {
			max = append(max, s[i])
		}
	}
	array := []int{}
	array = append(array, fastSort(min)...)
	array = append(array, stand)
	array = append(array, fastSort(max)...)
	return array
}
