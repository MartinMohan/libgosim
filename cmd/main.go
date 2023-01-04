package main

import (
	"github.com/MartinMohan/libgosim"
	"fmt"
	"github.com/spf13/pflag"
)

func main() {
	var mysignal string
	var ErrorOFF bool
	var TimeoutOFF bool
	var noloop bool
	var status bool

	pflag.StringVarP(&mysignal,"mysignal","m", "rampx1", "<rampx1 | rampx2 | error | timeout>")
	pflag.BoolVarP(&status,"status","s",false,"Return status")
	pflag.BoolVarP(&ErrorOFF, "ErrorOFF", "e", false, "Return rampx3 instead of error")
	pflag.BoolVarP(&TimeoutOFF, "TimeoutOFF", "t", false, "Return rampx4 instead of error")
	pflag.BoolVarP(&noloop, "noloop", "n", false, "Do not loop result indefinetely")
	pflag.Parse()
	//    fmt.Println("sig:", *mysignal)
	if ErrorOFF {
		mysignal="error"
		libgosim.SetErrorOFF(true)
//		fmt.Printf("ErrorOFF has been set to %v\n",ErrorOFF)
	}
	if TimeoutOFF {
		mysignal="timeout"
		libgosim.SetTimeoutOFF(true)
//		fmt.Printf("TimeoutOFF has been set to %v\n",TimeoutOFF)
	}
	if noloop {
		value,err := libgosim.Ramp(mysignal)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s %.6f\n",mysignal,value)
		return
	}
	if status {
		fmt.Printf("mysignal=%s TimeoutOFF=%v ErrorOFF=%v noloop=%v\n",mysignal,libgosim.GetTimeoutOFF(),libgosim.GetErrorOFF(),noloop)
		return
	}
	// Infinte loop
	for{
		value,err := libgosim.Ramp(mysignal)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s %.6f\n",mysignal,value)
	}
}
