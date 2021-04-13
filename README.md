# StateMachine

[![Go Reference](https://pkg.go.dev/badge/github.com/jeremymichel/statemachine.svg)](https://pkg.go.dev/github.com/jeremymichel/statemachine)

A dead simple implementation of a finite state machine in Go.

## Usage

Go get the module 
```shell
go get -u github.com/jeremmymichel/go-statemachine
```
The module works by creating a new `StateMachine`, defining states and adding them to the state machine.

### New State Machine

```go
// Create a new State Machine
stateMachine := statemachine.NewStateMachine() // returns a *statemachine.StateMachine type
```

### Define a state

The `StateMachine` only accepts types implementing the `State` interface, if you choose not to create a compliant struct, you may use the `AnonymousState` struct provided.

The `State` interface requires an implementation of 3 functions,
1. `Enter()` Will be called when the state machine switches to this state (gives you the opportunity to do some setup)
2. `Update()` Will be called on every update of the state machine, when this state is activated
3. `Leave()` Will be called once the state machine is about to switch to a new state (a great place to do some cleanup)

As mentioned, an `AnonymousState` struct is available if you do not need a dedicated struct for a state.

```go
// Create a new AnonymousState

anonymousState := statemachine.NewAnonymousState(func () {}, func () {}, func () {})

// The 3 funcs passed above are the enter, update and leave func respectively, if you do not need one or any of them, just pass in nil
```

### Add a state to the machine

When adding a new `State` to the `StateMachine`, you will need to pass in a name, this is the key to your state in the internal map.

```go
// Add a state struct compliant with the State interface
stateMachine.Add("running", *SomeRunningState)
```

### Activate a state

Activating a state requires the name of said state. `Activate` returns an `error` when the passed in name does not correspond to any state of the machine

```go
// Activate a state
err := stateMachine.Activate("running")
```
