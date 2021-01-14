package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// test
// Question and answer data structure
type mathematicsQuiz struct {
	question struct {
		number1 int
		number2 int
	}
	answer int
}

var (
	inputVal              int
	mathChoice            string
	mathChoiceInt         int
	difChoice             string
	difChoiceInt          int
	symbolMath            string
	mathquestionlist      []*mathematicsQuiz
	defaultNumofQuestions     = 20
	mathMenu                  = []string{"Addition", "Subtraction", "Multiplication", "Division"}
	difList                   = []string{"Beginner", "Intermediate", "Advanced", "Expert"}
	input                     = bufio.NewScanner(os.Stdin)
	timeLimit             int = 5
)

func main() {
	// Cli config
	// math = choose category - 1.Addition 2.Subtraction 3.Multiplication 4.Division
	// level = choose level - 1.Beginner 2.Intermediate 3.Advanced 4.Expert
	// math and level is compulsory to choose when using cli command
	// numq and time is alternatively, can be set or not
	// numq = number of questions (default is 20), time = time limit for each question (default is 5s)
	// example: ./golang-cli-math-quiz -math=(1-4) -level=(1-4) -numq=(1-n) -time=(1-n)
	flag.IntVar(&mathChoiceInt, "math", 0, "Mathematics Choice\n1) Addition \n2) Subtraction\n3) Multiplication\n4) Division")
	flag.IntVar(&difChoiceInt, "level", 0, "Level Choice\n1) Beginner\n2) Intermediate\n3) Advanced\n4) Expert")
	flag.IntVar(&defaultNumofQuestions, "numq", defaultNumofQuestions, "How many questions you want to challenge, default is 20")
	flag.IntVar(&timeLimit, "time", timeLimit, "-Time limit per question\n-if you don't want to set a limit, just put 0, default is 5 secs per question")
	flag.Parse()

	// Looping for user can continue do the challenge
	for {
		// if the cli command input not within the criteria will jump to manual selection
		// math and level is compulsory to input together
		if mathChoiceInt > 4 || mathChoiceInt < 1 || difChoiceInt > 4 || difChoiceInt < 1 || defaultNumofQuestions < 1 {
			fmt.Println(strings.Repeat("*", 60))
			fmt.Println("Mathematics Challenge Quiz (Train Your Calculation Speed)")
			fmt.Println(strings.Repeat("*", 60))
			fmt.Println(strings.Repeat("#", 15) + " Please Choose Below Category " + strings.Repeat("#", 15))

			// Manual selection for mathematics and level
			for k, v := range mathMenu {
				fmt.Printf("%d) %v\n", k+1, v)
			}

			// Looping for mathematics choices, the criteria input is 1-4, other than the criteria /
			// will keep looping until get the correct selection
			for {
				fmt.Printf("Select Your Mathematics Choice: ")
				input.Scan()
				if input.Text() == "" {
					fmt.Printf("error: selection can't be blank.\n")
					continue
				}
				inputVal, _ = strconv.Atoi(input.Text())
				if inputVal == 0 || inputVal > 4 {
					fmt.Printf("err: please input a valid number from 1 to 4.\n")
					continue
				}
				mathChoice = mathMenu[inputVal-1]
				break
			}

			// Select the math quiz difficulty level
			fmt.Println(strings.Repeat("#", 15)+" Please Select The Difficulty", strings.Repeat("#", 15))

			for k, v := range difList {
				fmt.Printf("%d) %v\n", k+1, v)
			}

			// Looping for level selection, the criteria input is 1-4,
			// if other than criteria will keep looping until get the correct selection
			for {
				fmt.Printf("Select Your Difficulty Choice: ")
				input.Scan()
				if input.Text() == "" {
					fmt.Printf("error: selection can't be blank.\n")
					continue
				}
				inputVal, _ = strconv.Atoi(input.Text())
				if inputVal == 0 || inputVal > 4 {
					fmt.Printf("err: please input a valid number from 1 to 4.\n")
					continue
				}
				difChoice = difList[inputVal-1]
				break
			}
		} else {
			mathChoice = mathMenu[mathChoiceInt-1]
			difChoice = difList[difChoiceInt-1]
		}

		fmt.Printf(strings.Repeat("-", 10)+" You have selected `%v` with `%v` level "+strings.Repeat("-", 10)+"\n", mathChoice, difChoice)
		fmt.Println()
		fmt.Println()

		// Create the questions based on math and level selection
		var correct int = 0
		switch mathChoice {
		case "Addition":
			symbolMath = "+"
			mathquestionlist = addition(difChoice, defaultNumofQuestions)
		case "Subtraction":
			symbolMath = "-"
			mathquestionlist = subtraction(difChoice, defaultNumofQuestions)
		case "Multiplication":
			symbolMath = "*"
			mathquestionlist = multiplicationdivision(mathChoice, difChoice, defaultNumofQuestions)
		case "Division":
			symbolMath = "/"
			mathquestionlist = multiplicationdivision(mathChoice, difChoice, defaultNumofQuestions)
		}

		// Set the time limit for total question = time for each question * total questions
		timeperQ := timeLimit * defaultNumofQuestions
		timer := time.NewTimer(time.Duration(timeperQ) * time.Second)

		// Looping for print out the questions
	problemloop:
		for k, v := range mathquestionlist {
			answerCH := make(chan string)
			fmt.Printf("Question #%d : %v %v %v = ", k+1, v.question.number1, symbolMath, v.question.number2)

			go func() {
				input.Scan()
				answerCH <- input.Text()
			}()
			select {
			case <-timer.C:
				fmt.Printf("timeout: you have reached the time limit to answer questions.\n")
				break problemloop
			case <-answerCH:
				if input.Text() == "" {
					fmt.Println("you have skip the question, continue to next question.")
					fmt.Println()
					fmt.Println(strings.Repeat("<>", 30))
					fmt.Println()
					continue
				}
				inputVal, _ = strconv.Atoi(input.Text())
				if inputVal == v.answer {
					fmt.Printf("your answer %v is correct\n", inputVal)
					fmt.Println()
					fmt.Println(strings.Repeat("<>", 30))
					fmt.Println()
					correct++
				} else {
					fmt.Printf("your answer is `%v` incorrect, the correct answer is %v\n", inputVal, v.answer)
					fmt.Println()
					fmt.Println(strings.Repeat("<>", 30))
					fmt.Println()
				}
			}
		}

		// Print out results
		fmt.Println(strings.Repeat("+", 50))
		fmt.Printf("Results: you have scored %v out of %v\n", correct, len(mathquestionlist))
		fmt.Printf("You have completed the quiz.\n")
		fmt.Printf(strings.Repeat("-", 50) + "\n")

		// To continue the quiz or quit
		fmt.Printf("Press Enter to Continue or Enter 1 to Exit\n")
		fmt.Printf("or program will exit in 10 seconds\n")
		fmt.Printf("Please Press Enter or 1: ")

		// Set the time limit for decision for continue or quit
		exitTimer := time.NewTimer(time.Duration(10) * time.Second)
		exitCh := make(chan string)
		go func() {
			input.Scan()
			exitCh <- input.Text()
		}()
		select {
		case <-exitTimer.C:
			return
		case exitStr := <-exitCh:
			if exitStr == "" {
				break
			} else if exitStr == "1" {
				return
			}
		}
	}
}

