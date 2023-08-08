# Why
The Fiskil coding challenge

# What
Golang Question
Implement a small HTTP service in Golang that generates a unique sequence number for every request it receives.

Requirements:
*	The service should listen on a configurable port and respond to HTTP GET requests.
*	The service should return a unique sequence number in the response body for each request it receives.
*	The service should ensure that the sequence number generated is unique across all requests.
*	The service should handle multiple requests concurrently and should be thread-safe.

# How

Start the server and run the integration test
```
go install github.com/simonmittag/fiskil/cmd/server && server
go test ./integration
```

Run the unit tests
```
go test -v .
```