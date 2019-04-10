/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"testing"
)

func TestRoundUpBytes(t *testing.T) {
	tests := []struct {
		size     int64
		expected int64
	}{
		{
			int64(1500 * 1024 * 1024),
			int64(2 * GiB),
		},
	}
	for _, test := range tests {
		result := RoundUpBytes(test.size)
		if result != test.expected {
			t.Errorf("input: %d, RoundUpBytes result: %d, expected: %d", test.size, result, test.expected)
		}
	}
}

func TestRoundUpGiB(t *testing.T) {
	tests := []struct {
		size     int64
		expected int64
	}{
		{
			int64(1500 * 1024 * 1024),
			int64(2),
		},
	}
	for _, test := range tests {
		result := RoundUpGiB(test.size)
		if result != test.expected {
			t.Errorf("input: %d, RoundUpToGiB result: %d, expected: %d", test.size, result, test.expected)
		}
	}
}

func TestBytesToGiB(t *testing.T) {
	tests := []struct {
		size     int64
		expected int64
	}{
		{
			int64(2 * GiB),
			int64(2),
		},
	}
	for _, test := range tests {
		result := BytesToGiB(test.size)
		if result != test.expected {
			t.Errorf("input: %d, RoundUpToGiB result: %d, expected: %d", test.size, result, test.expected)
		}
	}
}

func TestGiBToBytes(t *testing.T) {
	tests := []struct {
		size     int64
		expected int64
	}{
		{
			int64(2),
			int64(2 * GiB),
		},
	}
	for _, test := range tests {
		result := GiBToBytes(test.size)
		if result != test.expected {
			t.Errorf("input: %d, RoundUpToGiB result: %d, expected: %d", test.size, result, test.expected)
		}
	}
}

func TestRoundUpSize(t *testing.T) {
	tests := []struct {
		size     int64
		unit     int64
		expected int64
	}{
		{
			int64(1500 * 1024 * 1024),
			int64(GiB),
			int64(2),
		},
	}
	for _, test := range tests {
		result := roundUpSize(test.size, test.unit)
		if result != test.expected {
			t.Errorf("input: %d, RoundUpSize result: %d, expected: %d", test.size, result, test.expected)
		}
	}
}
