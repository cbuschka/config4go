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
	"os"
)

// ConfigReader is a factory for configs from a source.
type ConfigReader struct {
}

// NewConfigReader creates a new ConfigReader.
func NewConfigReader() *ConfigReader {
	configReader := ConfigReader{}
	return &configReader
}

// ReadConfigFromFile reads a config from file fileName. Returns the
// config read or an error object in case of error.
func (configReader *ConfigReader) ReadConfigFromFile(fileName string) (*Config, error) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	in := bufio.NewReader(file)
	return configReader.ReadConfig(in)
}

// ReadConfig reads a config from in. Returns the config read or
// an error object in case of error.
func (configReader *ConfigReader) ReadConfig(in *bufio.Reader) (*Config, error) {
	parser := ConfigParser{}

	valueMap, err := parser.Parse(in)
	config := Config{valueMap: valueMap}

	return &config, err
}
