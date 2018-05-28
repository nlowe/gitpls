package main

import (
	"regexp"
	"time"

	"github.com/nlowe/gitpls/pkg/gitpls"
	log "github.com/sirupsen/logrus"
)

var patterns = []*regexp.Regexp{
	regexp.MustCompile("(?i)pl(?:ea)?s(?:e)?"),
}

var truncator = &gitpls.MessageTruncator{MaxLength: 280}
var foundCount = 0

func processPush(msg *string) bool {
	for _, r := range patterns {
		if r.MatchString(*msg) {
			foundCount++
			log.WithFields(log.Fields{
				"message": *msg,
				"truncatedMessage": truncator.Truncate(msg)
			}).Warning("Found a matching commit")
			return false
		}
	}

	return true
}

func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("Hello, World!")

	provider := gitpls.NewGithubProvider()
	for {
		gen, delay := provider.Provide()

		i := 0
		for {
			msg := gen()
			if msg == nil {
				break
			}

			processPush(msg)

			i++
		}

		log.WithFields(log.Fields{
			"count": i,
			"delay": delay,
			"found": foundCount,
		}).Infof("Processed Commits. Next tick in %d seconds", delay)

		time.Sleep(time.Duration(delay) * time.Second)
	}
}
