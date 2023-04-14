# Go worker example

This is an example of tool which makes http requests and prints the address of the
request along with the MD5 hash of the response.

## Usage

```bash
$ make build
$ /.myhttp [-parallel N] URL1 [URL2...]
```

Some examples: 

```
$ ./myhttp http://www.adjust.com http://google.com
$ ./myhttp -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com
```

or 
```bash
$ go run main.go http://www.adjust.com http://google.com
```

### Parameters:

The tool is able to limit the number of parallel requests, to prevent
exhausting local resources. The tool accept a flag to indicate this limit, and
by default it is equal to 10 if the flag is not provided.

1. `-parallel 3` The number of workers


##Run tests:

To run unit tests:

```bash
$ make test-unit
```

or

```bash
$ go test -cover -race ./... -coverprofile=/tmp/cover.out && go tool cover -html=/tmp/cover.out
```
