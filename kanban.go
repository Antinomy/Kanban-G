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
		param2 string
		param3 string
	)

	var usingTaskItem kt.TaskItem = kt.UNKNOWN

	// fmt.Scanln(&cmd, &param1, &param2)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		//控制循环退出
		if input.Text() == "exit" || input.Text() == "e" {
			break
		}

		// default
		kanban.IsShortMode = false

		var cmds []string = strings.Split(input.Text(), " ")

		if len(cmds) >= 1 {
			cmd = cmds[0]
		}

		if strings.ToLower(cmd) == "rekan" || strings.ToLower(cmd) == "r" || len(cmds) <= 1 {
			kanban = kb.BuildKanban(path)
			refreshScreen(kanban, usingTaskItem)
			continue
		}

		if len(cmds) >= 2 {
			param1 = cmds[1]
		}

		if len(cmds) >= 3 {
			param2 = cmds[2]

			if strings.ToLower(param2) == "short" || strings.ToLower(param2) == "s" {
				kanban.IsShortMode = true
			}
		}

		if strings.ToLower(cmd) == "kan" || strings.ToLower(cmd) == "k" {
			var taskItem kt.TaskItem = kt.GetTaskItem(param1)
			usingTaskItem = taskItem
			refreshScreen(kanban, usingTaskItem)
			continue
		}

		if strings.ToLower(cmd) == "changeban" || strings.ToLower(cmd) == "cb" {
			var key = param1
			var banPrefix = param2
			var changeSpec kb.ChangeSpec = kb.ChangeBan(kanban, usingTaskItem, key, banPrefix)

			var err = kb.ChangeOne(path, changeSpec)

			if err != nil {
				println(err)
			}

			continue
		}

		if len(cmds) >= 4 {
			param3 = cmds[3]

			if strings.ToLower(cmd) == "changetask" || strings.ToLower(cmd) == "ct" {
				var key = param1
				var newTaskItem = param2
				var changeContent = param3

				var changeSpec = kb.ChangeTask(kanban, usingTaskItem, key, newTaskItem, changeContent)
				var err = kb.ChangeOne(path, changeSpec)

				if err != nil {
					println(err)
				}

				continue
			}
		}

	}

}

func refreshScreen(kanban ban.Kanban, item kt.TaskItem) {
	//	clear screen
	println("\033[H\033[2J")
	kb.Kan(kanban, item)
	print("Input Cmd $ ")

}
