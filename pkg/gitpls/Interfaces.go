package gitpls

type CommitMessageProvider interface {
	Provide() CommitMessageQueue
}

type CommitMessageQueue func() *string
