package main
import ("fmt")
func sumarray(arr[5]int)int{
	sum:=0
	for _, value :=range arr{
		sum+=value
	}
	return sum
}
func  main(){
	numbers:=[5]int{1,2,3,4,5}
	result:=sumarray(numbers)
	fmt.Println(result)
}