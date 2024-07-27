package main

import (
	"GolangPractice/utils/logger"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

/*
quick start:
go_pkg_cleaner -dir /tmp -deadline 20240606
*/

var (
	path       = flag.String("dir", "", "the PATH you want to clear")
	deadline   = flag.String("deadline", "", "packages created before the date are going to be cleared. Layout is like 20060102")
	safePeriod = flag.String("safe", "", "only if there is a package's creation time later than `date+safe`, older ones can be eliminated. It's measured in days.")
)


const (
	DefaultSafePeriod = time.Hour * 24 * 15 // 15 days
	TimeLayOut = "20060102"
)

func main() {
	ts, err := parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	doubleConfirmation()


	packageCleaner := NewPackageCleaner(*path, ts)

	err = packageCleaner.Clean()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func parse() (time.Time, error) {
	flag.Parse()

	ts, err := time.Parse(TimeLayOut, *deadline)
	if err != nil {
		fmt.Println("wrong date get:", deadline)
		fmt.Println(err)
		return time.Time{}, err
	}

	return ts, nil
}

func doubleConfirmation() {
	fmt.Printf("this is the directory you want to clear:%s\n confirm it[y/n] and click [Enter] key to submit your choice", *path)
	confirm, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		logger.Errorln("program unfortunately exit:", err)
		os.Exit(-1)
	}
	if confirm != "y\n" {
		os.Exit(0)
	}
}

type PackageCleaner struct {
	Path       string
	Deadline   time.Time
	SafePeriod time.Duration
	DelWait     map[string][]string
	SafeDelMark map[string]bool
}

func NewPackageCleaner(path string, deadline time.Time) *PackageCleaner {
	return &PackageCleaner{
		Path:        path,
		Deadline:    deadline,
		SafePeriod:  DefaultSafePeriod,
		DelWait:     map[string][]string{},
		SafeDelMark: map[string]bool{},
	}
}

func (p *PackageCleaner) Clean() error {
	categories, err := os.ReadDir(p.Path)
	if err != nil {
		logger.Errorln("read directory fail:", err)
		return err
	}

	for _, category := range categories {
		p.Determine(category)
	}

	p.Remove()
	return nil
}

func (p *PackageCleaner) Determine(dir os.DirEntry) {
	info, err := dir.Info()
	if err != nil {
		logger.Errorf("addDir err: %s, %v", err.Error(), dir)
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
	if createTime.After(p.Deadline.Add(p.SafePeriod)) {
		p.SafeDelMark[fileSymbol] = true
	} else if createTime.Before(p.Deadline) {
		p.DelWait[fileSymbol] = append(p.DelWait[fileSymbol], info.Name())
	}
}

func (p *PackageCleaner) Remove() {
	for fileSymbol := range p.SafeDelMark {
		for _, file := range p.DelWait[fileSymbol] {
			err := os.RemoveAll(p.Path + "/" + file)
			if err != nil {
				fmt.Printf("[ERROR] remove file %v err:%v\n", file, err)
				continue
			}
			fmt.Println("[INFO] file", file, "has been removed")
		}
	}
}