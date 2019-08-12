package main

//use a segment tree to significantly improve the performance of calculating metrics alert info
//basic structure, to be continued

type NumArray struct {
	Vals       *[][]int
	Sum        []int
	IndexStart int
	IndexEnd   int
	Left       *NumArray
	Right      *NumArray
}

func (this *NumArray) Update(i int, val []int) {
	var diff []int
	for j := 0; j < len(val); j++ {
		change := val[j] - (*this.Vals)[i][j]
		diff = append(diff, change)
	}
	treeUpdate(i, val, diff, this)
}

func (this *NumArray) SumRange(i int, j int) []int {
	if i > len(*this.Vals)-1 || j < 0 {
		return []int{}
	}
	return getSum(i, j, this)
}

func SumNums(nums [][]int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	var sum []int
	metricsNum := len(nums[0])
	for index := 0; index < metricsNum; index++ {
		sum = append(sum, 0)
	}
	for index1 := 0; index1 < len(nums); index1++ {
		for index2 := 0; index2 < metricsNum; index2++ {
			sum[index2] += nums[index1][index2]
		}
	}
	return sum
}

func getLeftRight(root *NumArray, start int, end int, LR int, nums *[][]int) *NumArray { //root是父节点
	if start > end || (root != nil && (root.IndexStart == root.IndexEnd)) {
		return nil
	} else {
		var rootTemp NumArray
		rootTemp.IndexStart = start
		rootTemp.IndexEnd = end
		rootTemp.Vals = nums
		rootTemp.Sum = SumNums((*rootTemp.Vals)[start : end+1])
		//fmt.Println("rootTemp.Sum: ",rootTemp.Sum)//!!!!!!!!!
		if LR == -1 && root != nil { //-1==left
			root.Left = &rootTemp
		} else if LR == 1 && root != nil { //1==right
			root.Right = &rootTemp
		}
		mid := (start + end) / 2
		getLeftRight(&rootTemp, start, mid, -1, nums)
		getLeftRight(&rootTemp, mid+1, end, 1, nums)
		return &rootTemp
	}
}

func getSum(i, j int, root *NumArray) []int {
	if root == nil {
		return []int{}
	}
	if i == root.IndexStart && j == root.IndexEnd {
		return root.Sum
	}
	mid := (root.IndexStart + root.IndexEnd) / 2
	if j <= mid {
		return getSum(i, j, root.Left)
	} else if i >= mid+1 {
		return getSum(i, j, root.Right)
	} else {
		return sumSilces(getSum(i, mid, root.Left), getSum(mid+1, j, root.Right))
	}
}

func sumSilces(slice1, slice2 []int) []int {
	if len(slice1) != 0 && len(slice2) != 0 && len(slice1) != len(slice2) {
		panic("slice1's length is not equal to slice2's length")
	}
	var result []int
	if len(slice1) != 0 && len(slice2) != 0 {
		for index := 0; index < len(slice1); index++ {
			result = append(result, slice1[index]+slice2[index])
		}
	} else {
		if len(slice1) != 0 {
			return slice1
		} else {
			return slice2
		}
	}

	return result
}

func treeUpdate(i int, val []int, diff []int, root *NumArray) {
	if root == nil {
		return
	}
	if i < root.IndexStart || i > root.IndexEnd {
		return
	}
	(*root.Vals)[i] = val
	for index := 0; index < len(root.Sum); index++ {
		(*root).Sum[index] += diff[index]
	}
	treeUpdate(i, val, diff, root.Left)
	treeUpdate(i, val, diff, root.Right)
}

func segmentTreeConstructor(nums [][]int) NumArray {
	if len(nums) == 0 {
		panic("the data for the segmentTree is nil!")
	}
	var result NumArray
	if resultTemp := getLeftRight(nil, 0, len(nums)-1, 0, &nums); resultTemp == nil {
		var sumTemp []int
		for index := 0; index < len(nums[0]); index++ {
			sumTemp = append(sumTemp, 0)
		}
		result.Sum = sumTemp
	} else {
		result = *resultTemp
	}
	return result
}
