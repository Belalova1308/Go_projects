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
		c := 3
		resStruct := []Present{
			{Value: 5, Size: 2},
			{Value: 5, Size: 1},
		}
		res := grabPresents(presents, c)
		for i := range res {
			if res[i] != resStruct[i] {
				t.Error("Result was incorrect")
			}
		}
	})
	t.Run("Test 2 in Parallel", func(t *testing.T) {
		t.Parallel()
		presents := []Present{
			{Value: 5, Size: 1},
			{Value: 4, Size: 5},
			{Value: 5, Size: 1},
			{Value: 5, Size: 2},
		}
		c := 20
		resStruct := []Present{
			{Value: 5, Size: 2},
			{Value: 5, Size: 1},
			{Value: 5, Size: 1},
			{Value: 4, Size: 5},
		}
		res := grabPresents(presents, c)
		for i := range res {
			if res[i] != resStruct[i] {
				t.Error("Result was incorrect")
			}
		}
	})

	t.Run("Test 3 in Parallel", func(t *testing.T) {
		t.Parallel()
		presents := []Present{
			{Value: 5, Size: 1},
			{Value: 4, Size: 5},
			{Value: 5, Size: 1},
			{Value: 5, Size: 2},
		}
		c := 0
		resStruct := []Present{}
		res := grabPresents(presents, c)
		for i := range res {
			if res[i] != resStruct[i] {
				t.Error("Result was incorrect")
			}
		}
	})
}
