package main

import (
	"fmt"
	"testing"
)

func TestColorDummy(t *testing.T) {
	randomColor := randomColor()
	for _, color := range colors {
		fmt.Println(color)
		if color == randomColor {
			t.Log("ok")
		}
		t.Skip(0)
	}
}
