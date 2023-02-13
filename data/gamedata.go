package data

import (
	"bytes"
	"encoding/gob"
)

func Fetch[T any](key string) (value T, err error) {
	row := db.QueryRow("SELECT value FROM GameData WHERE key = ?", key)

	var b []byte
	err = row.Scan(&b)
	if err != nil {
		return
	}

	dec := gob.NewDecoder(bytes.NewReader(b))
	err = dec.Decode(&value)
	return
}

func Store[T any](key string, value T) (err error) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err = enc.Encode(value)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT OR REPLACE INTO GameData (key, value) VALUES (?, ?)", key, b.Bytes())
	return
}
