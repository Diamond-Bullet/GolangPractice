package engineering

import (
	"GolangPractice/pkg/logger"
	"fmt"
	"syscall"
	"testing"
	"time"
)

// see more layouts in src/time/format.go
const (
	TimeFormatDate  = "2006-01-02"
	TimeFormatDate2 = "20060102"

	TimeFormatTime  = "2006-01-02 15:04:05"
	TimeFormatTime2 = "2006/01/02 15:04:05"

	TimeFormatMS = "2006-01-02 15:04:05.000"

	TimeFormatISO = "2006-01-02T15:04:05Z"
)

func TestTimeFormat(t *testing.T) {
	// Format current time.
	nowStamp := time.Now()

	timeStr := nowStamp.Format(TimeFormatMS)
	logger.Infof("Current Time: %s", timeStr)
}

func TestTimeParse(t *testing.T) {
	// Assign location.
	// time.Local, time.UTC.
	locIdentifier := "America/New_York"
	loc, _ := time.LoadLocation(locIdentifier)

	templateTime, _ := time.ParseInLocation(TimeFormatTime2, "2021/11/02 15:04:05", loc)
	logger.Infof("%s %s", locIdentifier, templateTime.String())

	// Specify time zone.
	logger.Info(time.Local.String(), templateTime.In(time.Local))
}

func TestTimeCalc(t *testing.T) {
	logger.Info(time.Now().AddDate(1, 1, 1))
	logger.Info(time.Now().AddDate(0, 0, -1))
	logger.Info(time.Now().AddDate(0, 0, -26))
}

func TestGetFileCreationTime(t *testing.T) {
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
