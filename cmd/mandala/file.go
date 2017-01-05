package main

import (
	"os"
	"os/exec"
	"fmt"
	"log"
	"time"
	"bytes"
	"strconv"
	"strings"
	"path/filepath"
)

type status struct {
	Mode  string
	Size  int64
	Date  time.Time
	Owner string
	Group string
}

type unit interface {
	View(len int) []byte  
	Open()        *[]file
	Status()      *status
	IsDir()       bool
}

type file struct {
	Name string
	Path string
}

func (s *status)String() string {
	return fmt.Sprintf("%s %d %s:%s %s", s.Mode, s.Size, s.Owner, s.Group, s.Date.String())
}

func strToFile(name string) file {
	path, err := filepath.Abs(name)
	if err != nil {log.Fatal(err)}
	
	return file {
		Name: name,
		Path: path}
}

func (f *file) View(len int) []byte {
	if f.IsDir() {
		tmp, err := os.Open(f.Path)
		defer tmp.Close()
		if err != nil {log.Fatal(err)}

		names, err := tmp.Readdirnames(len)
		if err != nil {log.Fatal(err)}

		return []byte(strings.Join(names, "\n"))
	} else {

		bin  := "cat"

		var b bytes.Buffer

		proc := exec.Command(bin, f.Path )
		filt := exec.Command("head", "-n", strconv.Itoa(len))
	
		procOut, err := proc.StdoutPipe()
		
		filt.Stdin  = procOut
		filt.Stdout = &b

		proc.Start()
		if err != nil {log.Fatal(err)}
		filt.Run()
		if err != nil {log.Fatal(err)}
	
		return b.Bytes()
	}
}

func (f *file) Open() *[]file {
	if f.IsDir() {
		tmp, err := os.Open(f.Path)
		defer tmp.Close()
		if err != nil {log.Fatal(err)}

		names, err := tmp.Readdirnames(0)
		if err != nil {log.Fatal(err)}

		rval := make([]file, len(names))
		
		for i, name := range(names) {
			rval[i] = strToFile(name)
		}
		return &rval
	} else {
		bin  := "vi"

		proc := exec.Command(bin, f.Path)
		proc.Stdin  = os.Stdin
		proc.Stdout = os.Stdout
		proc.Stderr = os.Stderr
		err := proc.Start()
		if err != nil {log.Fatal(err)}
		err = proc.Wait()
		if err != nil {
			log.Printf("Error while editing. Error: %v\n", err)
		} else {
			log.Printf("Successfully edited.")
		}
		return nil
	}
}

func (f *file) Status() *status {
	info, err := os.Stat(f.Path)
	if err != nil {log.Fatal(err)}

	rval := &status {
		info.Mode().String(),
		info.Size(),
		info.ModTime(),
		"user",
		"test"}
	return rval
}

func (f *file) IsDir() bool {
	info, err := os.Stat(f.Path)
	if err != nil {log.Fatal(err)}
	
	return info.IsDir()
}

func (f *file) String() string {
	return f.Name
}
