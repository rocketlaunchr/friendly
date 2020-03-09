// Copyright 2020 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func getCerts(gen bool) ([]byte, []byte, error) {

	if gen {
		return nil, nil, errors.New("generate key")
	}

	exePath, err := os.Executable()
	if err != nil {
		return nil, nil, err
	}
	exePath = filepath.Dir(exePath)

	content, err := ioutil.ReadFile(filepath.Join(exePath, "friendly.cert"))
	if err != nil {
		return nil, nil, err
	}

	certs := map[string]string{}
	err = json.Unmarshal(content, &certs)
	if err != nil {
		return nil, nil, err
	}

	return []byte(certs["public"]), []byte(certs["private"]), nil
}

func saveCerts(public []byte, private []byte) error {

	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	exePath = filepath.Dir(exePath)

	data, _ := json.Marshal(map[string]string{
		"private": string(private),
		"public":  string(public),
	})

	return ioutil.WriteFile(filepath.Join(exePath, "friendly.cert"), data, 0644)
}

func deleteCerts() error {
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	exePath = filepath.Dir(exePath)

	return os.Remove(filepath.Join(exePath, "friendly.cert"))
}
