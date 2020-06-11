package main

import "strings"

type Cmds struct {
	cmd     string
	param1  string
	param2  string
	param3  string
	cmdType CmdType
	length  int
}

// TaskItem desc
type CmdType string

const (
	EXIT       CmdType = "EXIT"
	HELP       CmdType = "HELP"
	KAN        CmdType = "KAN"
	REKAN      CmdType = "REKAN"
	CREATE     CmdType = "CREATE"
	CHANGEBAN  CmdType = "CHANGEBAN"
	CHANGETASK CmdType = "CHANGETASK"

	UNKNOWN CmdType = "UNKNOWN"
)

func buildCmd(cmdstr string) Cmds {
	var result Cmds

	var cmdArray []string = strings.Split(cmdstr, " ")

	var length = len(cmdArray)
	result.length = length

	if length >= 1 {
		result.cmd = cmdArray[0]
	}

	if length >= 2 {
		result.param1 = cmdArray[1]
	}

	if length >= 3 {
		result.param2 = cmdArray[2]
	}

	if length >= 4 {
		result.param3 = cmdArray[3]
	}

	var cmdLower string = strings.ToLower(result.cmd)

	switch cmdLower {

	case "exit", "e":
		result.cmdType = EXIT

	case "help", "h":
		result.cmdType = HELP

	case "rekan", "r":
		result.cmdType = REKAN

	case "kan", "k":
		result.cmdType = KAN

	case "create", "c":
		result.cmdType = CREATE

	case "changeban", "cb":
		result.cmdType = CHANGEBAN

	case "changetask", "ct":
		if length >= 4 {
			result.cmdType = CHANGETASK
		}

	default:
		result.cmdType = UNKNOWN
	}

	if result.cmdType == UNKNOWN && length <= 1 {
		result.cmdType = REKAN
	}

	return result
}
