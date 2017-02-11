package mygoutils

import (
	"encoding/gob"
	"bytes"
	"fmt"
)

func GetBytes(key interface{}) ([]byte, error) {
	gob.Register(map[string]interface{}{})
	gob.Register([]interface{}{})

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	return buf.Bytes(), nil
}
