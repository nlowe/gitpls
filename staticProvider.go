package main

type staticProvider struct{}

func newStaticProvider() *staticProvider {
	return &staticProvider{}
}

func (*staticProvider) Provide() commitMessageQueue {
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
