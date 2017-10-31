package types

type TestEmail struct {
	RecipientEmail string `json:"recipient_email" mapstructure:"recipient_email"`
	TestString     string `json:"test_string" mapstructure:"test_string"`
}
