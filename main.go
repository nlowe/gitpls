package main

import (
	"context"
	"fmt"
	"regexp"

	"github.com/google/go-github/github"
)

var patterns = []*regexp.Regexp{
	regexp.MustCompile("(?i)pl(?:ea)?s(?:e)?"),
}

func processPush(ev *github.PushEvent) bool {
	if len(ev.Commits) > 0 {
		for _, c := range ev.Commits {
			for _, r := range patterns {
				if r.MatchString(c.GetMessage()) {
					fmt.Printf("LETS GO: %s\n", c.GetMessage())
					return false
				}
			}
		}
	}

	return true
}

func main() {
	client := github.NewClient(nil)

	opt := &github.ListOptions{PerPage: 30}
search:
	for {
		events, resp, err := client.Activity.ListEvents(context.Background(), opt)

		if err != nil {
			panic(err)
		}

		for _, raw := range events {
			parsed, err := raw.ParsePayload()
			if err != nil {
				panic(err)
			}

			switch ev := parsed.(type) {
			case *github.PushEvent:
				if !processPush(ev) {
					break search
				}
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage

		fmt.Printf("%d requests remaining in rate (resets at %s)\n", resp.Rate.Remaining, resp.Rate.Reset)
	}
}
