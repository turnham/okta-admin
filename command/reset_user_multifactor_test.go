package command

import (
	"strings"
	"testing"
)

func createTestResetUserMultifactorsCommand(globalOptsHelpText string) *ResetUserMultifactorsCommand {
	return &ResetUserMultifactorsCommand{
		Command: createTestCommand(globalOptsHelpText, "test_reset_user_mfa_cmd"),
	}
}

func TestResetUserMultifactorsCommand_Help(t *testing.T) {
	t.Parallel()
	const globalHelpMsg = `
Welcome to Hogwarts!
`
	c := createTestResetUserMultifactorsCommand(globalHelpMsg)
	if !strings.Contains(c.Help(), globalHelpMsg) {
		t.Errorf("Expected final help message to contain \"%s\"", globalHelpMsg)
	}
}

func TestResetUserMultifactorsCommand_ParseArgs(t *testing.T) {
	t.Parallel()

	c := createTestDeactivateUserCommand("")
	args := []string{
		"-email", "harry.potter@hogwarts.co.uk",
	}

	cfg, err := c.ParseArgs(args)
	if err != nil {
		t.Fatalf("Failed to parse arguments: %v", err)
	}

	if cfg.EmailID != args[1] {
		t.Errorf("Expected email id to be %s, received %s", args[1], cfg.EmailID)
	}
}
