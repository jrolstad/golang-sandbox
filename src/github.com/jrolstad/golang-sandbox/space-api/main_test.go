package main

import "testing"

func Test_GetAstronauts_NoInputs_ReturnsAstronautData(t *testing.T) {
	result := GetAstronauts()

	AssertDataReceived(t, result)
	AssertAstronautCountIsLargerThanZero(t, result)
}

func AssertDataReceived(t *testing.T, data people) {
	if data == (people{}) {
		t.Error("No data received")
	}
}

func AssertAstronautCountIsLargerThanZero(t *testing.T, data people) {
	if data.Number <= 0 {
		t.Error("No Astronaut data received")
	}
}
