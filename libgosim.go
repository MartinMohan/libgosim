package main

import (
	"time"
	"errors"
	"fmt"
	"os"
)

func Plus(a int, b int) int {
	return a + b
}

func Plusf(mystring string) (float64, error) {
	fmt.Printf("In Plusf %s",mystring)
	var valuef float64 = 1.23
	return valuef,nil 
}

// Ramp
func Ramp(dev string) (float64, error) {
	now := time.Now()
	var valuef float64 = 1.23
	valuef = float64(now.Second()) + float64(now.Nanosecond())/1e9

	switch dev {
		case "rampx1": // 0-59.999999 = ramp
		return valuef, nil

		case "rampx2": // 0-119.999999 = rampx2
		return valuef*2, nil

		case "timeout": // timeout
		//		time.Sleep(1000)
		time.Sleep(9999999999999)
		return valuef, nil

		case "error": // error
		return valuef,errors.New("Error Device")

		default: // error
		return valuef, errors.New("Unknown Device")
	}
	return valuef, nil
}

func main() {
	//	fmt.Printf("os.Args[0] %s len(os.Args) %d\n",os.Args[0],len(os.Args))
	if (len(os.Args)<2){
		fmt.Printf("Usage %s <rampx1 | rampx1 | error | timeout>\n",os.Args[0])
		os.Exit(0)
	}
	for{
		value,err := Ramp(os.Args[1])
		if err != nil {
			panic(err)
		}
		fmt.Printf("rampx1 %.6f\n",value)
	}
}
