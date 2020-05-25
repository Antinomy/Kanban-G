package main

import (
	"bufio"
	"kanban/ban"
	kb "kanban/ban"
	kt "kanban/task"
	"os"
	"strings"
)

func main() {

	var path = os.Args[1]
	var kanban kb.Kanban = kb.BuildKanban(path)

	refreshScreen(kanban, kt.UNKNOWN)

	var (
		cmd    string
		param1 string
		// param2 string
	)

	// fmt.Scanln(&cmd, &param1, &param2)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		//控制循环退出
		if input.Text() == "exit" || input.Text() == "e" {
			break
		}

		var cmds []string = strings.Split(input.Text(), " ")

		if len(cmds) >= 1 {
			cmd = cmds[0]
		}

		if strings.ToLower(cmd) == "rekan" || strings.ToLower(cmd) == "r" {
			kanban = kb.BuildKanban(path)
			refreshScreen(kanban, kt.UNKNOWN)
			continue
		}

		if len(cmds) <= 1 {
			refreshScreen(kanban, kt.UNKNOWN)
			continue
		}

		if len(cmds) >= 2 {
			param1 = cmds[1]
		}
		// param2 = cmds[2]

		if strings.ToLower(cmd) == "kan" || strings.ToLower(cmd) == "k" {
			var taskItem kt.TaskItem = kt.GetTaskItem(param1)
			refreshScreen(kanban, taskItem)
		}
	}

}

func refreshScreen(kanban ban.Kanban, item kt.TaskItem) {
	//	clear screen
	println("\033[H\033[2J")
	kb.Kan(kanban, item)
	print("Input Cmd $ ")

}
