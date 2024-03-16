package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArgs(t *testing.T) {
	var testCases = []struct {
		name           string
		args           []string
		spmArgs        []string
		commandAndArgs []string
	}{
		{
			name:           "simple",
			args:           []string{"ls"},
			spmArgs:        nil,
			commandAndArgs: []string{"ls"},
		},
		{
			name:           "help",
			args:           []string{"-h"},
			spmArgs:        []string{"-h"},
			commandAndArgs: nil,
		},
		{
			name:           "pid",
			args:           []string{"-pid", "/var/run/pid", "ls"},
			spmArgs:        []string{"-pid", "/var/run/pid"},
			commandAndArgs: []string{"ls"},
		},
		{
			name:           "pid with double dash",
			args:           []string{"-pid", "/var/run/pid", "--", "ls"},
			spmArgs:        []string{"-pid", "/var/run/pid"},
			commandAndArgs: []string{"ls"},
		},
		{
			name:           "pid with double dash * 2",
			args:           []string{"-pid", "/var/run/pid", "--", "--"},
			spmArgs:        []string{"-pid", "/var/run/pid"},
			commandAndArgs: []string{"--"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			spmArgs, commandAndArgs := parseArgs(tc.args)
			assert.Equal(t, spmArgs, tc.spmArgs)
			assert.Equal(t, commandAndArgs, tc.commandAndArgs)
		})
	}
}
