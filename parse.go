package phonenumbers

import (
	"bytes"
	// "fmt"
	"regexp"
	"strings"
)

var (
	plus_chars_pattern         = regexp.MustCompile("[" + PLUS_CHARS + "]")
	unwanted_end_char_pattern  = regexp.MustCompile("[[\\P{N}&&\\P{L}]&&[^#]]+$")
	valid_phone_number_pattern = regexp.MustCompile(VALID_PHONE_NUMBER + "(?:" + EXTN_PATTERNS_FOR_PARSING + ")?")
	valid_start_char_pattern   = regexp.MustCompile("[+\uFF0B\\p{Nd}]")
)

func Parse(number_to_parse, default_region string, keep_raw_input, check_region bool) (*PhoneNumber, error) {
	var (
		national_number string
	)

	if number_to_parse == "" {
		return nil, &NumberParseError{"The phone number supplied was empty."}
	} else if len(number_to_parse) > MAX_INPUT_STRING_LENGTH {
		return nil, &NumberParseError{"The string supplied was too long to parse."}
	}

	if national_number = buildNationalNumberForParsing(number_to_parse); !isViablePhoneNumber(national_number) {
		return nil, &NumberParseError{"The string supplied did not seem to be a phone number."}
	}

	if check_region && !checkRegionForParsing(national_number, default_region) {
		return nil, &NumberParseError{"Missing or invalid default region."}
	}

	return nil, nil
}

func buildNationalNumberForParsing(number string) string {
	buff := bytes.NewBuffer([]byte{})

	if phone_context_index := strings.Index(number, RFC3966_PHONE_CONTEXT); phone_context_index > 0 {
		phone_context_start := phone_context_index + len(RFC3966_PHONE_CONTEXT)

		// If the phone context contains a phone number prefix, we need to capture it, whereas domains
		// will be ignored.
		if string(number[phone_context_start]) == PLUS_SIGN {
			// Additional parameters might follow the phone context. If so, we will remove them here
			// because the parameters after phone context are not important for parsing the
			// phone number.
			phone_context_end := strings.Index(number[phone_context_start:], ";")

			if phone_context_end > 0 {
				buff.WriteString(number[phone_context_start : phone_context_start+phone_context_end])
			} else {
				buff.WriteString(number[phone_context_start:])
			}
		}

		// Now append everything between the "tel:" prefix and the phone-context. This should include
		// the national number, an optional extension or isdn-subaddress component.
		buff.WriteString(number[strings.Index(number, RFC3966_PREFIX)+len(RFC3966_PREFIX) : phone_context_index])
	} else {
		// Extract a possible number from the string passed in (this strips leading characters that
		// could not be the start of a phone number.)
		buff.WriteString(extractPossibleNumber(number))
	}

	national_number := buff.String()

	// Delete the isdn-subaddress and everything after it if it is present. Note extension won't
	// appear at the same time with isdn-subaddress according to paragraph 5.3 of the RFC3966 spec,
	if isdn_index := strings.Index(national_number, RFC3966_ISDN_SUBADDRESS); isdn_index > 0 {
		national_number = national_number[:isdn_index]
	}

	// If both phone context and isdn-subaddress are absent but other parameters are present, the
	// parameters are left in nationalNumber. This is because we are concerned about deleting
	// content from a potential number string when there is no strong evidence that the number is
	// actually written in RFC3966.

	return national_number
}

// Checks to see that the region code used is valid, or if it is not valid, that the number to
// parse starts with a + symbol so that we can attempt to infer the region from the number.
// Returns false if it cannot use the region provided and the region cannot be inferred.
func checkRegionForParsing(number_to_parse, default_region string) bool {
	var loc []int

	if !isValidRegionCode(default_region) {
		if loc = plus_chars_pattern.FindStringIndex(number_to_parse); loc == nil {
			return false
		}

		// We MUST validate that the phone number STARTS with the discovered +
		if len(number_to_parse) == 0 || loc[0] != 0 {
			return false
		}
	}

	return true
}

func extractPossibleNumber(number string) string {
	var loc []int

	if loc = valid_start_char_pattern.FindStringIndex(number); loc == nil {
		return ""
	}

	number = number[loc[0]:]

	// Remove trailing non-alpha non-numerical characters.
	if loc = unwanted_end_char_pattern.FindStringIndex(number); loc == nil {
		return number
	}

	return number[:loc[0]]
}

func isValidRegionCode(region_code string) bool {
	return region_code != "" // Check for supported region codes :1045 in original library
}

func isViablePhoneNumber(number string) bool {
	if len(number) < MIN_LENGTH_FOR_NSN {
		return false
	}

	return valid_phone_number_pattern.MatchString(number)
}
