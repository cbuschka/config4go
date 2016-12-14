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
	"unicode"
)

const (
	initial   = 0
	inKey     = 1
	postKey   = 2
	eqSeen    = 4
	inValue   = 3
	inComment = 5
	done      = 11
)

// ConfigParser is a parser capable to parse config streams.
type ConfigParser struct {
	state       int
	keyBuffer   bytes.Buffer
	valueBuffer bytes.Buffer
}

// Parse reads a config from reader in into map dest. Returns error
// in case of error.
func (configParser *ConfigParser) Parse(in *bufio.Reader) (map[string]string, error) {
	dest := make(map[string]string)

	configParser.state = initial

	for configParser.state != done {
		rune, _, err := in.ReadRune()
		if err != nil && err != io.EOF {
			return nil, err
		}

		switch configParser.state {
		case initial:
			if err == io.EOF {
				configParser.state = done
			} else if '#' == rune {
				configParser.state = inComment
			} else if unicode.IsSpace(rune) {
				// skip
			} else if unicode.IsLetter(rune) || '_' == rune {
				configParser.keyBuffer.WriteRune(rune)
				configParser.state = inKey
			} else {
				return nil, errors.New("Invalid input.")
			}
			break
		case inKey:
			if err == io.EOF {
				return nil, errors.New("Unexpected end of input.")
			} else if unicode.IsSpace(rune) {
				configParser.state = postKey
			} else if '=' == rune {
				configParser.state = eqSeen
			} else if unicode.IsDigit(rune) || unicode.IsLetter(rune) || '_' == rune {
				configParser.keyBuffer.WriteRune(rune)
			} else {
				return nil, errors.New("Invalid input.")
			}
			break
		case postKey:
			if err == io.EOF {
				return nil, errors.New("Unexpected end of input.")
			} else if unicode.IsSpace(rune) {
				// skip
			} else if '=' == rune {
				configParser.state = eqSeen
			} else {
				return nil, errors.New("Invalid input.")
			}
			break
		case eqSeen:
			if err == io.EOF {
				key := configParser.keyBuffer.String()
				value := configParser.valueBuffer.String()
				dest[key] = value

				configParser.state = done
			} else if unicode.IsSpace(rune) {
				// skip
			} else if '\n' == rune {
				key := configParser.keyBuffer.String()
				value := configParser.valueBuffer.String()
				dest[key] = value

				configParser.keyBuffer.Truncate(0)
				configParser.valueBuffer.Truncate(0)

				configParser.state = initial
			} else {
				configParser.valueBuffer.WriteRune(rune)
			}
			break
		case inValue:
			if err == io.EOF {
				key := configParser.keyBuffer.String()
				value := configParser.valueBuffer.String()
				dest[key] = value

				configParser.state = done
			} else if '\n' == rune {
				key := configParser.keyBuffer.String()
				value := configParser.valueBuffer.String()
				dest[key] = value

				configParser.keyBuffer.Truncate(0)
				configParser.valueBuffer.Truncate(0)

				configParser.state = initial
			} else {
				configParser.valueBuffer.WriteRune(rune)
			}
			break
		case inComment:
			if '\n' == rune {
				configParser.state = inKey
			}
			break
		case done:
			return nil, errors.New("Invalid state.")
		default:
			return nil, errors.New("Invalid input.")
		}
	}

	return dest, nil
}
