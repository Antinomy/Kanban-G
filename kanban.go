package main

import (
	"bufio"
	"fmt"
	"kanban/ban"
	kb "kanban/ban"
	kt "kanban/task"
	"log"
	"os"
	"time"
)

const autoGitCounter int = 3

func main() {

	var path = os.Args[1]
	var kanban kb.Kanban = kb.BuildKanban(path)

	refreshScreen(kanban, kt.UNKNOWN)

	var usingTaskItem kt.TaskItem = kt.UNKNOWN

	input := bufio.NewScanner(os.Stdin)

	var IsShortMode bool = false

	var gitcouter int = 1

CommandMode:
	for input.Scan() {
		// default
		kanban.IsShortMode = IsShortMode

		var cmds Cmds = buildCmd(input.Text())

		switch cmds.cmdType {
		case EXIT:
			break CommandMode

		case SHORTMODE:
			IsShortMode = !IsShortMode
			fmt.Println("ShortMode TurnOn: ", IsShortMode)
			continue

		case HELP:
			printHelp()
			continue

		case OPEN:
			err := kb.OpenTask(kanban, cmds.param1, usingTaskItem)

			if err != nil {
				println(err)
			}

			autoGit(&gitcouter, path)

			continue

		case GIT:
			lasyGit(path)

			continue

		case REKAN:
			kanban = kb.BuildKanban(path)
			kanban.IsShortMode = IsShortMode
			refreshScreen(kanban, kt.UNKNOWN)

			autoGit(&gitcouter, path)

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

			autoGit(&gitcouter, path)

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

			autoGit(&gitcouter, path)

			continue

		case CHANGEBAN:
			var taskKey = cmds.param1
			var banKey = cmds.param2
			var changeSpec kb.ChangeSpec = kb.ChangeBan(kanban, usingTaskItem, taskKey, banKey)

			var err = kb.ChangeOne(path, changeSpec)

			if err != nil {
				println(err)
			}

			autoGit(&gitcouter, path)

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

				autoGit(&gitcouter, path)

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

func lasyGit(execPath string) {
	var lasyGitShell = execPath + "/lazyGit.sh"
	var commitStr = "'Git Sync On : [" + time.Now().Format("2006-01-02 15:04:05"+"]'")
	err := kb.Exec("/bin/bash", "-c", lasyGitShell+" "+commitStr)

	if err != nil {
		println(err)
	}
}

func autoGit(gitcouter *int, execPath string) {

	if *gitcouter > autoGitCounter {
		*gitcouter = 1
		lasyGit(execPath)
	}

	*gitcouter++
}
