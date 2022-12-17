package test

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

// int to string
func TestIntToString(t *testing.T) {
	vint := 3
	// int to string
	itoa := strconv.Itoa(vint)
	t.Log(fmt.Sprintf("int to string: %s %#v", itoa, itoa))
}

// byte to string
func TestByteToString(t *testing.T) {
	vbyte := []byte("hello")
	// byte to string
	btos := string(vbyte)
	t.Log(fmt.Sprintf("byte to string: %s %#v", btos, btos))
}

// string to byte
func TestStringToByte(t *testing.T) {
	vstring := "hello"
	// string to byte
	stob := []byte(vstring)
	t.Log(fmt.Sprintf("string to byte: %s %#v", stob, stob))
}

// btye to map
func TestByteToMap(t *testing.T) {
	vbyte := []byte(`{"name":"hello"}`)
	// byte to map
	var btom map[string]interface{}
	err := json.Unmarshal(vbyte, &btom)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Sprintf("byte to map: %#v %#v", btom, btom["name"]))
}
