package searchengine

import (
	"strconv"
	"strings"
)

// feature 7 Search engine has many services that are chained or recursive. For optimization efforts,
// calculate the individual time taken by each service to run.

// Solution
// Using stack

func serviceTime(n int, logs []string) []int {
	var stack Stack

	funcTimes := make([]int, n)
	delim := ":"

	funcStr := strings.Split(logs[0], delim)
	temp, _ := strconv.Atoi(funcStr[0])
	stack = stack.Push(temp)
	time, _ := strconv.Atoi(funcStr[2])
	for i := 0; i < len(logs); i++ {
		funcStr = strings.Split(logs[i], delim)
		if strings.Contains(funcStr[1], "start") {
			if !stack.Empty() {
				temp, _ = strconv.Atoi(funcStr[2])
				funcTimes[stack.Top()] += (temp - time)
			}
			temp, _ = strconv.Atoi(funcStr[0])
			stack = stack.Push(temp)
			time, _ = strconv.Atoi(funcStr[2])
		} else {
			temp, _ = strconv.Atoi(funcStr[2])
			funcTimes[stack.Top()] += (temp - time + 1)
			stack, _ = stack.Pop()
			time = temp + 1
		}
	}
	return funcTimes
}
