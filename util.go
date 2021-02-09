/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2018 Aliaksandr Valialkin
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 * Author: Aliaksandr Valialkin <valyala@gmail.com>
 */
package libconfig

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func s2b(s string) (b []byte) {
	strh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh.Data = strh.Data
	sh.Len = strh.Len
	sh.Cap = strh.Len
	return b
}

const maxStartEndStringLen = 80

func startEndString(s string) string {
	if len(s) <= maxStartEndStringLen {
		return s
	}
	start := s[:40]
	end := s[len(s)-40:]
	return start + "..." + end
}

func hex2dec(val string) int {
	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		return 0
	}
	return int(n)
}

func dec2hex(n int) string {
	return fmt.Sprintf("0x%X", n)
}

func scanMatch(path string) []string {
	dir := filepath.Dir(path)
	filenameMatching := filepath.Base(path)

	return scanMatchDir(dir, filenameMatching)
}

func matchFile(filename string, matching string) bool {
	matchs := strings.Split(matching, "*")

	if matchs[0] != "" {
		if !strings.HasPrefix(filename, matchs[0]) {
			return false
		}

		filename = strings.TrimPrefix(filename, matchs[0])
		matchs = matchs[1:]
	}
	if matchs[len(matchs)-1] != "" {
		if !strings.HasSuffix(filename, matchs[len(matchs)-1]) {
			return false
		}

		filename = strings.TrimSuffix(filename, matchs[len(matchs)-1])
		matchs = matchs[:len(matchs)-1]
	}

	for _, match := range matchs {
		n := strings.Index(filename, match)
		if n == -1 {
			return false
		}

		filename = filename[n:]
	}

	return true
}

func scanMatchDir(path string, matching string) (matchFiles []string) {

	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			//scanDir(path + "/" + file.Name(), matching)
			continue
		}

		if matchFile(file.Name(), matching) {
			matchFiles = append(matchFiles, path+"/"+file.Name())
		}
	}
	return
}