// Addition function, create questions with answer
// Criteria = based on level selected, n + n = 2n (n value will not exceeed the maximum number)
func addition(level string, defaultNumQ int) []*mathematicsQuiz {
	addQuesList := make([]*mathematicsQuiz, defaultNumQ)
	var maxnum int
	switch level {
	case "Beginner":
		maxnum = 10
	case "Intermediate":
		maxnum = 30
	case "Advanced":
		maxnum = 100
	case "Expert":
		maxnum = 1000
	}
	for i := 0; i < defaultNumQ; i++ {
		addQuesList[i] = &mathematicsQuiz{}
		for {
			addQuesList[i].question.number1 = rand.Intn(maxnum)
			addQuesList[i].question.number2 = rand.Intn(maxnum)
			if addQuesList[i].question.number1 == 0 {
				continue
			} else if addQuesList[i].question.number2 == 0 {
				continue
			} else {
				break
			}
		}
		addQuesList[i].answer = addQuesList[i].question.number1 + addQuesList[i].question.number2
	}
	return addQuesList
}

// Subtraction function, create questions with answer
// criteria = based on level selected, x - y = n (y value will not greater than x and x will not exceed to the maximum number)
func subtraction(level string, defaultNumQ int) []*mathematicsQuiz {
	subQuesList := make([]*mathematicsQuiz, defaultNumQ)
	var maxnum int
	switch level {
	case "Beginner":
		maxnum = 10
	case "Intermediate":
		maxnum = 30
	case "Advanced":
		maxnum = 100
	case "Expert":
		maxnum = 1000
	}
	for i := 0; i < defaultNumQ; i++ {
		subQuesList[i] = &mathematicsQuiz{}
		for {
			subQuesList[i].question.number1 = rand.Intn(maxnum)
			subQuesList[i].question.number2 = rand.Intn(maxnum)
			if subQuesList[i].question.number1 <= 1 {
				continue
			} else {
				for {
					subQuesList[i].question.number2 = rand.Intn(maxnum)
					if subQuesList[i].question.number2 >= subQuesList[i].question.number1 || subQuesList[i].question.number2 == 0 {
						continue
					} else {
						break
					}
				}
				break
			}
		}
		subQuesList[i].answer = subQuesList[i].question.number1 - subQuesList[i].question.number2
	}
	return subQuesList
}

