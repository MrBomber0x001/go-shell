## Intro

Native Shell processor written in Go, the application is currently in progress!

[![asciicast](https://asciinema.org/a/V4tqMsnHN7qznKuN6iRUzve3a.svg)](https://asciinema.org/a/V4tqMsnHN7qznKuN6iRUzve3a)

## Functionalites

All of the following commands are implemented natively *without* using `cmd.Execute()`

commands: [cat ls pwd cd grep curl find]

## Logs

- [ ] Write user errors to logs! [logging]
- [ ] Setup Github actions & Dockerize the application
- [ ] Write unit tests and table-driven tests
- [ ] integrate proper error and validation handling
- [ ] [FEATURE] implement auto-suggestions based on frequent commands hit by the user
- [ ] Handling Signals
