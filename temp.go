package main

import "sort"

func beautifulSubarrays(nums []int) int64 {
	// subtracting 2^k doesn't move or the alter the bits from their position
	// just bit at position k turn to 0
	var count int64 = 0
	arrLen := len(nums)
	for i := 0; i < arrLen; i++ {
		for j := i; j < arrLen; j++ {
			// Fix the subarray
			flag := true
			// iterate over all the 32 bit
			// count if there are even number of bits in each position
			for bit := 0; bit < 32; bit++ {
				if !(bitCountAtIndex(nums, i, j, bit)&1 == 0) {
					flag = false
					break
				}
			}
			if flag {
				count++
			}
		}
	}
	return count
}

func bitCountAtIndex(nums []int, start, end, bitPos int) int {
	mask := 1 << bitPos
	bitCount := 0
	for i := start; i <= end; i++ {
		if nums[i]&mask > 0 {
			bitCount++
		}
	}
	return bitCount
}

func passThePillow(n int, time int) int {
	sig := 1
	num := 1
	for t := 1; t <= time; t++ {
		num += sig
	}
	return -1
}

func coloredCells(n int) int64 {
	deltaX := []int{0, 1, -1, 0}
	deltaY := []int{1, 0, 0, -1}

	type Point struct {
		X int
		Y int
	}
	pointSet := make(map[Point]struct{})
	q := []Point{
		{
			X: 0,
			Y: 0,
		},
	}
	// make initial point as visited
	var count int64 = 0
	time := 0
	for time < n {
		levelLen := len(q)
		for i := 0; i < levelLen; i++ {
			curr := q[0]
			q = q[1:]
			if _, ok := pointSet[curr]; !ok {
				pointSet[curr] = struct{}{}
				count++
			}
			// generate next set of point
			for k := 0; k < 4; k++ {
				nx := curr.X + deltaX[k]
				ny := curr.Y + deltaY[k]
				auxNewPoint := Point{
					X: nx,
					Y: ny,
				}
				q = append(q, auxNewPoint)
			}
		}
		time++
	}
	return count
}

func maximizeGreatness(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := 0
	for _, val := range nums {
		if val > nums[res] {
			res++
		}
	}
	return res
}

func findScore(nums []int) int64 {
	type ValIndexPair struct {
		First  int
		Second int
	}
	numsLen := len(nums)
	pairs := make([]ValIndexPair, numsLen)
	for idx, val := range nums {
		pairs[idx] = ValIndexPair{
			First:  val,
			Second: idx,
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].First == pairs[j].First {
			return pairs[i].Second < pairs[j].Second
		}
		return pairs[i].First < pairs[j].First
	})
	markedIndex := make(map[int]struct{})
	score := int64(0)
	for i := 0; i < numsLen; i++ {
		p := pairs[i]
		if _, ok := markedIndex[p.Second]; ok {
			continue
		}
		score += int64(p.First)
		markedIndex[p.Second] = struct{}{}
		markedIndex[p.Second+1] = struct{}{}
		markedIndex[p.Second-1] = struct{}{}
	}
	return score
}
