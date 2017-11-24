// G00334621
// Data Representation Project 2017
// Christian Olim
// https://github.com/data-representation/eliza/blob/master/eliza.go
// https://github.com/data-representation/eliza/blob/master/data/substitutions.txt
// https://golang.org/pkg/regexp/
// https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/
// https://golang.org/pkg/math/rand/
// https://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go

package eliza

import (
	"time"	
	"bufio"
	"math/rand"
	"strings"
	"regexp"
	"fmt"
	"os"
)

// This will create a structure for our responses
type Response struct {
	Patterns *regexp.Regexp
	Answers  []string
}

// This reads our file, splits it and then makes patterns
func makeResponses(path string) []Response {
	fullFile, _ := scanLines(path)
	responses := make([]Response, 0)
	for i := 0; i < len(fullFile); i += 2 {
		allPatterns := strings.Split(fullFile[i], ";")
		allResponses := strings.Split(fullFile[i+1], ";")
		for _, pattern := range allPatterns {
			pattern = "(?i)" + pattern
			Patterns := regexp.MustCompile(pattern)
			responses = append(responses, Response{Patterns: Patterns, Answers: allResponses})
		}
	}
	return responses
}

func PrintResponses(path string) {
	response := makeResponses(path)
	fmt.Printf("%+v\n", response)
}

// This function will read strings line by line
func scanLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		readLine := scanner.Text()
		// This skips past any commentary added to our files
		if hideComments(readLine) {
			continue
		}
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Here is our function for skipping commentary
func hideComments(readLine string) bool {
	return strings.HasPrefix(readLine, "//") || len(strings.TrimSpace(readLine)) == 0
}


func mapSubstitutions(inputString string) string {
	// This splits up strings in sections
	splitStr := strings.Fields(inputString)

	// Substitutions for Eilza to add human-like responses
	substitutions := map[string]string{
		"i":      "you",
		"you'll": "I will",
		"are":    "am",
		"you've": "I have",
		"was":    "were",
		"my":     "your",
		"me":     "you",
		"you're": "I’m",
		"your":   "my",
		"you":    "I",		
	}

	// This will scan for simularites and the replace characters
	for index, word := range splitStr {
		if value, ok := substitutions[strings.ToLower(word)]; ok {
			splitStr[index] = value
		}
	}
	return strings.Join(splitStr, " ")
}

// This function is used for substituting strings
func stingReplace(pattern *regexp.Regexp, input string) string {
	match := pattern.FindStringSubmatch(input)
	if len(match) == 1 {
	}
	wordSwap := match[1]
	wordSwap = mapSubstitutions(wordSwap)
	return wordSwap
}

// Function for a response that has a  "%s" formatter
func responseBuilder(response, wordSwap string) string {
	if strings.Contains(response, "%s") {
		return fmt.Sprintf(response, wordSwap)
	}
	return response
}

// AskEliza will begin creating a response for the user's input
func AskEliza(input string) string {
	// Creating a new instance of Eliza
	response := makeResponses("./dat/elizaResponses.dat")
	randomResponse, _ := scanLines("./dat/generalResponses.dat")
	rand.Seed(time.Now().Unix())

	for _, response := range response {
		if response.Patterns.MatchString(input) {
			wordSwap := stingReplace(response.Patterns, input)
			genResp := response.Answers[rand.Intn(len(response.Answers))]
			genResp = responseBuilder(genResp, wordSwap)
			return genResp
		}
	}
	// In the case that we can't respond to the user's input, we will generate a random response
	return randomResponse[rand.Intn(len(randomResponse))]
}
