package main

import (
	"reflect"
	"testing"
)

func Test_parseCommand(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantRs  rollSet
		wantErr bool
	}{
		{
			"Parses a simple command",
			args{
				"1d6",
			},
			rollSet{1, 6, 0},
			false,
		},
		{
			"Fails to parse an invalid command",
			args{
				"1dork6",
			},
			rollSet{0, 0, 0},
			true,
		},
		{
			"Parses a command with addition",
			args{
				"3d12+123",
			},
			rollSet{3, 12, 123},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRs, err := parseCommand(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRs, tt.wantRs) {
				t.Errorf("parseCommand() = %v, want %v", gotRs, tt.wantRs)
			}
		})
	}
}

func Test_rollSet_roll(t *testing.T) {
	type args struct {
		seed int64
	}
	tests := []struct {
		name       string
		rs         *rollSet
		args       args
		wantResult int
		wantErr    bool
	}{
		{
			"it rolls...",
			&rollSet{1, 6, 0},
			args{
				1, // Using a static seed
			},
			6,
			false,
		},
		{
			"it rolls...",
			&rollSet{10, 10, 10},
			args{
				1, // Using a static seed
			},
			64,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := tt.rs.roll(tt.args.seed)
			if (err != nil) != tt.wantErr {
				t.Errorf("rollSet.roll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("rollSet.roll() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
