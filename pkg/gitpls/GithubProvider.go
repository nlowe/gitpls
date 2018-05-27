package gitpls

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GithubCommitMessageProvider struct {
	ctx    context.Context
	client *github.Client
}

func NewGithubProvider() *GithubCommitMessageProvider {
	ctx := context.Background()

	var tc *http.Client

	if key, present := os.LookupEnv("GITHUB_TOKEN"); present {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: key},
		)
		tc = oauth2.NewClient(ctx, ts)
	}

	return &GithubCommitMessageProvider{
		ctx:    ctx,
		client: github.NewClient(tc),
	}
}

func (p *GithubCommitMessageProvider) Provide() CommitMessageQueue {
	done := false
	messages := make([]*string, 0)

	// Keep trying to populate while github gives us (be lazy and don't worry about pages)
	events, resp, err := p.client.Activity.ListEvents(p.ctx, nil)
	fmt.Printf("%d requests remaining in rate (resets at %s)\n", resp.Rate.Remaining, resp.Rate.Reset)

	if err != nil {
		done = true
		fmt.Printf("Failed to fetch new events! %s\n", err)

		return nil
	}

	for _, raw := range events {
		parsed, err := raw.ParsePayload()
		if err != nil {
			fmt.Printf("Failed to parse event! %s\n", err)
		}

		switch ev := parsed.(type) {
		case *github.PushEvent:
			for _, c := range ev.Commits {
				messages = append(messages, c.Message)
			}
		}
	}

	return func() *string {
		if done || len(messages) == 0 {
			done = true
			return nil
		}

		var ret *string
		ret, messages = messages[0], messages[1:]

		return ret
	}
}
