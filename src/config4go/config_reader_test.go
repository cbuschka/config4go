/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2016 Cornelius Buschka.
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
 */

package config4go

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestReadIntoWithNewlineAtEof(t *testing.T) {
	configReader := NewConfigReader()
	in := bufio.NewReader(strings.NewReader("# empty\nkey = value\n"))
	config, err := configReader.ReadConfig(in)
	if err != nil {
		t.Error(err.Error())
	}

	dest := config.ToMap()

	if len(dest) != 1 {
		message := fmt.Sprintf("Dest is not of size 1, but %d, %s.", len(dest), dest)
		t.Error(message)
	}

	if dest["key"] != "value" {
		t.Error("Expected value of key to be 'value'.")
	}
}

func TestReadIntoNoNewlineAtEof(t *testing.T) {
	configReader := NewConfigReader()
	in := bufio.NewReader(strings.NewReader("# empty\nkey = value"))
	config, err := configReader.ReadConfig(in)
	if err != nil {
		t.Error(err.Error())
	}

	dest := config.ToMap()

	if len(dest) != 1 {
		message := fmt.Sprintf("Dest is not of size 1, but %d, %s.", len(dest), dest)
		t.Error(message)
	}

	if dest["key"] != "value" {
		t.Error("Expected value of key to be 'value'.")
	}
}
