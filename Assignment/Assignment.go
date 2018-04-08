// G00334621 - Christian Olim - Year 3 (Group C)
// Resources: https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
// Resources: https://en.wikipedia.org/wiki/Thompson%27s_construction
// Resources: https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d
// Resources: https://gobyexample.com/regular-expressions
// Resources: https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
// Resources: https://en.wikipedia.org/wiki/Shunting-yard_algorithm
// Resources: https://brilliant.org/wiki/shunting-yard-algorithm/
// Resources: https://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa

package main

import (
	"fmt"
	"os"
    "bufio"
    "strings"
)

//-------------------- Below is the intopost function for the Shunt.go file -----------------------------------------------
//-------------------------------------------------------------------------------------------------------------------------

// Return type infix is returned from this function.
func intopost(infix string) string {
    specials := map[rune]int{'*': 10, '.': 9, '|': 8}

    // pofix = postfix, s = stack.
    pofix, s := []rune{}, []rune{} 

    // This for loop will loop over the infix and will then give us an index of read characters.
    for _, r:= range infix {
        
        switch {

        case r == '(': 
            s = append(s, r) 

        case r == ')': 

        for s[len(s)-1] != '(' {
            pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
        }

        s = s[:len(s)-1]

        case specials[r] > 0: 
            for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
                pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
            }
            s = append(s, r)

        default:
          pofix = append(pofix, r)
        }
    }

    for len(s) > 0 {
        pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
    }

    // Postfix is returned in a string format.
    return string(pofix)
}//-------------------------------------------------------------------------------------------------------------------------



//-------------------- Below is a match funtion for a string and a post fixed regular --------------------------------------
//-------------------- expression. We will then change an infix regexp to postfix regexp. ----------------------------------
//-------------------- This function can be found in the NFA.go file -------------------------------------------------------
//--------------------------------------------------------------------------------------------------------------------------

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
//-------------------------------------------------------------------------------------------------------------------------


// How to read an input from the command line.
func userInput() (string, error){
    reader := bufio.NewReader(os.Stdin)
    str, err := reader.ReadString('\n')
    return strings.TrimSpace(str),err
}


//-------------------- Main function with User Interface ------------------------------------------------------------------
//-------------------------------------------------------------------------------------------------------------------------

func main() {
	// User Interface.
	fmt.Println("===============================================================")
	fmt.Println("========================= G00334621 ===========================")
	fmt.Println("================== Graph Theory Project 2018 ==================")
	fmt.Println("=============================================================== \n")
	fmt.Println("Please Choose From One Of The Options Below: ")
	fmt.Println("1.) Convert a Regular Expression from Infix notation to Postfix notation string.")
	fmt.Println("2.) Convert a Regular Expression to a Nondeterministic Finite Automaton (NFA).")
	fmt.Println("3.) Exit.")

	// User's choice.
	var option int

	// Read User's Input.
	fmt.Scanln(&option)

	// Switch Statement.
	switch option {
	case 1:
		fmt.Println("You entered option ", option, ". ")	
		fmt.Print("Please enter a Regular/ Infix expression: ")
	
		reader := bufio.NewReader(os.Stdin)
		regex, _ := reader.ReadString('\n')

		fmt.Println("Postfix = ", Intopost(regex))
	}

	case 2:
		fmt.Println("You entered option ", option, ". ")	
		fmt.Print("Please enter a Regular/ Infix expression: ")
		reader := bufio.NewReader(os.Stdin)
		regex, _ := reader.ReadString('\n')
		fmt.Println("NFA: ", Poregtonfa(regex))

		fmt.Println("Please enter a string to see if it matches the NFA: ")
		stringInput, _ := reader.ReadString('\n')
		stringInput = Intopost(stringInput)

		if pomatch(regex, stringInput) == false {
			fmt.Println("The string you entered doesn't match correctly.")
		} else if pomatch(regex, stringInput) == true {
			fmt.Println("The string you entered matches correctly with the regular expression.")
		} else {
			fmt.Println("Sorry there was a problem.")
		}

	default:
		fmt.Println("Goodbye.")
}