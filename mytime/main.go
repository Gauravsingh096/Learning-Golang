package main

import (
	"fmt"
	"time"
)
func main() {
	fmt.Println("Hello, World!")
	present := time.Now()
	fmt.Println(present)
	fmt.Println("The time is", time.Now())
	fmt.Println("The year is", present.Year())
	fmt.Println("The month is", present.Month())
	fmt.Println("The day is", present.Day())
	fmt.Println("The hour is", present.Hour())
	fmt.Println("The minute is", present.Minute())
	fmt.Println(present.Format("Monday, January 2, 2006"))
	fmt.Println(present.Format("02/01/2006"))
	fmt.Println(present.Format("02-01-2006 15:04:05"))

	// creatingdate
	// year,time.month,date hour minute sec
	ceratedDate := time.Date(2021, time.April, 10, 23, 23, 23, 3, time.UTC)
	fmt.Println(ceratedDate)
	fmt.Println(ceratedDate.Format("01-01-2004 Monday"))

}
