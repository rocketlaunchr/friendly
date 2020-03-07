// Copyright 2020 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package main

type ThrowAway struct{}

func (*ThrowAway) Write(p []byte) (n int, err error) {
	return len(p), nil
}
