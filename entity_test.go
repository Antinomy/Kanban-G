package main

import (
	"testing"
)

func TestCreateCmd(t *testing.T) {

	var cmds Cmds

	cmds = buildCmd("e")
	if cmds.cmd != "e" || cmds.cmdType != EXIT {
		printError(cmds, t)
	}

	cmds = buildCmd("exit")
	if cmds.cmdType != EXIT {
		printError(cmds, t)
	}

	cmds = buildCmd("")
	if cmds.cmdType != REKAN {
		printError(cmds, t)
	}
}

func printError(cmds Cmds, t *testing.T) {
	t.Log(cmds)
	t.Errorf("Failed")
}
