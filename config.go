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
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

// Config is a config read by ConfigReader.
type Config struct {
	valueMap map[string]string
}

// ToMap returns the config as a new map.
func (config *Config) ToMap() map[string]string {
	newMap := make(map[string]string)
	for key, value := range config.valueMap {
		newMap[key] = value
	}

	return newMap
}

// Fill fills the target struct.
func (config *Config) Fill(target interface{}) {

	targetValue := reflect.ValueOf(target).Elem()

	for key, value := range config.valueMap {
		log.Printf("trying %s->%s\n", key, value)
		setField(targetValue, key, value)
	}
}

func setField(targetValue reflect.Value, name string, value interface{}) error {

	targetFieldValue := targetValue.FieldByName(strings.Title(name))
	if !targetFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s.", name)
	}

	if !targetFieldValue.CanSet() {
		return fmt.Errorf("Cannot set value of %s.", name)
	}

	targetFieldType := targetFieldValue.Type()
	val := reflect.ValueOf(value)
	if targetFieldType != val.Type() {
		return errors.New("Provided value type didn't match target field type.")
	}

	targetFieldValue.Set(val)
	return nil
}
