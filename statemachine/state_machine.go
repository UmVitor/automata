package statemachine

import (
	"errors"
	"fmt"
	"os"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

type State struct {
	Name string
}

type StateMachine struct {
	graph   graph.Graph[string, *State]
	current *State
}

func NewStateMachine(initialState *State) (*StateMachine, error) {
	g := graph.New(func(state *State) string {
		return state.Name
	}, graph.Directed())

	if err := g.AddVertex(initialState); err != nil {
		return nil, fmt.Errorf("failed to add initial state: %w", err)
	}

	return &StateMachine{graph: g, current: initialState}, nil
}

func (sm *StateMachine) AddState(state *State) error {
	return sm.graph.AddVertex(state)
}

func (sm *StateMachine) AddTransition(from, to *State) error {
	if _, err := sm.graph.Vertex(from.Name); err != nil {
		return fmt.Errorf("state %s does not exists: %s", from.Name, err.Error())
	}

	if _, err := sm.graph.Vertex(to.Name); err != nil {
		return fmt.Errorf("state %s does not exists: %s", to.Name, err.Error())
	}

	return sm.graph.AddEdge(from.Name, to.Name)
}

func (sm *StateMachine) CanTransition(to *State) bool {
	_, err := sm.graph.Edge(sm.current.Name, to.Name)
	return err == nil
}

func (sm *StateMachine) Fire(to *State) error {
	if !sm.CanTransition(to) {
		return errors.New("trasition not allowed")
	}
	sm.current = to
	return nil
}

func (sm *StateMachine) GetCurrentState() *State {
	return sm.current
}

func (sm *StateMachine) PermitReentry(state *State) error {
	return sm.graph.AddEdge(state.Name, state.Name)
}

// DrawStateMachine draws the state machine graph in DOT language
func (sm *StateMachine) DrawStateMachine() {
	file, _ := os.Create("./my-state-machine.gv")
	_ = draw.DOT(sm.graph, file)
}
