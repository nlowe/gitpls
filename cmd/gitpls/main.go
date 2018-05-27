package main

import (
	"fmt"
	"regexp"
	"time"

	"github.com/nlowe/gitpls/pkg/gitpls"
	log "github.com/sirupsen/logrus"
)

var patterns = []*regexp.Regexp{
	regexp.MustCompile("(?i)pl(?:ea)?s(?:e)?"),
}

func processPush(msg *string) bool {
	for _, r := range patterns {
		if r.MatchString(*msg) {
			fmt.Printf("LETS GO: %s\n", *msg)
			return false
		}
	}

	return true
}

func main() {
	log.Info("Hello, World!")

	provider := gitpls.NewGithubProvider()
	for {
		gen := provider.Provide()

		for {
			msg := gen()
			if msg == nil || !processPush(msg) {
				break
			}
		}

		time.Sleep(2 * time.Second)
	}
}
