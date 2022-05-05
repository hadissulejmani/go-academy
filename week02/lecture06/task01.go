package main

import "fmt"

type Item struct {
	Value    int
	PrevItem *Item
}

type MagicList struct {
	LastItem *Item
}

func main() {
	l := &MagicList{}
	add(l, 3)
	add(l, 4)
	add(l, 41)
	add(l, 92)
	add(l, 33)
	add(l, 54)
	add(l, 41)
	add(l, 72)
	add(l, 13)
	add(l, 24)

	fmt.Println(toSlice(l))
}

func add(l *MagicList, value int) {
	newItem := Item{}
	newItem.Value = value
	if l.LastItem == nil {
		l.LastItem = &newItem
		return
	}
	tmp := l.LastItem
	for {
		if tmp.PrevItem == nil {
			tmp.PrevItem = &newItem
			return
		}
		tmp = tmp.PrevItem
	}
}

func toSlice(l *MagicList) []int {
	var result []int
	for {
		if l.LastItem == nil {

			break
		}
		result = append(result, l.LastItem.Value)
		l.LastItem = l.LastItem.PrevItem
	}
	return result
}
