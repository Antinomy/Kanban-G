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

	"github.com/c-bata/go-prompt"
)

const autoGitCounter int = 5

var gitcouter int = 0

func main() {

	var path = os.Args[1]
	var kanban kb.Kanban = kb.BuildKanban(path)

	refreshScreen(kanban, kt.UNKNOWN)

	var usingTaskItem kt.TaskItem = kt.UNKNOWN

	var IsShortMode bool = false

CommandMode:
	for {
		// default
		kanban.IsShortMode = IsShortMode

		t := prompt.Input("InputCmd $", completer)

		var cmds Cmds = buildCmd(t)

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

			autoGit(path)

			continue

		case GIT:
			lasyGit(path)
			gitcouter = 0
			continue

		case REKAN:
			kanban = kb.BuildKanban(path)
			kanban.IsShortMode = IsShortMode

			refreshScreen(kanban, usingTaskItem)

			autoGit(path)
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
	println("AutoGit  :", gitcouter, "/", autoGitCounter)
	kb.Kan(kanban, item)

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
}

func lasyGit(execPath string) {
	var lasyGitShell = execPath + "/lazyGit.sh"
	var commitStr = "GitSyncOn:[" + time.Now().Format("2006-01-02 15:04:05"+"]")
	err := kb.Exec(lasyGitShell, commitStr)

	if err != nil {
		println(err)
	}
}

func autoGit(execPath string) {

	if gitcouter >= autoGitCounter {
		gitcouter = 0
		go lasyGit(execPath)
		return
	}

	gitcouter++
}

func completer(d prompt.Document) []prompt.Suggest {
	suggest := []prompt.Suggest{
		{Text: "h [help]", Description: "help doc"},
		{Text: "e [exit]", Description: "exit kanban"},
		{Text: "k [kan]", Description: "k [kan] <i[priority] / o[owner] / j[project] /d[deadline]>"},
		{Text: "r [rekan]", Description: "refresh kanban"},
		{Text: "o [open]", Description: "c [create] $taskKey "},
		{Text: "s [short / shortmode]", Description: "short mode turn on/off"},
		{Text: "ct [changetask]", Description: "ct [changetask] $taskKey $TaskItem context"},
		{Text: "cb [changeban]", Description: "cb [changeban] $taskKey $banPrefix "},
		{Text: "c [create]", Description: "c [create] taskname $banPrefix  "},
		{Text: "g [git]", Description: "commit & push to git "},
	}
	return prompt.FilterHasPrefix(suggest, d.GetWordBeforeCursor(), true)
}
