package main

import "testing"

func TestValidPhoneNumber(t *testing.T) {

	// test empty string
	result := ValidPhoneNumber("")
	if result == "" {
		t.Errorf("ValidPhoneNumber FAILED. It allowed an empty string.")
	} else {
		t.Logf("Don't allow empty strings: PASSED")
	}

	// test a phone number less than 3 digits long
	result = ValidPhoneNumber("12")
	if result == "" {
		t.Errorf("ValidPhoneNumber FAILED. It allowed 2 digit long phone number.")
	} else {
		t.Logf("Must be 3 digits or longer: PASSED.")
	}

	result = ValidPhoneNumber("123456789012")
	if result == "" {
		t.Errorf("ValidPhoneNumber FAILED. It allowed a 12 digit long phone number.")
	} else {
		t.Logf("Must be 11 digits or less: PASSED.")
	}

	result = ValidPhoneNumber("test")
	if result == "" {
		t.Errorf("ValidPhoneNumber FAILED. It allowed a string of text")
	} else {
		t.Logf("Must be numbers only: PASSED.")
	}
}
