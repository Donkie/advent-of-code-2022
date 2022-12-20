package lib

func InsertAtIndex[V any](slice []V, idx int, value V) []V {
	slice = append(slice[:idx+1], slice[idx:]...)
	slice[idx] = value
	return slice
}

func RemoveAtIndex[V any](slice []V, idx int) []V {
	return append(slice[:idx], slice[idx+1:]...)
}
