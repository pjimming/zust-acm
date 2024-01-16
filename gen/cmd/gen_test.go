package main

import (
	"io"
	"os"
	"testing"
)

func TestReadTpl(t *testing.T) {
	f, err := os.Open("../dao.tpl")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fd, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	t.Log(string(fd))
}
