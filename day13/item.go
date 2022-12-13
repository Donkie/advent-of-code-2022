package main

// An item
// Can either be:
// * A list of sub-items
// * A value item
// If the value field is -1, it's a list of sub-items, otherwise it's a value item.
type Item struct {
	children []*Item
	value    int
}

func (i Item) IsValue() bool {
	return i.value != -1
}

// Creates a new item with a specified value
// Put to -1 to make it be a list of sub-items
func newItem(val int) *Item {
	item := new(Item)
	item.value = val
	return item
}

func valueItemToList(valItem *Item) *Item {
	newItem := newItem(-1)
	newItem.children = []*Item{valItem}
	return newItem
}

// Returns:
// 1  : Right order
// -1 : Wrong order
// 0  : No decision
func compare(left *Item, right *Item) int {
	if left.IsValue() && right.IsValue() {
		if left.value == right.value {
			return 0
		} else if left.value < right.value {
			return 1
		} else {
			return -1
		}
	} else if !left.IsValue() && !right.IsValue() {
		for i, item := range left.children {
			if i >= len(right.children) {
				// Right ran out of children
				return -1
			}

			cmp := compare(item, right.children[i])
			if cmp != 0 {
				return cmp
			}
		}

		if len(left.children) < len(right.children) {
			// Left ran out of children
			return 1
		} else {
			return 0
		}
	} else if left.IsValue() && !right.IsValue() {
		// Convert left to a list item
		return compare(valueItemToList(left), right)
	} else {
		// Convert right to a list item
		return compare(left, valueItemToList(right))
	}
}

type Pair struct {
	p1 Item
	p2 Item
}

func (pair Pair) IsOrdered() (ordered bool) {
	return compare(&pair.p1, &pair.p2) == 1
}

func GetSumOfOrderedPairIndices(pairs []Pair) (sum int) {
	for idx, pair := range pairs {
		if pair.IsOrdered() {
			sum += idx + 1
		}
	}
	return
}