// Multiplication and Division function, create questions with answer
// criteria = based on level selected, will create multiplication table to generate questions and answers
// multiplication criteria = x * x = y (x will not exceed the end number)
// division criteria = x / x = y (x will not exceed the end number)
func multiplicationdivision(math, level string, defaultNumQ int) []*mathematicsQuiz {
	var muldivQuesList = make([]*mathematicsQuiz, defaultNumQ)
	var startnum, endnum int
	switch level {
	case "Beginner":
		startnum = 2
		endnum = 12
	case "Intermediate":
		startnum = 2
		endnum = 24
	case "Advanced":
		startnum = 10
		endnum = 24
	case "Expert":
		startnum = 20
		endnum = 30
	}
	multable := multiplicationTable(startnum, endnum)
	var randList []int
	for {
		if len(randList) != defaultNumQ {
			x := rand.Intn(len(multable))
			if len(randList) > 0 {
				for _, v := range randList {
					if v == x {
						break
					} else {
						randList = append(randList, x)
						break
					}
				}
			} else {
				randList = append(randList, x)
			}
		} else {
			break
		}
	}
	switch math {
	case "Multiplication":
		for k, v := range randList {
			muldivQuesList[k] = &mathematicsQuiz{}
			muldivQuesList[k].question.number1 = multable[v].question.number1
			muldivQuesList[k].question.number2 = multable[v].question.number2
			muldivQuesList[k].answer = multable[v].answer
		}

	case "Division":
		for k, v := range randList {
			muldivQuesList[k] = &mathematicsQuiz{}
			muldivQuesList[k].question.number1 = multable[v].answer
			muldivQuesList[k].question.number2 = multable[v].question.number2
			muldivQuesList[k].answer = multable[v].question.number1
		}
	}
	return muldivQuesList
}

// Create multiplication table
// Purpose is to easy to create multiplication and division questions and answers with less duplication Q&A
func multiplicationTable(start, end int) []*mathematicsQuiz {
	var multablelist []*mathematicsQuiz
	for i := start; i <= end; i++ {
		for x := start; x <= end; x++ {
			multablelist = append(multablelist, &mathematicsQuiz{
				question: struct {
					number1 int
					number2 int
				}{number1: i, number2: x},
				answer: i * x,
			})
		}

	}
	return multablelist
}
