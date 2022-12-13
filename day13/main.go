package main

import "log"

func main() {
	pairs := ParseInputPairsFile("input.txt")
	metric := GetSumOfOrderedPairIndices(*pairs)

	log.Printf("Part 1 - Sum of ordered pair indices: %d", metric)

	list := newPacketListFromPairs(*pairs)
	log.Printf("Part 2 - Decoder key: %d", list.GetDecoderKey())
}
