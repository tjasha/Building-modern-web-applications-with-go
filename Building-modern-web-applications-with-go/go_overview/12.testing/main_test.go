package main

import "testing"

//file has to have the same name as tested file with _test in the end
// we run test with "go test -v" command
// we can run "go test -cover" to see how much of the functionality is tested
// command "go test -coverprofile=coverage.out && go tool cover -html=coverage.out" should create html report of the coverage an open it in the browser

//case1: 

//function should start with Test, otherwise will be ignored
func TestDivide(t *testing.T){ //build in package testing
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

func TestBadDivide(t *testing.T){ //build in package testing
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Did not get an error when we should not have")
	}
}


//case2 - table tests

var tests = []struct {
	name string
	dividend float32
	divisor float32
	expected float32
	isErr bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 50, 10, 5, false},
	{"fraction", -1.0, -777.0, 0.0012870013, false},

}

func TestDivision(t *testing.T){
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}

