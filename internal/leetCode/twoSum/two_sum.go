package twoSum

type Positions struct {
	pos1 int
	pos2 int
}

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)

	myMap := make(map[int]*Positions)

	for pos, val := range nums {
		p, ok := myMap[val]
		if !ok {
			myMap[val] = &Positions{pos1: pos}
		} else {
			p.pos2 = pos
		}

	}

	for i := 1; i <= target; i++ {
		if i == target-i {
			pCenter, okCenter := myMap[i]
			if okCenter {
				res[0] = pCenter.pos1
				res[1] = pCenter.pos2
				break
			}
		}

		p1, ok1 := myMap[i]
		p2, ok2 := myMap[target-i]

		if ok1 && ok2 {
			res[0] = p1.pos1
			res[1] = p2.pos1
		}
	}

	return res
}
