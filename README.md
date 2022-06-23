# Backend Developer Tests

Hey there! Do you want to build the next generation of edge services? Do you 
want to work on cutting-edge technology with a great team full of talented and 
motivated individuals? StackPath is [hiring backend developers](https://stackpath.applytojob.com/apply/)! 
StackPath's backend powers not just our [customer portal](https://control.stackpath.com) 
and [API](https://developer.stackpath.com/) but is the core behind many of our 
amazing products. 

We love [golang](https://golang.org/) at StackPath. Most of our services are 
written in go, and we are always looking to add bright and awesome people to our 
ranks. If you think this sounds great and if you have what it takes then 
apply for one of our backend service positions. If being a backend engineer isn't your 
thing then we have [many open positions](https://stackpath.applytojob.com/) to go for.

We employ a lot of modern technology and processes at StackPath. These three 
exercises are intended to demonstrate your basic knowledge of go and its 
applications. These problems have many solutions. Our managers are most interested in 
how you choose to solve them and why. There are no wrong answers, as long as 
these examples compile and work in at least go 1.14 and use [go modules](https://golang.org/ref/mod)
for package management.

## Unit Testing

We're pretty serious about testing at StackPath. In addition to QA, all of our 
code includes unit tests that are run in builds. Builds must pass before code 
can go live. 

The `unit-testing` folder contains a [FizzBuzz](https://imranontech.com/2007/01/24/using-fizzbuzz-to-find-developers-who-grok-coding/) 
example implemented in go. It takes three optional command line arguments:
- The number of numbers to iterate over
- The multiple to "Fizz" on
- The multiple to "Buzz" on

For instance:

```
$ go run main.go 16 3 5
SP// Backend Developer Test - FizzBuzz

FizzBuzzing 16 number(s), fizzing at 3 and buzzing at 5:
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
16

Done
```

It works well enough, but doesn't have any unit tests. Without tests, we can't 
prove to everyone that the code works and that it doesn't affect other services. 

Look in the `unit-testing/pkg/fizzbuzz/fizzbuzz_test.go` file and implement unit 
tests that flex the `FizzBuzz()` function in `unit-testing/pkg/fizzbuzz.go`. Think 
of the different kinds of inputs that can be passed to the function and how it 
should act when common and edge cases are used against it. Did your tests find 
any bugs in `FizzBuzz()`? If so then fix 'em up! Can you get greater than 80% 
code coverage with your tests? 

Run `go test -v -cover ./...` from the `unit-testing` directory and let us know 
how you did.

### Candidate's Notes
I was able to get to 100% coverage of `FizzBuzz()` and found a couple of potential bugs. Normally I would address some
of these bugs by using Go's builtin error type, but I didn't want to change the function signature of `FizzBuzz()` so I 
changed the behavior of `FizzBuzz()` as I saw best fit to address any identified bugs.

The bugs I did find and addressed are the following:

- `make()` panic: If the `total` parameter is negative, it will cause `make()` to panic. I added a guard that returns empty results for negative totals.
- divide by zero panic: If either `fizzAt` or `buzzAt`

## Web Services

Web services are our bread and butter. Our services talk to each other over 
[gRPC](https://grpc.io/) and [REST](https://en.wikipedia.org/wiki/Representational_state_transfer). 
The `rest-service` directory contains a simple `Person` model and a set of 
sample data that needs a REST service in front of it. This service should:

- Respond with JSON output
- Respond to `GET /people` with a 200 OK response containing all people in the 
  system
- Respond to `GET /people/:id` with a 200 OK response containing the requested 
  person or a 404 Not Found response if the `:id` doesn't exist
- Respond to `GET /people?first_name=:first_name&last_name=:last_name` with a 
  200 OK response containing the people with that first and last name or an 
  empty array if no people were found
- Respond to `GET /people?phone_number=:phone_number` with a 200 OK response 
  containing the people with that phone number or an empty array if no people 
  were found

You can implement the service with go's built-in routines or import a framework 
or router if you like. The `Person` model and all of the backend code is in the 
`rest-service/pkg/models/person.go` file, the service should be initialized in 
`rest-service/main.go`, and should run by running `go run main.go` from the 
`rest-service` directory.

Implementing the service is a good start, but are there any extras you can throw 
in? How would you test this service? How would you audit it? How would an ops 
person audit it?

### Candidate's Notes
The implementation built meets the requirements / acceptance criteria but is rather unpolished
in its current state. If more time was available the following would be done to clean it up
and make it more resilient. This todo list includes the following:

- **Make application more dry**: The current structure and implementation of the handlers has a lot of repeated code that should be broken
out into separate callables. Most of this has to do with handling query parameters, writing headers, and writing response bodies.
- **Improve error handling**: The current error handling could be made more sophisticated and robust in order to better handle errors
raised by the `net/http` package and `json` packages. Some of these errors might recoverable and not necessitate treating
them as fatal.
- **Standardize logging behavior**: A consistent pattern of when to log information and what to log should be adopted. This can include
patterns such as weather to log on requests resulting in errors or all requests, whether to include request parameters in
the log entries, and whether any networking information in the log.
- **Add Unit Tests**: In order to confirm correctness, unit tests should be incorporated to validate the behavior of the `person` model
and the HTTP handler functions. The handlers can be tested by passing in `http.Request` types and inspecting the response that is
returned.
- **Make runtime options configurable**: To make the service more flexible, runtime parameters such as network protocol and
TCP/UDP port should be selectable by passing in these configuration values to the executable at startup. These were hardcoded
to TCP and 8000 for the sake of expediency.

For testing this service, I would rely on unit tests to check the correctness of the handlers and models. For integration testing
I would use Go's standard library package `httptest`. For end-to-end testing, I would leverage some HTTP testing utility that can
be scripted based on expected behavior.

For auditing, I would leverage test coverage utilities and RESTful schema documentation tools such as Swagger.


## Input Processing

StackPath operates a gigantic worldwide network to power our edge services. These 
nodes are in constant communication with each other and various central systems. 
Our services have to be robust enough to handle this communication at scale.

The third programming test in the `input-processing` directory contains a 
program that reads STDIN and should output every line that contains the word "error" 
to STDOUT. We've taken care of most of the boilerplate, but the rest is up to you. 

Consider scale when implementing this. How well will this work if a line is 
hundreds of megabytes long? What if 10 gigabytes of information is passed to it? 
What if entries are streamed to it? How would you differentiate between errors read 
from the stream vs program errors? How would you test this? Assume that `\n` ends a 
line of input. Was with the REST service test you're free to use any built-ins or 
import any frameworks you like to do this.

## Concurrency

In some advanced situations, StackPath relies on Go's concurrency primitives to
perform asynchronous tasks. The `concurrency` directory contains two interfaces for
asynchronous task pools, `SimplePool` and `AdvancedPool`, whose implementation
requirements are documented on the interface. Please write an implementation for
`NewSimplePool`. As an extra bonus challenge, but by no means required or necessary,
feel free to also write an implementation for `NewAdvancedPool`.

While third party libraries can assist, many times our unique needs rely on channels
directly, so please refrain from using anything but the Go standard library for this
exercise. Unit tests are not necessary, but feel free to write them if it helps
during the development.

## Contributing

What did you think of these? Are these too easy? Too hard? We're open to change 
and accept [issues](https://github.com/stackpath/backend-developer-tests/issues) 
and [pull requests](https://github.com/stackpath/backend-developer-tests/pulls) 
for these tests. See our [contributing guide](https://github.com/stackpath/backend-developer-tests/blob/master/.github/contributing.md) 
for more information. Thanks for giving these a try, and we hope to hear from you 
soon!
