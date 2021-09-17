package netflix

import (
	"fmt"
	"testing"
)

func TestGroupTitles(t *testing.T) {
	titles := []string{"duel", "dule", "speed", "spede", "deul", "cars"}
	query := "spede"
	output := groupTitles(titles)
	for _, o := range output {
		if find(o, query) {
			fmt.Println("found anagrams")
			return
		}
	}
	t.Error("anagrams not found")
}
