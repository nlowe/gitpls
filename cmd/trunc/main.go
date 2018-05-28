package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nlowe/gitpls/pkg/gitpls"
)

var sep string

func main() {
	flag.StringVar(&sep, "-key", "please", "Separator key")
	flag.Parse()

	t := &gitpls.MessageTruncator{MaxLength: 280}

	src, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	sourceString := string(src)

	fmt.Println(t.Truncate(&sourceString, sep))
}
