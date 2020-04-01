
package sh

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)


// Shell run command on shell and get back output and error if get one
func Shell(format string, args ...interface{}) (string, error) {
	return sh(context.Background(), format, true, true, true, args...)
}

// ShellContext run command on shell and get back output and error if get one
func ShellContext(ctx context.Context, format string, args ...interface{}) (string, error) {
	return sh(ctx, format, true, true, true, args...)
}

// ShellMuteOutput run command on shell without logging the output
func ShellMuteOutput(format string, args ...interface{}) (string, error) {
	return sh(context.Background(), format, true, false, true, args...)
}

// ShellMuteOutputError run command on shell without logging the output or errors
func ShellMuteOutputError(format string, args ...interface{}) (string, error) {
	return sh(context.Background(), format, true, false, false, args...)
}

// ShellSilent runs commmand on shell without logging the command, output or errors
func ShellSilent(format string, args ...interface{}) (string, error) {
	return sh(context.Background(), format, false, false, false, args...)
}

// ShellBackground starts a background process and return the Process if succeed
func ShellBackground(format string, args ...interface{}) (*os.Process, error) {
	return shBackground(format, true, true, args...)
}

// ShellBackgroundMuteError starts a background process without logging errors
func ShellBackgroundMuteError(format string, args ...interface{}) (*os.Process, error) {
	return shBackground(format, true, false, args...)
}

// ShellBackgroundSilent starts a background process without logging the command or errors
func ShellBackgroundSilent(format string, args ...interface{}) (*os.Process, error) {
	return shBackground(format, false, false, args...)
}

func sh(ctx context.Context, format string, logCommand, logOutput, logError bool, args ...interface{}) (string, error) {
	command := fmt.Sprintf(format, args...)
	if logCommand {
		log.Printf("Running command: \n%s", command) 
	}
	c := exec.CommandContext(ctx, "sh", "-c", command)  // #nosec
	bytes, err := c.CombinedOutput()
	if logOutput {
		if output := strings.TrimSuffix(string(bytes), "\n"); len(output) > 0 {
			log.Printf("Command output: \n%s", output)
		}
	}

	if err != nil {
		if logError {
			log.Printf("Command error: \n%s", err)
		}
		return string(bytes), fmt.Errorf("command failed: %q %v", string(bytes), err)
	}
	return string(bytes), nil
}

func shBackground(format string, logCommand, logError bool, args ...interface{}) (*os.Process, error) {
	command := fmt.Sprintf(format, args...)
	if logCommand {
		log.Printf("Running command: \n%s", command) 
	}
	parts := strings.Split(command, " ")
	c := exec.Command(parts[0], parts[1:]...) // #nosec
	err := c.Start()
	if err != nil {
		if logError {
			log.Printf("Command error: \n%s", err)
		}
		return nil, err
	}
	return c.Process, nil
}
