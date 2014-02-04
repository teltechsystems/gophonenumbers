package phonenumbers

type PhoneNumber struct {
	CountryCode          *int8
	NationalNumber       *int64
	Extension            string // [0-9]+
	ItalianLeadingZero   bool
	NumberOfLeadingZeros *int
	RawInput             string
}
