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
	"bytes"
	"errors"
	"io"
	"os"
	"unicode"
)

const (
	initial    = 0
	in_key     = 1
	post_key   = 2
	eq_seen    = 4
	in_value   = 3
	in_comment = 5
	done       = 11
)

func ReadFromFileInto(fileName string, dest map[string]string) error {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return err
	}

	in := bufio.NewReader(file)
	return ReadInto(in, dest)
}

func ReadInto(in *bufio.Reader, dest map[string]string) error {
	state := initial

	var keyBuffer bytes.Buffer
	var valueBuffer bytes.Buffer
	for state != done {
		rune, _, err := in.ReadRune()
		if err != nil && err != io.EOF {
			return err
		}

		switch state {
		case initial:
			if err == io.EOF {
				state = done
			} else if '#' == rune {
				state = in_comment
			} else if unicode.IsSpace(rune) {
				// skip
			} else if unicode.IsLetter(rune) || '_' == rune {
				keyBuffer.WriteRune(rune)
				state = in_key
			} else {
				return errors.New("Invalid input.")
			}
			break
		case in_key:
			if err == io.EOF {
				return errors.New("Unexpected end of input.")
			} else if unicode.IsSpace(rune) {
				state = post_key
			} else if '=' == rune {
				state = eq_seen
			} else if unicode.IsDigit(rune) || unicode.IsLetter(rune) || '_' == rune {
				keyBuffer.WriteRune(rune)
			} else {
				return errors.New("Invalid input.")
			}
			break
		case post_key:
			if err == io.EOF {
				return errors.New("Unexpected end of input.")
			} else if unicode.IsSpace(rune) {
				// skip
			} else if '=' == rune {
				state = eq_seen
			} else {
				return errors.New("Invalid input.")
			}
			break
		case eq_seen:
			if err == io.EOF {
				key := keyBuffer.String()
				value := valueBuffer.String()
				dest[key] = value

				state = done
			} else if unicode.IsSpace(rune) {
				// skip
			} else if '\n' == rune {
				key := keyBuffer.String()
				value := valueBuffer.String()
				dest[key] = value

				keyBuffer.Truncate(0)
				valueBuffer.Truncate(0)

				state = initial
			} else {
				valueBuffer.WriteRune(rune)
			}
			break
		case in_value:
			if err == io.EOF {
				key := keyBuffer.String()
				value := valueBuffer.String()
				dest[key] = value

				state = done
			} else if '\n' == rune {
				key := keyBuffer.String()
				value := valueBuffer.String()
				dest[key] = value

				keyBuffer.Truncate(0)
				valueBuffer.Truncate(0)

				state = initial
			} else {
				valueBuffer.WriteRune(rune)
			}
			break
		case in_comment:
			if '\n' == rune {
				state = in_key
			}
			break
		case done:
			return errors.New("Invalid state.")
		default:
			return errors.New("Invalid input.")
		}
	}

	return nil
}
