package golang

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/buger/jsonparser"
	"github.com/gookit/color"
	"github.com/pkg/errors"
	"github.com/tealeg/xlsx"
)

// https://github.com/buger/jsonparser #Parse json data dynamically，support Get、Set。
func TestJsonParser(t *testing.T) {
	data := []byte(`{
  "person": {
    "name": {
      "first": "Leonid",
      "last": "Bugaev",
      "fullName": "Leonid Bugaev"
    },
    "github": {
      "handle": "buger",
      "followers": 109
    },
    "avatars": [
      { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
    ]
  },
  "company": {
    "name": "Acme"
  }
}`)

	val, err := jsonparser.Set(data, []byte("http://github.com"), "person", "avatars", "[0]", "backup_url")
	fmt.Println(string(val), " err: ", err)
}

// https://github.com/tealeg/xlsx #seems like the official's.
// https://github.com/qax-os/excelize is the most starred repo by now
func TestHandleExel(t *testing.T) {
	excelFileName := "/foo.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}

// https://github.com/pkg/errors
// see new feature of GO 2 on error handling.
func TestGoErrors(t *testing.T) {
	_, err := strconv.Atoi("12ab")
	if err != nil {
		println(errors.Wrap(err, "strconv.Atoi failed"))
	}
}

func TestColorfulPrint(t *testing.T) {
	color.Blueln("I am great")

	// work on linux or macOS.
	// https://stackoverflow.com/questions/5947742/how-to-change-the-output-color-of-echo-in-linux
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	colorWhite := "\033[37m"

	fmt.Println(colorRed, "test")
	fmt.Println(colorGreen, "test")
	fmt.Println(colorYellow, "test")
	fmt.Println(colorBlue, "test")
	fmt.Println(colorPurple, "test")
	fmt.Println(colorWhite, "test")
	fmt.Println(colorCyan, "test", colorReset)
	fmt.Println("next")
}
