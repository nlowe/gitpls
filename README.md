# gitpls [![Build Status](https://travis-ci.org/nlowe/gitpls.svg?branch=master)](https://travis-ci.org/nlowe/gitpls) [![Twitter Follow](https://img.shields.io/twitter/follow/gitpls.svg?style=social&label=Follow)](https://twitter.com/gitpls)

Unfiltered commit messages from well-mannered developers
using GitHub's API. Inspired by
[Developers Swearing](https://twitter.com/gitlost).

## Building

Easiest way to build is to use the makefile:

```bash
make
```

Then, run the binary at `./dist/gitpls`.

## Usage

You need a vew environment variables set:

* `GITHUB_TOKEN`: An OAuth2 token for a github user. Required, otherwise you'll run out of requests after 6 minutes by default due to rate limiting.
* `TWITTER_KEY`: Consumer Key for twitter application
* `TWITTER_SECRET`: Consumer secret for twitter application
* `TWITTER_ACCESS_TOKEN`: Access token for twitter account to tweet as
* `TWITTER_ACCESS_SECRET`: Access token secret for twitter account to tweet as

The app will poll the github API and pull back as many pages
of events as github will let it. Currently this is 10 pages of
30 events each. The app will then look for all commit messages
for any that match any regular expressions. For each matching
commit, the message is truncated to 280 characters and
tweeted (currently it's just printed to `stdout`).

After each set of pages is pulled, the application will wait
for the interval defined by github (usually 60 seconds during
normal load).


## License

Copyright 2018 Nathan Lowe

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.