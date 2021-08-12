package twoSum

type Positions struct {
	pos1 int
	pos2 int
}

func NewPositions(pos1 int) *Positions {
	return &Positions{
		pos1: pos1,
		pos2: -1,
	}
}

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)

	myMap := make(map[int]*Positions)

	for pos, val := range nums {
		p, ok := myMap[val]
		if !ok {
			myMap[val] = NewPositions(pos)
		} else {
			p.pos2 = pos
		}
	}

	for _, val := range nums {
		val2 := target - val

		if val == val2 {
			pCenter, okCenter := myMap[val]
			if okCenter && pCenter.pos2 != -1 {
				res[0] = pCenter.pos1
				res[1] = pCenter.pos2
				break
			}
		} else {
			p1, ok1 := myMap[val]
			p2, ok2 := myMap[val2]

			if ok1 && ok2 {
				res[0] = p1.pos1
				res[1] = p2.pos1
			}
		}
	}

	return res
}
