package sh

import (
	"strings"
	"testing"
)

const (
	testMessage      = "sample test message"
	testErrorMessage = "sample error message"
)

func TestShellSuccess(t *testing.T) {
	t.Run("TestShell", func(t *testing.T) {
		res, err := Shell("echo %s", testMessage)
		if err != nil {
			t.Errorf("expected no error, got command error: %v", err)
		}
		if !strings.Contains(res, testMessage) {
			t.Errorf("expected result message %s, got result : %v", testMessage, res)
		}
	})

	t.Run("TestShellMuteOutput", func(t *testing.T) {
		res, err := ShellMuteOutput("echo %s", testMessage)
		if err != nil {
			t.Errorf("expected no error, got command error: %v", err)
		}
		if !strings.Contains(res, testMessage) {
			t.Errorf("expected result message %s, got result : %v", testMessage, res)
		}
	})

	t.Run("TestShellMuteOutputError", func(t *testing.T) {
		res, err := ShellMuteOutputError("echo %s", testMessage)
		if err != nil {
			t.Errorf("expected no error, got command error: %v", err)
		}
		if !strings.Contains(res, testMessage) {
			t.Errorf("expected result message %s, got result : %v", testMessage, res)
		}
	})

	t.Run("TestShellSilent", func(t *testing.T) {
		res, err := ShellSilent("echo %s", testMessage)
		if err != nil {
			t.Errorf("expected no error, got command error: %v", err)
		}
		if !strings.Contains(res, testMessage) {
			t.Errorf("expected result message %s, got result : %v", testMessage, res)
		}
	})
}

func TestShellFail(t *testing.T) {
	t.Run("TestShell", func(t *testing.T) {
		_, err := Shell("echo %s && exit 1", testErrorMessage)
		if err == nil {
			t.Errorf("expected error message 'exit status 1', got command error nil")
		}
	})

	t.Run("TestShellMuteOutput", func(t *testing.T) {
		_, err := ShellMuteOutput("echo %s && exit 1", testErrorMessage)
		if err == nil {
			t.Errorf("expected error message 'exit status 1', got command error nil")
		}
	})

	t.Run("TestShellMuteOutputError", func(t *testing.T) {
		_, err := ShellMuteOutputError("echo %s && exit 1", testErrorMessage)
		if err == nil {
			t.Errorf("expected error message 'exit status 1', got command error nil")
		}
	})

	t.Run("TestShellSilent", func(t *testing.T) {
		_, err := ShellSilent("echo %s && exit 1", testErrorMessage)
		if err == nil {
			t.Errorf("expected error message 'exit status 1', got command error nil")
		}
	})
}
