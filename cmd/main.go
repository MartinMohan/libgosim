package main

import (
    "github.com/MartinMohan/libgosim"
	"fmt"
	"os"
)

func main() {
	//	fmt.Printf("os.Args[0] %s len(os.Args) %d\n",os.Args[0],len(os.Args))
	if (len(os.Args)<2){
		fmt.Printf("Usage %s <rampx1 | rampx2 | error | timeout>\n",os.Args[0])
		os.Exit(0)
	}
	for{
		value,err := libgosim.Ramp(os.Args[1])
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s %.6f\n",os.Args[1],value)
	}
}
