package qs

func QuickSort2(obj []int, left, right int) int {

pivot := obj[right]
i := left - 1

for j := left; j <= right - 1; j++ {
	if obj[j] <= pivot {
		i++
		swap2(obj, i, j)
	}

}

swap2(obj, i +1, right)


return i + 1

}

func swap2(obj []int, i, j int) {
	obj[j], obj[i] = obj[i], obj[j]
}
