package main
import("fmt")
func checkevenodd(num int)string{
	if num%2==0{
		return"even"

	}else {
		return "odd"
	}
}
func main(){
	num:=5
	result:=checkevenodd(num)
	fmt.Println(result)
}