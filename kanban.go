package main

import (
	"bufio"
	kb "kanban/ban"
	kt "kanban/task"
	"os"
	"strings"
)

func main() {

	var path = os.Args[1]
	var kanban kb.Kanban = kb.BuildKanban(path)

	//	clear screen
	println("\033[H\033[2J")
	kb.Kan(kanban, kt.UNKNOWN)
	print("Input Cmd $ ")

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

		//	clear screen
		println("\033[H\033[2J")

		var cmds []string = strings.Split(input.Text(), " ")

		if len(cmds) <= 1 {
			kb.Kan(kanban, kt.UNKNOWN)
			print("Input Cmd $ ")
			continue
		}

		cmd = cmds[0]
		param1 = cmds[1]
		// param2 = cmds[2]

		if strings.ToLower(cmd) == "kan" {
			var taskItem kt.TaskItem = kt.GetTaskItem(param1)

			kb.Kan(kanban, taskItem)
			print("Input Cmd $ ")
		}
	}

}
