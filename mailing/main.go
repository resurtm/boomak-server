package mailing

import (
	"sync"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/ses"
	tj "github.com/tj/go-ses"
	"github.com/resurtm/boomak-server/database"
	cfg "github.com/resurtm/boomak-server/config"
)

var (
	clients           map[byte]*tj.Client // aws clients pool
	lock              sync.RWMutex        // protecting pool for concurrent rw
	signupMailPayload chan database.User  // workers payload channel
)

func client(uid byte) *tj.Client {
	lock.RLock()
	client, exists := clients[uid]
	lock.RUnlock()

	if !exists {
		creds := credentials.NewCredentials(&awsCredsProvider{})
		config := aws.NewConfig().
			WithRegion(cfg.Config().Mailing.AWSRegion).
			WithCredentials(creds)
		client = tj.New(ses.New(session.New(config)))

		lock.Lock()
		clients[uid] = client
		lock.Unlock()
	}
	return client
}

func SendSignupEmail(user database.User) {
	signupMailPayload <- user
}

func init() {
	clients = map[byte]*tj.Client{}
	lock = sync.RWMutex{}
	signupMailPayload = make(chan database.User, cfg.Config().Mailing.WorkerQueueSize)

	for i := byte(1); i <= cfg.Config().Mailing.WorkerCount; i++ {
		go mailWorker(i, signupMailPayload)
	}
}
