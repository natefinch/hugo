---
title: "Contributing to Hugo"
date: "2013-07-01"
aliases: ["/doc/contributing/", "/meta/contributing/"]
groups: ["community"]
groups_weight: 30
---

We welcome all contributions. Feel free to pick something from the roadmap
or contact [spf13](http://spf13.com) about what may make sense
to do next. Go ahead and fork the project and make your changes.  *We encourage pull requests to discuss code changes.*

When you're ready to create a pull request, be sure to:

  * Have test cases for the new code.  If you have questions about how to do it, please ask in your pull request.
  * Run `go fmt`
  * Squash your commits into a single commit.  `git rebase -i`.  It's okay to force update your pull request.
  * Make sure `go test ./...` passes, and go build completes.  Our Travis CI loop will catch most things that are missing.  The exception: Windows.  We run on windows from time to time, but if you have access please check on a Windows machine too.

## Contribution Overview

1. Fork Hugo from https://github.com/spf13/hugo
2. git clone your repo into $GOPATH/src/github.com/spf13/hugo (this makes the go tool happy)
3. run 'go get' from that directory (this gets the dependencies)
4. Create your feature branch (`git checkout -b my-new-feature`)
5. Commit your changes (`git commit -am 'Add some feature'`)
6. Commit passing tests to validate changes.
7. Run `go fmt`
8. Squash commits into a single (or logically grouped) commits (`git rebase -i`)
9. Push to the branch (`git push origin my-new-feature`)
10. Create new Pull Request


# Building from source

## Clone locally (for contributors):

    mkdir -a $GOPATH/src/github.com/spf13/hugo
    cd $GOPATH/src/github.com/spf13/hugo
    git clone https://github.com/spf13/hugo .
    go get

## Running Hugo

Make sure $GOPATH/bin is in your $PATH.

    hugo

## Building Hugo

    cd $GOPATH/src/github.com/spf13/hugo
    go install

