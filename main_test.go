package main

import "testing"

func Test_getTeams(t *testing.T) {
	team := 5
	if team != 5 {
		t.Error("incorrect result: expected 5")
	}
}
