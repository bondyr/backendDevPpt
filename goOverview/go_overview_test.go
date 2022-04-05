package main

import (
	"common/common"
	"fmt"
	"testing"
)

func TestSqrtSunnyDay(t *testing.T) {
  val, _ := common.Sqrt(4)
  if val != 2 {
	  t.Error("The result should be 2")
  }
}

func TestSqrtError(t *testing.T) {
  _, err := common.Sqrt(-4)
  if err == nil {
	  t.Error("There should be an error here")
  }
}

func TestTableDriven(t *testing.T) {
  var tests = []struct {
 	number float64
 	expected float64
  }{
  	{0,  0},
  	{4,  2},
  	{9,  3},
 	{16, 4},
  }
 
  for _, testData := range tests {
	testname := fmt.Sprintf("%f,%f", testData.number, testData.expected)
	
	t.Run(testname, func(t *testing.T) {
	  val, _ := common.Sqrt(testData.number)
 	  if val != testData.expected {
  	    t.Errorf("got %f, want %f", val, testData.expected)
 	  }
	})
  }
}
