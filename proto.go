package main

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/tidwall/resp"
)

const (
	CommandSET = "SET"
)

type Command interface {
}

type SetCommand struct {
	// we fucking need what
	key string
	val string
}

func ParseCommand(raw string) (SetCommand, error) {
	rd := resp.NewReader(bytes.NewBufferString(raw))

	for {
		v, _, err := rd.ReadValue()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Read %s\n", v.Type())
		if v.Type() == resp.Array {
			for _, v := range v.Array() {
				switch v.String() {
				case CommandSET:

				default:

				}

			}
		}
	}
	return SetCommand{}, nil
}
