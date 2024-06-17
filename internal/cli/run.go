package cli

import (
	"context"
	"errors"
	"io"
	"os/exec"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`\w|\%|^(\+|,|\.|:|\/|\\|=|_|-|[a-z]|[A-Z]|[0-9]|\s)+$`)

func ValidateCliArgument(arg string) bool {
	return re.Match([]byte(arg))
}

func ValidateCliArguments(args []string) bool {
	for _, p := range args {
		if !ValidateCliArgument(p) {
			return false
		}
	}
	return true
}

func RunThroughCli(ctx context.Context, command string, args []string) (string, string, error) {
	if !ValidateCliArguments(args) || !ValidateCliArgument(command) {
		return "", "", errors.New("failed to execute cli command: " + command + " " + strings.Join(args, " ") + " some arguments are not allowed")
	} else {
		cmd := exec.CommandContext(ctx, command, args...)
		cmdStr := cmd.String()

		pStdout, err := cmd.StdoutPipe()
		if err != nil {
			return "", cmdStr, err
		}

		pStderr, err := cmd.StderrPipe()
		if err != nil {
			return "", cmdStr, err
		}

		err = cmd.Start()
		if err != nil {
			return "", cmdStr, err
		}

		stdout, _ := io.ReadAll(pStdout)
		stderr, _ := io.ReadAll(pStderr)

		err = cmd.Wait()
		if err != nil {
			return "", cmdStr, errors.New(err.Error() + " | " + string(stderr) + " | " + string(stdout))
		}
		return string(stdout), cmdStr, nil
	}
}
