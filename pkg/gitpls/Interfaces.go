package gitpls

type CommitMessageProvider interface {
	Provide() (CommitMessageQueue, int)
}

type CommitMessageQueue func() *string
