package main
import("fmt")
func getweekday(day int )string{
	switch day{
	case 1:
		return "monday"
	case 2:
		return "tuesday"
	case 3:
		return "wednesday"
	case 4:
		return "thursady"
	case 5:
		return "friday"
	case 6:
		return "saturday"
	case 7:
		return "sunday"
	default:
		return "invalid day "
	}
}
func main(){
	day := 3
	weekdayname := getweekday(day)
	fmt.Println(weekdayname)
}