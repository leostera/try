# try [![Travis-CI](https://api.travis-ci.org/ostera/try.svg)](https://travis-ci.org/ostera/try)
:recycle: A portable Go utility to retry commands with backoff

## Installation

```
go get github.com/ostera/try

brew install https://raw.githubusercontent.com/ostera/homebrew-core/master/Formula/go-try.rb
```

From source just run `make` and put the `try` executable somewhere handy.

## Usage

```
~ λ try

   Usage: try [options] <cmd>

   Sample: try -i=100ms docker pull ubuntu:latest

   Options:

     -h, --help                 this help page
     -v, --version              print out version

```

## Motivation

The main pain point was re-running commands that are affected by transient errors.

`curl`-ing an HTTP endpoint, `docker pull`-ing an image, any remote ssh transfers.
All of these are typically vulnerable to transient network errors.

It's fairly straightforward to do a while loop in bash to repeat the command if it
fails, but you don't want it to run infintely, and you also want to see the output
as it executes, and it should capture command failures gracefully and not explode.

As complexity arises, `bash` syntax gets more and more cryptic.

My interim solution for it was this:

```bash
#!/usr/bin/env bash
set -x
set -e

COMMAND=$*

exec 3>&1
exec 4>&2
DONE=$(
  trap 'exit 0;' ERR;
  COUNTER=0
  RESULT=$(
    while [[ $COUNTER -lt 10 ]] \
    && ! eval $COMMAND 1>&3 2>&4 \
    &&  sleep $(($COUNTER*10)); do
      echo "Retry $COUNTER..."
      let COUNTER++
    done
  )
  echo YES
)
```

Hacky, but does the job. Now it's time I have a version I can run anywhere without
thinking of pipe alias syntax, or sub-shell trapping.

## Contributing

If you don't have it yet, you can install `watch` to continuously build the project:

```
repos/try λ ./watch make
exit: 0
```
