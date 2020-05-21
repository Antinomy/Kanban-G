package ban

import (
	"testing"
)

func TestKan(t *testing.T) {

	var folderPath = ".././unittest/myTasks"

	var kanban Kanban = buildKanban(folderPath)

	kan(kanban)
}
