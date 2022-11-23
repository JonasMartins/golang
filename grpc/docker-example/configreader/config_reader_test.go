// Copyright (c) 2022 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.
package configreader

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadEnv(t *testing.T) {
	testCases := []struct {
		name                 string
		mockGodotenvLoad     func(filenames ...string) (err error)
		mockEnvconfigProcess func(prefix string, spec interface{}) error
		expectedConfigReader *Config
		expectedError        error
	}{
		{
			name: "happy path",
			mockGodotenvLoad: func(filenames ...string) (err error) {
				return nil
			},
			mockEnvconfigProcess: func(prefix string, spec interface{}) error {
				return nil
			},
			expectedConfigReader: &Config{},
		},
		{
			name: "error reading .env file",
			mockGodotenvLoad: func(filenames ...string) (err error) {
				return errors.New("random error")
			},
			expectedError: errors.New("reading .env file: random error"),
		},
		{
			name: "error processing env vars",
			mockGodotenvLoad: func(filenames ...string) (err error) {
				return nil
			},
			mockEnvconfigProcess: func(prefix string, spec interface{}) error {
				return errors.New("random error")
			},
			expectedError: errors.New("processing env vars: random error"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			godotenvLoad = tc.mockGodotenvLoad
			envconfigProcess = tc.mockEnvconfigProcess
			c, err := ReadEnv()
			if err != nil {
				if tc.expectedError == nil {
					t.Fatalf(`expected no errors, got "%v"`, err)
				}
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				if tc.expectedError != nil {
					t.Fatalf(`expected error "%v", got nil`, err)
				}
			}
			require.Equal(t, tc.expectedConfigReader, c)
		})
	}
}
