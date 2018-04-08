// G00334621 - Christian Olim - Year 3 (Group C)
// Resources: https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
// Resources: https://en.wikipedia.org/wiki/Shunting-yard_algorithm
// Resources: https://brilliant.org/wiki/shunting-yard-algorithm/

// Shunting-Yard Algorithm
// In computer science, Shunting-yard algorithm is a method for parsing mathematical
// expressions specified in infix notation. It can produce either a postfix notation string.

package main

import (
    "fmt"
)

func intopost(infix string) string {
    specials := map[rune]int{'*': 10, '.': 9, '|': 8}

    //pofix = postfix, s = stack.
    pofix, s := []rune{}, []rune{} 

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

    return string(pofix)
}

func main() {
    // The answer should be ab.c
    fmt.Println("infix: ", "a.b.c*")
    fmt.Println("postfix: ", intopost("a.b.c*"))

    // The answer should be abd|.*
    fmt.Println("infix: ", "(a.(b|d))*")
    fmt.Println("postfix: ", intopost("(a.(b|d))*"))

    // The answer should be abd|.c*.
    fmt.Println("infix ", "a.(b|d).c*")
    fmt.Println("postfix ", intopost("a.(b|d).c*"))
 
    // The answer should be abb.+.c.
    fmt.Println("infix ", "a.(b.b)+.c")
    fmt.Println("postfix ", intopost("a.(b.b)+.c"))
}