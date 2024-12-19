package main

import (
	"automata/statemachine"
	"fmt"
	"log"
)

// Use example
func main() {
	// Define states
	stateA := &statemachine.State{Name: "A"}
	stateB := &statemachine.State{Name: "B"}
	stateC := &statemachine.State{Name: "C"}

	//Add more states
	sm, err := statemachine.NewStateMachine(stateA)
	if err != nil {
		log.Print("Failed to create state machine: ", err.Error())
	}

	if err := sm.AddState(stateB); err != nil {
		log.Print("Failed to add state: ", err.Error())
	}

	if err := sm.AddState(stateC); err != nil {
		log.Print("Failed to add state: ", err.Error())
	}

	//Define the transitions
	if err := sm.AddTransition(stateA, stateB); err != nil {
		log.Print("failed to add transition A->B", err.Error())
	}

	if err := sm.AddTransition(stateB, stateC); err != nil {
		log.Print("failed to add transition A->B", err.Error())
	}

	if err := sm.PermitReentry(stateB); err != nil {
		log.Print("error in add reentry", err.Error())
	}

	//Perform state transitions
	if err := sm.Fire(stateB); err != nil {
		log.Print("Failed to transitions to B", err.Error())
	} else {
		log.Print("Transitioned to B")
	}

	if err := sm.Fire(stateB); err != nil {
		log.Print("Failed to transitions to B", err.Error())
	} else {
		log.Print("Transitioned to B")
	}

	if err := sm.Fire(stateC); err != nil {
		log.Print("Failed to transitions to B", err.Error())
	} else {
		log.Print("Transitioned to C")
	}

	fmt.Printf("test")
}
