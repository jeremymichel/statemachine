package statemachine

import (
	"errors"
	"fmt"
)

type StateMachine struct {
	CurrentState State
	states       map[string]State
}

// NewStateMachine defines a new state machine and initializes the states map
func NewStateMachine() *StateMachine{
	sm := StateMachine{
		states: make(map[string]State),
	}

	return &sm
}

// Add Adds a new state to the map and keys it by the given name
func (sm *StateMachine) Add (name string, state State) {
	// Add a new state to the map
	sm.states[name] = state
}

// Update Runs the Update func of the current state
func (sm *StateMachine) Update () {
	sm.callUpdate()
}

// Activate Calls the Leave function of the current state, switches to the new state
// based on the given name, and calls the Enter func of the newly activated state
// Will return an error in case the state could not be found in the map
func (sm *StateMachine) Activate (name string) error {
	// Leave the current state
	sm.callLeave()

	// Get the state from the map
	state, ok := sm.states[name]
	if !ok {
		return errors.New(fmt.Sprintf("could not find a state matching the name : %s", name))
	}

	// Switch the current state to the new one
	sm.CurrentState = state

	// Call the enter func on the new state
	sm.callEnter()

	return nil
}

func (sm *StateMachine) callEnter() {
	sm.CurrentState.Enter()
}

func (sm *StateMachine) callUpdate() {
	sm.CurrentState.Update()
}

func (sm *StateMachine) callLeave() {
	sm.CurrentState.Leave()
}