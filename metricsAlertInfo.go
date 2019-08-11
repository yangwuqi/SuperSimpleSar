package main

//use a segment tree to significantly improve the performance of calculating metrics alert info

type NumArray struct {
	Vals       *[]int
	Sum        int
	IndexStart int
	IndexEnd   int
	Left       *NumArray
	Right      *NumArray
}

func (this *NumArray) Update(i int, val int) {
	diff := val - (*this.Vals)[i]
	treeUpdate(i, val, diff, this)
}

func (this *NumArray) SumRange(i int, j int) int {
	if i > len(*this.Vals)-1 || j < 0 {
		return 0
	}
	return getSum(i, j, this)
}

func SumNums(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func getLeftRight(root *NumArray, start int, end int, LR int, nums *[]int) *NumArray { //root是父节点
	if start > end || (root != nil && (root.IndexStart == root.IndexEnd)) {
		return nil
	} else {
		var rootTemp NumArray
		rootTemp.IndexStart = start
		rootTemp.IndexEnd = end
		rootTemp.Vals = nums
		rootTemp.Sum = SumNums((*rootTemp.Vals)[start : end+1])
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

func getSum(i, j int, root *NumArray) int {
	if root == nil {
		return 0
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
		return getSum(i, mid, root.Left) + getSum(mid+1, j, root.Right)
	}
}

func treeUpdate(i int, val int, diff int, root *NumArray) {
	if root == nil {
		return
	}
	if i < root.IndexStart || i > root.IndexEnd {
		return
	}
	(*root.Vals)[i] = val
	root.Sum += diff
	treeUpdate(i, val, diff, root.Left)
	treeUpdate(i, val, diff, root.Right)
}

func Constructor(nums []int) NumArray {
	var tempNum NumArray
	if getLeftRight(nil, 0, len(nums)-1, 0, &nums) == nil {
		tempNum.Sum = 0
	} else {
		tempNum = *getLeftRight(nil, 0, len(nums)-1, 0, &nums)
	}
	return tempNum
}
