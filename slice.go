package main
import("fmt")
func maxinslice(nums[]int)int{
	if len(nums)==0{
		return 0

	}
	max := nums[0]
	for _, value:=range nums{
		if value > max{
			max = value
		}
	} 
	return max
}
func main(){
	numbers := []int{3,5,2,1,8}
	result := maxinslice(numbers)
	fmt.Println(result)
}

