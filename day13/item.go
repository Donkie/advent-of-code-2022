package main

type Item struct {
	children []*Item
	value    int
}

func newItem(val int) *Item {
	item := new(Item)
	item.value = val
	return item
}

type Pair struct {
	p1 Item
	p2 Item
}

func (pair Pair) IsOrdered() (ordered bool) {
	return
}

func GetSumOfOrderedPairIndices(pairs *[]Pair) (sum int) {
	return
}
