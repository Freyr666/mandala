package main

import (
	"os"
	"os/exec"
	"log"
	"time"
	"bytes"
	"strconv"
)

type unitType int
const (
        File = iota
        Dir
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
	Open()        unit
	Status()      status
	Type()        unitType
}

type file string
type dir  []string

func strToUnit(name *string) unit {
	tmp, err := os.Open(*name)
	defer tmp.Close()
	if err != nil {log.Fatal(err)}

	info, err := tmp.Stat()
	if err != nil {log.Fatal(err)}

	if info.IsDir() {
		nms, err := tmp.Readdirnames(0)
		if err != nil {log.Fatal(err)}
		return (*dir)(&nms)
	} else {
		return (*file)(name)
	}
}

func (f *file) View(len int) []byte {
	tmp, err := os.Open(string(*f))
	defer tmp.Close()
	if err != nil {log.Fatal(err)}

	bin  := "cat"

	var b bytes.Buffer

	proc := exec.Command(bin, tmp.Name() )
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

func (f *file) Open() unit {
	tmp, err := os.Open(string(*f))
	defer tmp.Close()
	if err != nil {log.Fatal(err)}
	
	bin  := "vi"
	args := []string{tmp.Name()}

	proc, err := os.StartProcess(bin, args, nil)
	if err != nil {log.Fatal(err)}
	_, err = proc.Wait()
	if err != nil {log.Fatal(err)}
	
	return nil
}

func (f *file) Status() status {
	tmp, err := os.Open(string(*f))
	defer tmp.Close()
	if err != nil {log.Fatal(err)}
	
	info, err := tmp.Stat()	
	if err != nil {log.Fatal(err)}
	rval := status {
		info.Mode().String(),
		info.Size(),
		info.ModTime(),
		"user",
		"test"}
	return rval
}

func (f *file) Type() unitType {
	return File
}

func (d *dir) View(len int) []byte {
	
	return nil
}

func (d *dir) Open() unit {
	
	return nil
}

func (d *dir) Status() status {
	rval := status {
		"d------",
		123,
		time.Now(),
		"user",
		"test"}
	return rval
}

func (d *dir) Type() unitType {
	return Dir
}
