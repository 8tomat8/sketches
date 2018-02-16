package main

import (
	"errors"
	"fmt"
)

func main() {
}

func formatV(e error) error {
	return fmt.Errorf("AJHVGCueyvbcuabefia ahjdbke dasd: %v", e)
}

func formatS(e error) error {
	return fmt.Errorf("AJHVGCueyvbcuabefia ahjdbke dasd: %s", e)
}

func concat(e error) error {
	return errors.New("AJHVGCueyvbcuabefia ahjdbke dasd: " + e.Error())
}

func formatSWithFunc(e error) error {
	return fmt.Errorf("AJHVGCueyvbcuabefia ahjdbke dasd: %s", e.Error())
}

func formatVWithFunc(e error) error {
	return fmt.Errorf("AJHVGCueyvbcuabefia ahjdbke dasd: %v", e.Error())
}
