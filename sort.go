package main

import (
	"io/fs"
	"sort"
)

type Files []fs.FileInfo

type SortDirection int

const (
	SortDirectionUp SortDirection = 1 << iota
	SortDirectionDown
)

func SortByMode(files Files, d SortDirection) {
	if d == SortDirectionDown {
		sort.Slice(files, func(i, j int) bool { return (files)[i].Mode() > (files)[j].Mode() })
		return
	}

	sort.Slice(files, func(i, j int) bool { return (files)[i].Mode() < (files)[j].Mode() })
}
func SortByTime(files Files, d SortDirection) {
	if d == SortDirectionDown {
		sort.Slice(files, func(i, j int) bool { return (files)[i].ModTime().After((files)[j].ModTime()) })
		return
	}

	sort.Slice(files, func(i, j int) bool { return (files)[i].ModTime().Before((files)[j].ModTime()) })
}
func SortBySize(files Files, d SortDirection) {
	if d == SortDirectionDown {
		sort.Slice(files, func(i, j int) bool { return (files)[i].Size() > (files)[j].Size() })
		return
	}

	sort.Slice(files, func(i, j int) bool { return (files)[i].Size() < (files)[j].Size() })
}
func SortByName(files Files, d SortDirection) {
	if d == SortDirectionDown {
		sort.Slice(files, func(i, j int) bool { return (files)[i].Name() > (files)[j].Name() })
		return
	}

	sort.Slice(files, func(i, j int) bool { return (files)[i].Name() < (files)[j].Name() })
}
