package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

const (
	SORTBY_SIZE Key = 1 << iota
	SORTBY_TIME
	SORTBY_MODE
	SORTBY_NAME
	WITH_ATTR
	WITH_COLOR
	RENDER_COLUMN
	RENDER_LIST
)

var (
	OPTS int64

	OPTIONS = map[string]Opts{
		"M": SortOpt{i: SORTBY_MODE, Direction: SortDirectionUp},
		"S": SortOpt{i: SORTBY_SIZE, Direction: SortDirectionUp},
		"T": SortOpt{i: SORTBY_TIME, Direction: SortDirectionUp},
		"N": SortOpt{i: SORTBY_NAME, Direction: SortDirectionUp},

		"m": SortOpt{i: SORTBY_MODE, Direction: SortDirectionDown},
		"s": SortOpt{i: SORTBY_SIZE, Direction: SortDirectionDown},
		"t": SortOpt{i: SORTBY_TIME, Direction: SortDirectionDown},
		"n": SortOpt{i: SORTBY_NAME, Direction: SortDirectionDown},

		"a": Opt{i: WITH_ATTR},
		"c": Opt{i: WITH_COLOR},
		"1": Opt{i: RENDER_COLUMN},
		"l": Opt{i: RENDER_LIST},
	}
)

func main() {
	directory := "."
	if len(os.Args) > 1 {
		if !strings.HasPrefix(os.Args[len(os.Args)-1], "-") {
			directory = os.Args[len(os.Args)-1]
		}
	}

	for _, flag := range os.Args[1:] {
		if strings.HasPrefix(flag, "-") {
			for _, opt := range flag[1:] {
				OPTS = OPTS | OPTIONS[string(opt)].Int64()
			}
		}
	}

	f, err := os.Open(directory)
	if err != nil {
		log.Fatal(err)
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	files := Files(list)
	for _, o := range OPTIONS {
		if reflect.TypeOf(o).String() == "main.SortOpt" {
			SortByName(files, o.(SortOpt).Direction)
			if OPTS&SORTBY_MODE.Int64() == SORTBY_MODE.Int64() {
				SortByMode(files, o.(SortOpt).Direction)
			} else if OPTS&SORTBY_TIME.Int64() == SORTBY_TIME.Int64() {
				SortByTime(files, o.(SortOpt).Direction)
			} else if OPTS&SORTBY_SIZE.Int64() == SORTBY_SIZE.Int64() {
				SortBySize(files, o.(SortOpt).Direction)
			}
		}
	}

	Render(files)
}

type File struct {
	Name    string
	Attr    string
	ModTime time.Time
	Size    int64
}

func Render(files Files) {
	for _, file := range files {

		var name = file.Name()

		if OPTS&WITH_COLOR.Int64() == WITH_COLOR.Int64() {
			fn := strings.Split(name, ".")
			if len(fn) > 1 {
				if v, ok := EXT_MAPPING[fn[1]]; ok {
					name = fmt.Sprintf("%s%s%s", v, name, Clear)
				}
			}
		}

		if OPTS&WITH_ATTR.Int64() == WITH_ATTR.Int64() {
			fmt.Printf("%s %s | %8s |  %s\n", file.Mode(), file.ModTime().Format("Jan 02 2006 15:04:05"), humanize.Bytes(uint64(file.Size())), name)
		} else {
			fmt.Printf("%s\n", name)
		}
	}
}
