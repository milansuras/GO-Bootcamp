package main

import "fmt"

func main() {
	l := list{"A", "B", "C"}
	input := list{"D", "E", "F"}
	interresult := l.intersect(input)
	fmt.Println("Intersection:=", interresult)

	l.remove("B")
	fmt.Println("After removal", l)
}

type list []string

func (l list) intersect(input list) list {
	mp := make(map[string]bool)
	for _, ele := range l {
		mp[ele] = true
	}

	var result list
	for _, ele := range input {
		if mp[ele] {
			result = append(result, ele)
		}
	}
	return result
}

func (l *list) remove(in string) {
	index := l.indexof(in)
	if index != -1 {
		*l = append((*l)[:index], (*l)[index+1:]...)
	}
}

func (l list) indexof(in string) int {
	for index, ele := range l {
		if ele == in {
			return index
		}
	}
	return -1
}
