# Graph Theory Assiignment: Year 3
Christian Olim - G00334621

# Contents
1. NFA.go
2. Rega.go
3. Shunt.go
4. Assignment.go

# Question Given To Us
You must write a program in the Go programming language [2] that can build a non-deterministic finite automaton (NFA) from a regular 
expression, and can use the NFA to check if the regular expression matches any given string of text. You must write the program from
scratch and cannot use the regexp package from the Go standard library nor any other external library. A regular expression is a string
containing a series of characters, some of which may have a special meaning. For example, the three characters “.”, “|”, and “∗ ” have the
special meanings “concatenate”, “or”, and “Kleene star” respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1, and 1∗ means
any number of 1’s. These special characters must be used in your submission. Other special characters you might consider allowing as input
are brackets “()” which can be used for grouping, “+” which means “at least one of”, and “?” which means “zero or one of”. You might also
decide to remove the concatenation character, so that 1.0 becomes 10, with the concatenation implicit. You may initially restrict the non-
special characters your program works with to 0 and 1, if you wish. However, you should at least attempt to expand these to all of the 
digits, and the characters a to z, and A to Z. You are expected to be able to break this project into a number of smaller tasks that are
easier to solve, and to plug these together after they have been completed.


# Resources
1. https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
2. https://en.wikipedia.org/wiki/Thompson%27s_construction
3. https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d
4. https://gobyexample.com/regular-expressions
5. https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
6. https://en.wikipedia.org/wiki/Shunting-yard_algorithm
7. https://brilliant.org/wiki/shunting-yard-algorithm/
8. https://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa


