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

func TestReadEmpty(t *testing.T) {
	configReader := NewConfigReader()
	in := bufio.NewReader(strings.NewReader(""))
	config, err := configReader.ReadConfig(in)
	if err != nil {
		t.Error(err.Error())
	}

	dest := config.ToMap()

	if len(dest) != 0 {
		message := fmt.Sprintf("Dest is not of size 0, but %d, %s.", len(dest), dest)
		t.Error(message)
	}
}

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

func TestReadMultipleKeyValues(t *testing.T) {
	configReader := NewConfigReader()
	in := bufio.NewReader(strings.NewReader("# empty\nkey1 = value1\nkey2=value2\n"))
	config, err := configReader.ReadConfig(in)
	if err != nil {
		t.Error(err.Error())
	}

	dest := config.ToMap()

	if len(dest) != 2 {
		message := fmt.Sprintf("Dest is not of size 2, but %d, %s.", len(dest), dest)
		t.Error(message)
	}

	if dest["key1"] != "value1" {
		t.Error("Expected value of key1 to be 'value1'.")
	}
	if dest["key2"] != "value2" {
		t.Error("Expected value of key2 to be 'value2'.")
	}
}
