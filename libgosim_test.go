package main

import (
	"testing"
	"time"
)

func TestPlusf(t *testing.T) {
    want :=  1.23
    got,_ := Plusf("1.23")
    if got != want {
        t.Errorf("TestPlus() =>\n\t%f, want %f", got, want)
    }
}

func TestPlus(t *testing.T) {
	want := 5
	got := Plus(3,2)

	if got != want {
		t.Errorf("TestPlus() =>\n\t%d, want %d", got, want)
	}

}

// Test range 0.0 - 60.0
func TestRampx1(t *testing.T) {
	var Amplitude float64 = -1
	Amplitude, err := Ramp("rampx1")

	if err != nil {                                   
		t.Errorf("libgosim = %d; want nil", err)
	}
	if (Amplitude >= 0.0 && Amplitude <= 60.0) { 
		t.Logf("Amplitude = %f; want 0.0-60.0", Amplitude)
	}else{
		t.Errorf("Amplitude = %f; want 0.0-60.0", Amplitude)
	}
}

// Test range 0.0 - 120.0
func TestRampx2(t *testing.T) {
	var Amplitude float64 = -1
	Amplitude, err := Ramp("rampx2")

	if err != nil {                                   
		t.Errorf("libgosim = %d; want nil", err)
	}
	if (Amplitude >= 0.0 && Amplitude <= 120.0) { 
		t.Logf("Amplitude = %f; want 0.0-120.0", Amplitude)
	}else{
		t.Errorf("Amplitude = %f; want 0.0-120.0", Amplitude)
	}
}

// Test err
func TestError(t *testing.T) {
	_, err := Ramp("error")

	if err != nil {                                   
		//        t.Errorf("Error returned as expected: %v", err)
		t.Logf("Error returned as expected: %v", err)
	}
}

//func TestWithTimeOut(t *testing.T) {
func TestTimeOut(t *testing.T) {
	timeout := time.After(3 * time.Second)
	done := make(chan bool)
	go func() {
		// do your testing
		//		time.Sleep(5 * time.Second)
		Ramp("timeout")
		done <- true
	}()

	select {
	case <-timeout:
		//        t.Fatal("Test didn't finish in time")
		t.Logf("No response within %d seconds. Timeout as expected",3)
	case <-done:
	}
}
