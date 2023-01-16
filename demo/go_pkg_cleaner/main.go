package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

/*
quick start:
go_pkg_cleaner -dir /tmp -date 2021-11-20
*/

const (
	DefaultSafePeriod = time.Hour * 24 * 15
)

var (
	path       = flag.String("dir", "", "the PATH you want to clear")
	deadline   = flag.String("date", "", "packages created before the date are going to be cleared. Layout is like 2006-01-02")
	safePeriod = flag.String("safe", "", "only if there is a package's creation time later than `date+safe`, older ones can be eliminated. It's measured in days.")
)

func main() {
	ts := parse()
	categories, err := os.ReadDir(*path)
	if err != nil {
		fmt.Println("read directory fail:", err)
		os.Exit(-1)
	}

	dc := NewDirControl(*path, ts)

	for _, d := range categories {
		dc.AddDir(d)
	}

	dc.DoDel()
}

func parse() time.Time {
	flag.Parse()

	ts, err := time.Parse("2006-01-02", *deadline)
	if err != nil {
		fmt.Println("wrong date get:", deadline)
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Printf("this is the directory you want to clear:%s\n confirm it[y/n] and click [Enter] key to submit your choice", *path)
	confirm, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("program unfortunately exit:", err)
		os.Exit(-1)
	}
	if confirm != "y\n" {
		os.Exit(0)
	}

	return ts
}

type DirControl struct {
	Path        string
	DelDate     time.Time
	SafePeriod  time.Duration
	DelWait     map[string][]string
	SafeDelMark map[string]bool
}

func NewDirControl(path string, delDate time.Time) *DirControl {
	return &DirControl{
		Path:        path,
		DelDate:     delDate,
		SafePeriod:  DefaultSafePeriod,
		DelWait:     map[string][]string{},
		SafeDelMark: map[string]bool{},
	}
}

func (d *DirControl) AddDir(dir os.DirEntry) {
	info, err := dir.Info()
	if err != nil {
		fmt.Printf("addDir err: %v, %v\n", err, dir)
		return
	}

	fileMeta := strings.SplitN(info.Name(), "@", 2)
	if len(fileMeta) != 2 {
		return
	}

	// Get creation time from file name like v0.0.0-20220606... or v0.0.0-0.20220606
	timeInfo := strings.SplitN(fileMeta[1], "-", 2)
	if len(timeInfo) <= 1 || len(timeInfo[1]) < 8 {
		return
	}
	timeDate := timeInfo[1][:8]
	if timeInfo[1][:2] == "0." && len(timeInfo[1]) >= 10 {
		timeDate = timeInfo[1][2:10]
	}
	createTime, err := time.Parse("20060102", timeDate)
	if err != nil {
		fmt.Println(fileMeta[1], "Parse creation time err: ", err)
		return
	}

	fileSymbol := fileMeta[0]
	if createTime.After(d.DelDate.Add(d.SafePeriod)) {
		d.SafeDelMark[fileSymbol] = true
	} else if createTime.Before(d.DelDate) {
		d.DelWait[fileSymbol] = append(d.DelWait[fileSymbol], info.Name())
	}
}

func (d *DirControl) DoDel() {
	for fileSymbol := range d.SafeDelMark {
		for _, file := range d.DelWait[fileSymbol] {
			err := os.RemoveAll(d.Path + "/" + file)
			if err != nil {
				fmt.Printf("[ERROR] remove file %v err:%v\n", file, err)
				continue
			}
			fmt.Println("[INFO] file", file, "has been removed")
		}
	}
}
