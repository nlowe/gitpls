package gitpls

import (
	"context"
	"net/http"
	"os"
	"strconv"

	"github.com/google/go-github/github"
	"github.com/gregjones/httpcache"
	"golang.org/x/oauth2"

	log "github.com/sirupsen/logrus"
)

type GithubCommitMessageProvider struct {
	ctx    context.Context
	client *github.Client
}

func NewGithubProvider() *GithubCommitMessageProvider {
	ctx := context.Background()

	cache := httpcache.NewMemoryCacheTransport()
	tc := &http.Client{Transport: cache}
	if key, present := os.LookupEnv("GITHUB_TOKEN"); present {
		cache.Transport = &oauth2.Transport{
			Source: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: key}),
		}
	}

	return &GithubCommitMessageProvider{
		ctx:    ctx,
		client: github.NewClient(tc),
	}
}

func (p *GithubCommitMessageProvider) Provide() (CommitMessageQueue, int) {
	done := false
	messages := make([]*string, 0)

	// Keep trying to populate while github gives us (be lazy and don't worry about pages)
	delay := 60
	listOptions := github.ListOptions{Page: 1, PerPage: 30}
	for {
		events, resp, err := p.client.Activity.ListEvents(p.ctx, &listOptions)
		log.WithFields(log.Fields{
			"remainingRequests": resp.Rate.Remaining,
			"resetAt":           resp.Rate.Reset,
		}).Debug("Requested more events")
		delay, _ := strconv.Atoi(resp.Response.Header.Get("X-Poll-Interval"))

		if err != nil {
			done = true
			log.WithError(err).Error("Failed to fetch new events")

			return nil, delay
		}

		for _, raw := range events {
			parsed, err := raw.ParsePayload()
			if err != nil {
				log.WithError(err).Error("Failed to parse event")
			}

			switch ev := parsed.(type) {
			case *github.PushEvent:
				for _, c := range ev.Commits {
					messages = append(messages, c.Message)
				}
			}
		}

		if resp.NextPage == 0 {
			break
		}

		listOptions.Page = resp.NextPage
	}

	return func() *string {
		if done || len(messages) == 0 {
			done = true
			return nil
		}

		var ret *string
		ret, messages = messages[0], messages[1:]

		return ret
	}, delay
}
