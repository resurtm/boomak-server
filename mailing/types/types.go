package types

const (
	TestMailJob           = iota
	EmailVerifyMailJob    = iota
	SignupFinishedMailJob = iota
)

type MailJob struct {
	Kind    byte
	Payload interface{}
}
