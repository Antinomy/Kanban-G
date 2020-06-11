package main

import (
	"bufio"
	"fmt"
	"kanban/ban"
	kb "kanban/ban"
	kt "kanban/task"
	"log"
	"os"
	"strings"
)

func main() {

	var path = os.Args[1]
	var kanban kb.Kanban = kb.BuildKanban(path)

	refreshScreen(kanban, kt.UNKNOWN)

	var usingTaskItem kt.TaskItem = kt.UNKNOWN

	input := bufio.NewScanner(os.Stdin)

CommandMode:
	for input.Scan() {
		// default
		kanban.IsShortMode = false

		var cmds Cmds = buildCmd(input.Text())

		if cmds.length >= 3 {
			if strings.ToLower(cmds.param2) == "short" || strings.ToLower(cmds.param2) == "s" {
				kanban.IsShortMode = true
			}
		}

		switch cmds.cmdType {
		case EXIT:
			break CommandMode

		case HELP:
			printHelp()
			continue

		case REKAN:
			kanban = kb.BuildKanban(path)
			refreshScreen(kanban, usingTaskItem)
			continue

		case KAN:
			var taskItem kt.TaskItem = kt.GetTaskItem(cmds.param1)
			usingTaskItem = taskItem
			refreshScreen(kanban, usingTaskItem)
			continue

		case CREATE:
			var task = cmds.param1
			var banPrefix = "t"
			if cmds.length >= 3 {
				banPrefix = cmds.param2
			}
			kb.CreateBanTask(kanban, task, banPrefix)
			continue

		case CHANGETASK:
			var key = cmds.param1
			var newTaskItem = cmds.param2
			var changeContent = cmds.param3

			var changeSpec = kb.ChangeTask(kanban, usingTaskItem, key, newTaskItem, changeContent)
			var err = kb.ChangeOne(path, changeSpec)

			if err != nil {
				println(err)
			}

			continue

		case CHANGEBAN:
			var taskKey = cmds.param1
			var banKey = cmds.param2
			var changeSpec kb.ChangeSpec = kb.ChangeBan(kanban, usingTaskItem, taskKey, banKey)

			var err = kb.ChangeOne(path, changeSpec)

			if err != nil {
				println(err)
			}

			continue

		default:
			// change ban by ban key
			if cmds.length == 2 {
				var banKey = cmds.cmd
				var taskKey = cmds.param1
				var ban kb.Ban = kb.GetBan(kanban, banKey)

				if len(ban.Name) == 0 {
					continue
				}

				var changeSpec kb.ChangeSpec = kb.ChangeBan(kanban, usingTaskItem, taskKey, banKey)

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

func printHelp() {
	file, err := os.Open("README.md")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	print("Input Cmd $ ")
}
