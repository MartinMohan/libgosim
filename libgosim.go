package libgosim

import (
	"time"
	"errors"
	//	"fmt"
	//	"os"
)
var ErrorOFF bool = false
var TimeoutOFF bool = false

func SetErrorOFF(onoff bool)(){
	ErrorOFF=onoff
}
func SetTimeoutOFF(onoff bool)(){
	TimeoutOFF=onoff
}
func GetErrorOFF()(bool){
	return ErrorOFF
}
func GetTimeoutOFF()(bool){
	return TimeoutOFF
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

		case "error": // error
		if ErrorOFF{
			return valuef*3,nil
		}
		return valuef,errors.New("Error Device")

		case "timeout": // timeout
		//		time.Sleep(1000)
		if TimeoutOFF{
			return valuef*4,nil
		}
		time.Sleep(9999999999999)
		return valuef, nil


		default: // error
		return valuef, errors.New("Unknown Device")
	}
	return valuef, nil
}
