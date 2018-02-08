package tlkn

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

// Debug will enable debug logging
var Debug bool

// ExitError alias
type ExitError = exec.ExitError

func logDebug(v ...interface{}) {
	if !Debug {
		return
	}
	fmt.Println("[tlkn debug]\n", fmt.Sprint(v...))
}

// BashCmd creates a *exec.Cmd using the given bash command string
// writes to os.Stderr and os.Stdout by default
func BashCmd(ctx context.Context, cmd string) *exec.Cmd {
	cmd = trimLefts(cmd)
	bash := exec.CommandContext(ctx, "bash", "-c", cmd)
	bash.Stderr = os.Stderr
	bash.Stdout = os.Stdout
	return bash
}

// Bash creates an unexecuted BashCmd func
func Bash(ctx context.Context, cmd string) func() error {
	return BashCmd(ctx, cmd).Run
}

func Parallel(fns ...func() error) error {
	errCh := make(chan error)
	var wg sync.WaitGroup

	for _, fn := range fns {
		wg.Add(1)
		go func(fn func() error) {
			err := fn()
			if err != nil {
				errCh <- err
			}
			wg.Done()
		}(fn)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		return err
	}

	return nil
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
