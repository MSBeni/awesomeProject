package timenow

import ( "fmt"; "time" ) 

// Calculating time and speed
func Timenowcalc(){
	// var(
	// 	now	time.Time
	// 	speed int
	// )
	speed, now := 100, time.Now()
	fmt.Println("The speed is: ", speed, "The time is: ", now)
} 