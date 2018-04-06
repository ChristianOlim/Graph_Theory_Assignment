// G00334621 - Christian Olim - Year 3 (Group C)
// Resources: https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
// Resources: https://en.wikipedia.org/wiki/Thompson%27s_construction

// Thompson's Construction
// In computer science, Thompson's construction is an algorithm used to change a regular
// expression into an equivalent nondeterministic finite automaton (NFA). This NFA can then
// be used to find strings against the regular expression.

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
	return nfastack[0]
}

func main(){
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
}