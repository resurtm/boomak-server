package common

type AuthEntry struct {
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
}

type TestEmail struct {
	RecipientEmail string `json:"recipient_email" mapstructure:"recipient_email"`
	TestString     string `json:"test_string" mapstructure:"test_string"`
}
