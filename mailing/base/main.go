package base

const (
	TestMailJob           = iota
	EmailVerifyMailJob
	SignupFinishedMailJob
)

type MailJob struct {
	Kind    byte
	Payload interface{}
}

var MailJobsQueue chan MailJob

func EnqueueTestMailJob(data interface{}) {
	MailJobsQueue <- MailJob{Kind: TestMailJob, Payload: data}
}

func EnqueueEmailVerifyMailJob(user interface{}) {
	MailJobsQueue <- MailJob{Kind: EmailVerifyMailJob, Payload: user}
}

func EnqueueSignupFinishedMailJob(user interface{}) {
	MailJobsQueue <- MailJob{Kind: SignupFinishedMailJob, Payload: user}
}
