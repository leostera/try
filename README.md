# try [![Travis-CI](https://api.travis-ci.org/ostera/try.svg)](https://travis-ci.org/ostera/try)
:recycle: A portable Go utility to retry commands with backoff

## Installation

```
go get github.com/ostera/try
```

From source just run `make` and put the `try` executable somewhere handy.

## Usage

```
~ λ try

   Usage: try [options] <cmd>

   Sample: try -i=10s -r=10 docker pull ubuntu:latest

           Run the command up to 10 times, with the start interval of 10 seconds,
           doubling the interval on every iteration.

   Options:

     -i, --interval             start interval time (default to 1s)
     -r, --retries              amount of retries (default to 10)
     -f, --factor               multiply interval by this factor (default to 2)

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
#!/bin/bash -xe

COMMAND=$*

exec 3>&1
exec 4>&2
DONE=$(
  trap 'exit 0;' ERR;
  COUNTER=0
  RESULT=$(
    while [[ $COUNTER -lt 10 ]] \
    && ! eval $COMMAND 1>&3 2>&4 \
    && sleep $(($COUNTER*10)); do
      echo "Retry $COUNTER..."
      let COUNTER++
    done
  )
  echo YES
)

if [ -z "$DONE" ] || [ "$DONE" == \'\' ]; then
  exit 1
fi
```

Hacky, but does the job. Now it's time I have a version I can run anywhere without
thinking of pipe alias syntax, or sub-shell trapping.

## Contributing

If you don't have it yet, you can install `watch` to continuously build the project:

```
repos/try λ watch make
/Users/leostera/.go/bin/gometalinter @.gometalinter
/usr/local/bin/go vet
/usr/local/bin/go build -o ./watch
/usr/local/bin/go test
PASS
ok      _/Users/leostera/repos/watch    0.013s
/usr/local/bin/go test -bench .
PASS
BenchmarkRunSuccessfully-4           500           3860598 ns/op
BenchmarkRunExit-4                   300           3834855 ns/op
ok      _/Users/leostera/repos/watch    3.908s
exit: 0
```
