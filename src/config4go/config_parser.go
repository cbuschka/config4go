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
	dest        map[string]string
}

func (configParser *ConfigParser) reset() {
	configParser.dest = make(map[string]string)
	configParser.state = initial
}

// Parse reads a config from reader in into map dest. Returns error
// in case of error.
func (configParser *ConfigParser) Parse(in *bufio.Reader) (map[string]string, error) {
	configParser.reset()

	for configParser.state != done {
		symbol, _, err := in.ReadRune()
		if err != nil && err != io.EOF {
			return nil, err
		}

		if err := configParser.handleInput(symbol, err); err != nil {
			return nil, err
		}
	}

	return configParser.dest, nil
}

func (configParser *ConfigParser) handleInput(symbol rune, err error) error {
	switch configParser.state {
	case initial:
		return configParser.handleInitial(symbol, err)
	case inKey:
		return configParser.handleInKey(symbol, err)
	case postKey:
		return configParser.handlePostKey(symbol, err)
	case eqSeen:
		return configParser.handleEqSeen(symbol, err)
	case inValue:
		return configParser.handleInValue(symbol, err)
	case inComment:
		return configParser.handleInValue(symbol, err)
	case done:
		return errors.New("Invalid state.")
	default:
		return errors.New("Invalid input.")
	}

	return errors.New("Invalid state.")
}

func (configParser *ConfigParser) handleInitial(symbol rune, err error) error {
	if err == io.EOF {
		configParser.state = done
	} else if '#' == symbol {
		configParser.state = inComment
	} else if unicode.IsSpace(symbol) {
		// skip
	} else if unicode.IsLetter(symbol) || '_' == symbol {
		configParser.keyBuffer.WriteRune(symbol)
		configParser.state = inKey
	} else {
		return errors.New("Invalid input.")
	}

	return nil
}

func (configParser *ConfigParser) handleInKey(symbol rune, err error) error {
	if err == io.EOF {
		return errors.New("Unexpected end of input.")
	} else if unicode.IsSpace(symbol) {
		configParser.state = postKey
	} else if '=' == symbol {
		configParser.state = eqSeen
	} else if unicode.IsDigit(symbol) || unicode.IsLetter(symbol) || '_' == symbol {
		configParser.keyBuffer.WriteRune(symbol)
	} else {
		return errors.New("Invalid input.")
	}

	return nil
}

func (configParser *ConfigParser) handlePostKey(symbol rune, err error) error {

	if err == io.EOF {
		return errors.New("Unexpected end of input.")
	} else if unicode.IsSpace(symbol) {
		// skip
	} else if '=' == symbol {
		configParser.state = eqSeen
	} else {
		return errors.New("Invalid input.")
	}
	return nil
}

func (configParser *ConfigParser) handleEqSeen(symbol rune, err error) error {
	if err == io.EOF {
		configParser.addKeyValue()
		configParser.state = done
	} else if unicode.IsSpace(symbol) {
		// skip
	} else if '\n' == symbol {
		configParser.addKeyValue()
		configParser.state = initial
	} else {
		configParser.valueBuffer.WriteRune(symbol)
	}

	return nil
}

func (configParser *ConfigParser) handleInValue(symbol rune, err error) error {
	if err == io.EOF {
		configParser.addKeyValue()
		configParser.state = done
	} else if '\n' == symbol {
		configParser.addKeyValue()
		configParser.state = initial
	} else {
		configParser.valueBuffer.WriteRune(symbol)
	}

	return nil
}

func (configParser *ConfigParser) handleInComment(symbol rune, err error) error {
	if '\n' == symbol {
		configParser.state = inKey
	} else {
		// ok
	}

	return nil
}

func (configParser *ConfigParser) addKeyValue() {

	key := configParser.keyBuffer.String()
	value := configParser.valueBuffer.String()
	if key != "" {
		configParser.dest[key] = value
	}
	configParser.keyBuffer.Truncate(0)
	configParser.valueBuffer.Truncate(0)
}
