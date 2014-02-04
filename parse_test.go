package phonenumbers

import (
	"testing"
)

func TestBuildNationalNumberForParsing(t *testing.T) {
	var national_number string

	if national_number = buildNationalNumberForParsing("+15038884341"); national_number != "+15038884341" {
		t.Errorf("Unexpected return value : %s", national_number)
	}

	if national_number = buildNationalNumberForParsing("tel:+15038884341" + RFC3966_PHONE_CONTEXT + "example.com"); national_number != "+15038884341" {
		t.Errorf("Unexpected return value : %s", national_number)
	}

	if national_number = buildNationalNumberForParsing("tel:8884341" + RFC3966_PHONE_CONTEXT + "+1503"); national_number != "+15038884341" {
		t.Errorf("Unexpected return value : %s", national_number)
	}

	if national_number = buildNationalNumberForParsing("tel:8884341" + RFC3966_PHONE_CONTEXT + "+1503;"); national_number != "+15038884341" {
		t.Errorf("Unexpected return value : %s", national_number)
	}

	if national_number = buildNationalNumberForParsing("tel:8884341" + RFC3966_PHONE_CONTEXT + "+1503" + RFC3966_ISDN_SUBADDRESS + "test"); national_number != "+15038884341" {
		t.Errorf("Unexpected return value : %s", national_number)
	}

	if national_number = buildNationalNumberForParsing("tel:+15038884341" + RFC3966_ISDN_SUBADDRESS + "test"); national_number != "+15038884341" {
		t.Errorf("Unexpected return value : %s", national_number)
	}
}

func TestCheckRegionForParsing(t *testing.T) {
	if checkRegionForParsing("+15038884341", "+1") == false {
		t.Errorf("A valid region code is provided, no reason for a failure")
	}

	if checkRegionForParsing("+15038884341", "") == false {
		t.Errorf("A valid region code is not provided, but the phone number is prefixed with a +, no reason for a failure")
	}

	if checkRegionForParsing("1+5038884341", "") == true {
		t.Errorf("An invalid region code and phone number were provided, this should have returned false")
	}
}

func TestExtractPossibleNumber(t *testing.T) {
	var number string

	if number = extractPossibleNumber("aadf+15038884341"); number != "+15038884341" {
		t.Errorf("Unexpected return value : %s", number)
	}

	if number = extractPossibleNumber(""); number != "" {
		t.Errorf("Unexpected return value : %s", number)
	}

	if number = extractPossibleNumber("+15038884341"); number != "+15038884341" {
		t.Errorf("Unexpected return value : %s", number)
	}
}

func TestIsValidRegionCode(t *testing.T) {
	if isValidRegionCode("") {
		t.Error("Empty region code should not be considered to be valid")
	}

	if !isValidRegionCode("+1") {
		t.Error("+1 should be a valid region code")
	}
}

func TestIsViablePhoneNumber(t *testing.T) {
	if isViablePhoneNumber("") == true {
		t.Errorf("A blank phone number should not be viable")
	}

	if isViablePhoneNumber("abc") == true {
		t.Errorf("This phone number should not be viable")
	}

	if isViablePhoneNumber("+15038884341") == false {
		t.Errorf("This valid phone number should be viable")
	}
}
