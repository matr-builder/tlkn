package tlkn

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Bash executes the given bash command string
func Bash(cmd string) error {
	cmd = trimLefts(cmd)
	c := BashCmd(cmd)
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	return c.Run()
}

// Bash executes the given bash command string
func BashCmd(cmd string) *exec.Cmd {
	cmd = trimLefts(cmd)
	bash := exec.Command("bash", "-c", cmd)
	bash.Stdout = os.Stdout
	bash.Stderr = os.Stderr
	return bash
}

// Prompt prompts user for input with default value.
func Prompt(prompt string, args ...interface{}) string {
	var s string
	fmt.Printf(prompt+": ", args...)
	fmt.Scanln(&s)
	return s
}

// PromptRequired prompts user for input with default value and requires an input.
func PromptRequired(prompt string, args ...interface{}) (s string) {
	for strings.Trim(s, " ") == "" {
		s = Prompt(prompt, args...)
	}
	return s
}

// PromptConfirm continues prompting until the input is boolean-ish.
func PromptConfirm(prompt string, args ...interface{}) bool {
	for {
		switch Prompt(prompt, args...) {
		case "Yes", "yes", "y", "Y":
			return true
		case "No", "no", "n", "N":
			return false
		}
	}
}

// PromptChoose prompts for a single selection from `list`, returning in the index.
func PromptChoose(prompt string, list []string) int {
	fmt.Println()
	for i, val := range list {
		fmt.Printf("  %d) %s\n", i+1, val)
	}

	fmt.Println()
	i := -1

	for {
		s := Prompt(prompt)
		// index
		n, err := strconv.Atoi(s)
		if err == nil {
			if n > 0 && n <= len(list) {
				i = n - 1
				break
			} else {
				continue
			}
		}

		// value
		i = indexOf(s, list)
		if i != -1 {
			break
		}
	}

	return i
}

func trimLefts(s string) string {
	nl := true
	return strings.Map(func(r rune) rune {
		// is space or tab at start of line
		if nl && r == 0x0020 || r == 0x0009 {
			return -1
		}

		// is newline
		if r == 0x000A {
			nl = true
			return r
		}

		nl = false
		return r
	}, s)
}

// index of `s` in `list`.
func indexOf(s string, list []string) int {
	for i, val := range list {
		if val == s {
			return i
		}
	}
	return -1
}
