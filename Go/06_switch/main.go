package main

func main() {
	// i:= 0

	// switch i {
	// case 1:
	// 	println("True")
	// case 0:
	// 	println("False")
	// default:
	// 	println("Invalid Input")
	// }


// multiple case

// switch time.Now().Weekday()	 {

// 	case time.Saturday, time.Sunday:{
// 		println("It's the weekend")
// 	}

// 	default:{
// 		println("It's a weekday")
// 	}
// }


whoIAm := func (i interface {})  {
	switch i.(type) {
	case int:
		println("It's an int")
	case string:
		println("It's a string")
	case bool:
		println("It's a bool")
	default:
		println("Other type")
	}
}


whoIAm(true)
}
