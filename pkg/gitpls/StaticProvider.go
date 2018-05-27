package gitpls

type StaticProvider struct{}

func NewStaticProvider() *StaticProvider {
	return &StaticProvider{}
}

func (*StaticProvider) Provide() CommitMessageQueue {
	messages := []string{
		"a",
		"b",
		"c",
	}

	i := 0

	return func() *string {
		if i == len(messages) {
			return nil
		}

		ret := messages[i]
		i++
		return &ret
	}
}

func foo() {

}
