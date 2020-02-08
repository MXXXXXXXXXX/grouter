# grouter
A simple HTTP router for Go. Use it to build RESTful APIs or your web framework.

## Motivation

I am a beginner in golang. I have developed on nginx for some years, I wanted a simple and high scalability HTTP GO router, which supports regexp. I prefer to support regexp is because otherwise it will need the logic to check the URL parameter type, thus increasing the program complexity.

# Features

* High scalability
* Support regex
* Support Routes groups
* Custom NotFoundHandler
* Custom PanicHandler
* HTTP Method Get、Post、Delete、Put、Patch Support
* Support multiple Middleware Chain
* No external dependencies (just Go stdlib)

## High scalability

  while do regex, we can create grouter like `/{userid:[0-9a-zA-Z_.]+}`, while append Middleware Chain, we can get userid instance from req.context() then wo can do what we want to do with userid in ServeHttp()

## Requirements

* golang 1.8+

## Installation

```
go get -u github.com/MXXXXXXXXXX/grouter
```