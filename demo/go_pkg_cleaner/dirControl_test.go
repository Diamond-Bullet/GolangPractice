package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestDirControl_AddDir(t *testing.T) {
	delTime, _ := time.Parse("2006-01-02 15:04:05", "")

	dirControl := NewDirControl("", delTime)

	type args struct {
		dir os.DirEntry
	}
	tests := []struct {
		name string
		args args
	}{
		{"0", args{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dirControl.AddDir(tt.args.dir)
		})
	}

	fmt.Println(dirControl)
}

func TestDirControl_DoDel(t *testing.T) {

	tests := []struct {
		name string
	}{
		{"0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// todo your files code
		})
	}
}
