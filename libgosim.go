package libgosim

import (
	"time"
	"errors"
//	"fmt"
//	"os"
)

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
