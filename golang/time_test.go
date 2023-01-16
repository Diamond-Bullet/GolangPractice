package golang

import (
	"fmt"
	"syscall"
	"testing"
	"time"
)

// reference: https://zhuanlan.zhihu.com/p/324922044

const (
	timeFormatDay = "2006-01-02"

	TimeFormatSecond = "2006-01-02 15:04:05"

	TimeFormatLog = "2006/01/02 15:04:05"

	TimeFormatMillisecond = "2006-01-02 15:04:05.000"

	TimeFormatISO = "2006-01-02T15:04:05Z"
)

func TestTime(t *testing.T) {
	// Parse current time.
	ti := time.Now()
	t1 := ti.Format(TimeFormatMillisecond)
	fmt.Printf("Current Time: %s\n", t1)

	// Assign location.
	loc, _ := time.LoadLocation("Local") // UTC， or other valid location name
	// another way to get local time zone is time.Local.
	// UTC is accessible if using time.UTC.
	t2, _ := time.ParseInLocation(TimeFormatLog, "2021/11/02 15:04:05", loc)
	fmt.Printf("Parse string as time.Time: %s\n", t2)
}

func TestGetCTime(t *testing.T) {
	var st syscall.Stat_t
	fileFullPath := "/root/demo.txt"
	err := syscall.Stat(fileFullPath, &st)
	if err != nil {
		fmt.Println(err)
		return
	}
	createTime := time.Unix(st.Ctim.Sec, 0)
	fmt.Println("file is created at ", createTime)
}
