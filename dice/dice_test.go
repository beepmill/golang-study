package main

import (
	"testing"
)

func Test_parseRoll(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name      string
		args      args
		wantRolls int
		wantSize  int
		wantAdd   int
		wantErr   bool
	}{
		{
			"Parses a simple command",
			args{
				"1d6",
			},
			1,
			6,
			0,
			false,
		},
		{
			"Parses a command with addition",
			args{
				"3d12+123",
			},
			3,
			12,
			123,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRolls, gotSize, gotAdd, err := parseRoll(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseRoll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRolls != tt.wantRolls {
				t.Errorf("parseRoll() gotRolls = %v, want %v", gotRolls, tt.wantRolls)
			}
			if gotSize != tt.wantSize {
				t.Errorf("parseRoll() gotSize = %v, want %v", gotSize, tt.wantSize)
			}
			if gotAdd != tt.wantAdd {
				t.Errorf("parseRoll() gotAdd = %v, want %v", gotAdd, tt.wantAdd)
			}
		})
	}
}

func TestRoll(t *testing.T) {
	type args struct {
		command string
		seed    int64
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
		wantErr    bool
	}{
		{
			"it rolls...",
			args{
				"1d6",
				1, // Using a static seed
			},
			6,
			false,
		},
		{
			"it rolls...",
			args{
				"10d10+10",
				1, // Using a static seed
			},
			64,
			false,
		},
		{
			"it requires a parsable command...",
			args{
				"1dork6",
				1, // Using a static seed
			},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := Roll(tt.args.command, tt.args.seed)
			if (err != nil) != tt.wantErr {
				t.Errorf("Roll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("Roll() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
