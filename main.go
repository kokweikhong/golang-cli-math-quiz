package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type mathematicsQuiz struct {
	question struct {
		number1 int
		number2 int
	}
	answer int
}

var (
	inputVal         int
	mathChoice       string
	difChoice        string
	symbolMath       string
	mathquestionlist []*mathematicsQuiz
)

// Addition, subtraction, multiplication and division
func main() {
	// multiplicationTable(2, 12)
	input := bufio.NewScanner(os.Stdin)
	fmt.Println(strings.Repeat("*", 60))
	fmt.Println("Mathematics Challenge Quiz (Train Your Calculation Speed)")
	fmt.Println(strings.Repeat("*", 60))
	fmt.Println(strings.Repeat("#", 15) + " Please Choose Below Category " + strings.Repeat("#", 15))

	// Display mathematics quiz menu list
	var mathMenu = []string{"Addition", "Subtraction", "Multiplication", "Division"}
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
	var difList = []string{"Beginner", "Intermediate", "Advanced", "Expert"}
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
	fmt.Println(mathChoice, difChoice)
	// var
	var correct int = 0
	switch mathChoice {
	case "Addition":
		symbolMath = "+"
		mathquestionlist = addition(difChoice, 20)
	case "Subtraction":
		symbolMath = "-"
		mathquestionlist = subtraction(difChoice, 20)
	case "Multiplication":
		symbolMath = "*"
		mathquestionlist = multiplicationdivision(mathChoice, difChoice, 20)
	case "Division":
		symbolMath = "/"
		mathquestionlist = multiplicationdivision(mathChoice, difChoice, 20)
	}

	for k, v := range mathquestionlist {
		fmt.Printf("Question #%d : %v %v %v = ", k+1, v.question.number1, symbolMath, v.question.number2)
		input.Scan()
		inputVal, _ = strconv.Atoi(input.Text())
		if input.Text() == "" {
			fmt.Printf("you have skip the question, continue to next question.\n")
			continue
		}
		if inputVal == v.answer {
			fmt.Printf("your answer %v is correct\n", inputVal)
			correct++
		} else {
			fmt.Printf("your answer is incorrect, the correct answer is %v\n", v.answer)
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
	fmt.Println(maxnum)
	for i := 0; i < defaultNumQ; i++ {
		subQuesList[i] = &mathematicsQuiz{}
		for {
			subQuesList[i].question.number1 = rand.Intn(maxnum)
			subQuesList[i].question.number2 = rand.Intn(maxnum)
			fmt.Println(subQuesList[i].question.number1)
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
		fmt.Println(subQuesList[i].question.number1)
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
			fmt.Println(x)
			fmt.Println(randList)
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
