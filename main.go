package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(match(readInput()))
}

func match(reg string, target string) bool {
	if reg == "" && target == "" {
		return true
	}

	if reg == "" && target != "" {
		return false
	}

	if reg != "" && reg[0] == '^' {
		return compareTokens(reg[1:], target, false)
	}

	if target == "" {
		return compareTokens(reg, target, false)
	}

	// if prev comparison did match. Cut target by one symbol each itteration
	for i := 0; i < len(target); i++ {
		// comparse regex with reduced verion of target
		res := compareTokens(reg, target[i:], false)
		if res {
			return true
		}
	}
	return false
}

func readInput() (reg string, target string) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	in := sc.Text()
	splited := strings.Split(in, "|")
	reg = splited[0]
	target = splited[1]
	return
}

func check(r, t byte, escape bool) bool {
	return (!escape && r == '.') || r == t
}

func compareTokens(reg, target string, escape bool) bool {
	// reg is consumed. Means sucsess match
	if len(reg) == 0 {
		return true
	}
	if reg == "$" && !escape {
		return len(target) == 0
	}
	if reg[0] == '\\' && !escape {
		return compareTokens(reg[1:], target, true)
	}

	if len(reg) > 1 && !escape {
		switch c := reg[1]; c {
		case '?':
			return compareWithOptional(reg, target)
		case '*':
			return compareWithStar(reg, target)
		case '+':
			return compareWithPlus(reg, target)
		}
	}

	// reg is not consumed, but target is consumed, mean unsecsess match
	if len(target) == 0 {
		return false
	}

	if !check(reg[0], target[0], escape) {
		return false
	}
	return compareTokens(reg[1:], target[1:], false)
}

func compareWithOptional(reg, target string) bool {
	nextTarget := target
	if len(target) > 0 && check(reg[0], target[0], false) {
		nextTarget = target[1:]
	}
	return compareTokens(reg[2:], nextTarget, false)
}

func compareWithStar(reg, target string) bool {
	for i := 0; i < len(target); i++ {
		if !check(reg[0], target[i], false) {
			break
		}
		if compareTokens(reg[2:], target[i+1:], false) {
			return true
		}
	}
	return compareTokens(reg[2:], target, false)
}

func compareWithPlus(reg, target string) bool {
	if len(target) == 0 || !check(reg[0], target[0], false) {
		return false
	}
	return compareWithStar(reg, target)
}
