package main

import (
	"testing"
)

func TestGetNCoolestPresents(t *testing.T) {
	t.Run("Test 1 in Parallel", func(t *testing.T) {
		t.Parallel()
		presents := []Present{
			{Value: 5, Size: 1},
			{Value: 4, Size: 5},
			{Value: 5, Size: 1},
			{Value: 5, Size: 2},
		}
		resStruct := []Present{
			{Value: 5, Size: 1},
			{Value: 5, Size: 1},
			{Value: 5, Size: 2},
		}
		res, _ := getNCoolestPresents(presents, 4)
		for i := range res {
			if res[i] != resStruct[i] {
				t.Error("Result was incorrect")
			}
		}
	})
	t.Run("Test 2 in Parallel", func(t *testing.T) {
		t.Parallel()
		presents := []Present{
			{Value: 6, Size: 1},
			{Value: 4, Size: 5},
			{Value: 5, Size: 1},
			{Value: 5, Size: 2},
		}
		resStruct := []Present{
			{Value: 6, Size: 1},
		}
		res, _ := getNCoolestPresents(presents, 4)
		for i := range res {
			if res[i] != resStruct[i] {
				t.Error("Result was incorrect")
			}
		}
	})
}
