// G00334621 - Christian Olim - Year 3 (Group C)
// Resources: https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d
// Resources: https://gobyexample.com/regular-expressions

// Use of a match funtion on a string and a post fixed regular
// expression. We will then change an infix regexp to postfix regexp.

package main

import (
    "fmt"
)

// This struct is used to store the states and arrows.
// The maximum number of arrow for each state is 2.
type state struct{
	symbol rune
	edge1 *state
	edge2 *state
}

// This keeps track of the initial state and the accept state
// of your fragment from the nondeterministic finite automaton.
type nfa struct{
	initial *state
	accept *state
}

// Poregtonfa = post fixed regular expression to NFA. 
func poregtonfa(pofix string) *nfa{
	nfastack := []*nfa{}

	// This will loop through the postfix regular expression. 
    for _, r := range pofix {

		switch r {
		// For this case in terms of concatenate, we pop two elements off the nfa stack.	
		case '.':
			frag2 := nfastack[len(nfastack)-1]
			// This will get rid of last thing off the nfastack
			nfastack = nfastack[:len(nfastack)-1]
            frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			
			// Adds frag1 to frag2.
			frag1.accept.edge1 = frag2.initial

			// Appends frag1 and frag2 to the of the stack.
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})

		// For this case in terms of or, we again pop two fragments off stack similar to previous case.
		case '|':
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		// For this case in terms of Kleene star, we pop one fragment off the nfa stack.
		case '*':
			frag := nfastack[len(nfastack)-1]
            nfastack = nfastack[:len(nfastack)-1]

            accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			
            frag.accept.edge1 = frag.initial
            frag.accept.edge2 = &accept
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
			
		// This will push an element to the stack.
		default:
			accept := state{}
            initial := state{symbol: r, edge1: &accept}

            nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}
	if len(nfastack) != 1 {
		fmt.Println("huh:", len(nfastack), nfastack)
	}
	return nfastack[0]
}

func addState(l []*state, s *state, a *state) []*state {
   
	l = append(l, s)
    // Any state has a value of 0 rune, then this state will have e arrows coming from it.
    if s != a && s.symbol == 0 {
        l = addState(l, s.edge1, a)
        if s.edge2 != nil {
            l = addState(l, s.edge2, a)
        }
    }
    return l
}

// This function will take a post fixed regular expression
// and a string which will return a boolean value.
func pomatch(po string, s string) bool {
	
	ismatch := false
	ponfa := poregtonfa(po)
	
    // Here we will create an array of pointers.
    current := []*state{}
    next := []*state{}

    // Now we will create a function to add to the current state.
    current = addState(current[:], ponfa.initial, ponfa.accept)

	// For loop to loop through s and current. This will then put all of
	// the current states into next, from current. Finally, it will clear next.
    for _, r := range s {
        for _, c := range current {
            if c.symbol == r {
                next = addState(next[:], c.edge1, ponfa.accept)
            }
        }
        current, next = next, []*state{}
    }

    // For loop through the current array. 
    for _, c := range current {
        if c == ponfa.accept {
            ismatch = true
            break
        }
    }
    return ismatch
}

func main(){
	fmt.Println(pomatch("ab.c*|", "mmm"))
}