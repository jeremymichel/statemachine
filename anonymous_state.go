package statemachine

type AnonymousState struct {
	EnterFunc func()
	UpdateFunc func()
	LeaveFunc func()
}

// NewAnonymousState Creates a new anonymous state and passes in the given funcs
func NewAnonymousState(enter, update, leave func()) *AnonymousState {
	as := AnonymousState{EnterFunc: enter, UpdateFunc: update, LeaveFunc: leave}
	return &as
}

func (a *AnonymousState) Enter()  {
	if a.EnterFunc != nil {
		a.EnterFunc()
	}
}

func (a *AnonymousState) Update()  {
	if a.UpdateFunc != nil {
		a.UpdateFunc()
	}
}

func (a *AnonymousState) Leave()  {
	if a.LeaveFunc != nil {
		a.LeaveFunc()
	}
}
