package worker

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
)

// Result represents the result of a job execution, containing the URL, an error if any occurred, and the hash of the response body if successful.
type Result struct {
	Url  Job
	Err  error
	Hash string
}

// String is a method attached to Result that returns a string representation of the result.
func (r *Result) String() string {
	if r.Err != nil {
		return fmt.Sprintf("%s %s", r.Url, r.Err.Error())
	}

	return fmt.Sprintf("%s %s", r.Url, r.Hash)
}

// Job represents a unit of work to be executed, in this case, an HTTP URL to be fetched and hashed.
type Job string

// execute is a method attached to Job that fetches the URL, calculates the hash of the response body, and returns a Result.
func (j Job) execute() Result {
	res, err := http.Get(string(j))
	if err != nil {
		return Result{Err: err, Url: j}
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Result{Err: err, Url: j}
	}
	defer res.Body.Close()

	sum := fmt.Sprintf("%x", md5.Sum(body))

	return Result{Hash: sum, Url: j}
}
