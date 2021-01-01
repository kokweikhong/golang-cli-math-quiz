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

// "Beginner", "Intermediate", "Advanced", "Expert"
func main() {
	flag.IntVar(&mathChoiceInt, "math", 0, "Mathematics Choice\n1) Addition \n2) Subtraction\n3) Multiplication\n4) Division")
	flag.IntVar(&difChoiceInt, "level", 0, "Level Choice\n1) Beginner\n2) Intermediate\n3) Advanced\n4) Expert")
	flag.IntVar(&defaultNumofQuestions, "numq", defaultNumofQuestions, "How many questions you want to challenge, default is 20")
	flag.IntVar(&timeLimit, "time", timeLimit, "-Time limit per question\n-if you don't want to set a limit, just put 0, default is 5 secs per question")
	flag.Parse()
	if mathChoiceInt > 4 || mathChoiceInt < 1 || difChoiceInt > 4 || difChoiceInt < 1 || defaultNumofQuestions < 1 {
		fmt.Println(strings.Repeat("*", 60))
		fmt.Println("Mathematics Challenge Quiz (Train Your Calculation Speed)")
		fmt.Println(strings.Repeat("*", 60))
		fmt.Println(strings.Repeat("#", 15) + " Please Choose Below Category " + strings.Repeat("#", 15))

		// Display mathematics quiz menu list
		for k, v := range mathMenu {
			fmt.Printf("%d) %v\n", k+1, v)
		}
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
	// var
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
	timeperQ := timeLimit * defaultNumofQuestions
	timer := time.NewTimer(time.Duration(timeperQ) * time.Second)
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
				continue
			}
			inputVal, _ = strconv.Atoi(input.Text())
			if inputVal == v.answer {
				fmt.Printf("your answer %v is correct\n", inputVal)
				correct++
			} else {
				fmt.Printf("your answer is `%v` incorrect, the correct answer is %v\n", inputVal, v.answer)
			}
		}

	}
	fmt.Println(strings.Repeat("+", 50))
	fmt.Printf("Results: you have scored %v out of %v\n", correct, len(mathquestionlist))
	fmt.Printf("You have completed the quiz.\n")
}

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
