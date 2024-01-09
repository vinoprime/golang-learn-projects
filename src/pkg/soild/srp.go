package solid

import "fmt"

var entrtyCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entrtyCount++
	entry := fmt.Sprintf("%d: %s", entrtyCount, text)
	j.entries = append(j.entries, entry)
	return entrtyCount
}

func (j *Journal) RemoveEntry(id int) bool {
	// .... code
	return false
}

/* Single Responsibility Principle*/
