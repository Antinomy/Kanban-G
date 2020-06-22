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
	"time"

	"github.com/c-bata/go-prompt"
)

const autoGitCounter int = 5

var gitcouter int = 0

var path string
var kanban kb.Kanban
var usingTaskItem kt.TaskItem = kt.UNKNOWN

// default
var IsShortMode bool = false

func main() {

	path = os.Args[1]

	refreshKanBan()

	kanbanPrompt := prompt.New(
		dummyExecutor,
		completer,
		prompt.OptionPrefix("InputCmd $"),
		prompt.OptionHistory([]string{"exit", "help"}),
		prompt.OptionPrefixTextColor(prompt.Yellow),
	)

CommandMode:
	for {
		inputCmd := kanbanPrompt.Input()
		inputCmd = strings.TrimLeft(inputCmd, " ")
		var cmds Cmds = buildCmd(inputCmd)

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
			refreshKanBan()
			continue

		case KAN:
			var taskItem kt.TaskItem = kt.GetTaskItem(cmds.param1)
			usingTaskItem = taskItem
			refreshKanView(kanban, usingTaskItem)
			continue

		case CREATE:
			var task = cmds.param1
			var banPrefix = "t"
			if cmds.length >= 3 {
				banPrefix = cmds.param2
			}
			kb.CreateBanTask(kanban, task, banPrefix)

			// refreshKanBan()
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

			// refreshKanBan()
			continue

		case CHANGEBAN:
			var taskKey = cmds.param1
			var banKey = cmds.param2
			var changeSpec kb.ChangeSpec = kb.ChangeBan(kanban, usingTaskItem, taskKey, banKey)

			var err = kb.ChangeOne(path, changeSpec)

			if err != nil {
				println(err)
			}

			// refreshKanBan()
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

				// refreshKanBan()
				continue
			}

		}

	}

}

func refreshKanBan() {
	kanban = kb.BuildKanban(path)
	kanban.IsShortMode = IsShortMode

	refreshKanView(kanban, usingTaskItem)

	autoGit(path)
}

func refreshKanView(kanban ban.Kanban, item kt.TaskItem) {
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
		{Text: "o [open]", Description: "o [open] $taskKey "},
		{Text: "s [short / shortmode]", Description: "short mode turn on/off"},
		{Text: "c [create]", Description: "c [create] taskname $banPrefix  "},
		{Text: "g [git]", Description: "commit & push to git "},
		{Text: "ct [changetask]", Description: "ct [changetask] $taskKey $TaskItem context"},
		{Text: "cb [changeban]", Description: "cb [changeban] $taskKey $banPrefix "},
		{Text: "k i", Description: "priority"},
		{Text: "k o", Description: "owner"},
		{Text: "k j", Description: "project"},
		{Text: "k d", Description: "deadline"},
		{Text: "h [help]", Description: "help doc"},
		{Text: "e [exit]", Description: "exit kanban"},
		{Text: "k [kan]", Description: "k [kan] <i / o / j / d>"},
		{Text: "r [rekan]", Description: "refresh kanban"},
	}
	return prompt.FilterHasPrefix(suggest, d.GetWordBeforeCursor(), true)
}

func dummyExecutor(in string) {}
