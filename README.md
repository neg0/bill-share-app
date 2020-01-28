# Bill Share
* Golang 1.3

## Internal and pkg
__internal__ packages are designed to be used only within this application (Microservice). However, __pkg__ packages are 
designed in a way that could be used by other microservices. For example if another microservice wishes to utilise 
Bill package or Money they could just import pkg into a new microservice to avoid code duplication all across our 
distributed systems.


## Domain Driven Design
Application designed according to DDD principles, please see diagram and brief description below.
 * __Application layer__: problem
 * __Domain layer__: participant, payback, receipt, and money
 * __Infrastructure layer__: bill, and cmd
[]()


## TDD & BDD
Tests are provided inside each package under directory named `test`. Ginkgo library used for creating BBD styled 
test suites and could be invoked from the root of the project by: `go test ./... -v`. Due to nature of mocking 
objects, I also created a mock Struct types inside few tests to avoid duplication and clear test files for better 
readability.


## Environment
Please ensure you create a `.env` file from `.env.dist`. This being used to define the JSON filename as the entries.

    ~$: cp .env.dist .env
    

## Makefile and Docker
Application is containerised using Docker and Makefile provided for convenience and shorten the docker 
commands to build and run the application. 

    ~$: make up
    ~$: make run

# Business Requirements & Scope
First Service is called Overview which gives an array of strings that contains an overview report of who owes whom.
Second Service is PayBack which gives an instruction of how to pay back each other

program included inside `cmd` named `problem_statement`. In order to run the program please ensure `testfile.json` is 
included at the root of the project. In order to run the program please run the following:

    ~$: go run cmd/problem_statement.go
    Tommen is owed £25.50
    Tommen owes Kelly £4.50
    Ola owes Sam nothing
    Sandy pays Kelly £9.50
    Sandy pays Tommen £5.50

Comments included where was possible but I ensured methods them-self are self-explanatory, hope this helps. Thank you


## Demo
program included inside `cmd` named `problem_statement`. In order to run the program please ensure `testfile.json` is 
included at the root of the project. Please run the following:

    ~$: go run cmd/problem_statement.go
    Tommen is owed £25.50
    Tommen owes Kelly £4.50
    Ola owes Sam nothing
    Sandy pays Kelly £9.50
    Sandy pays Tommen £5.50