package statemachine

type State interface {
	Enter() // Called when the state machine switches to this state
	Update() // Called on every update of the state machine when this state is active
	Leave() // Called upon switching to another state
}

