package main

import (
	"reflect"
	"testing"
)

func Test_rollSet_parseCommand(t *testing.T) {
	type fields struct {
		Command string
		Rolls   int
		Size    int
		Bonus   int
		Keep    int
	}
	tests := []struct {
		name    string
		fields  fields
		wantRs  rollSet
		wantErr bool
	}{
		{
			"Parses a simple command",
			fields{Command: "1d6"},
			rollSet{"1d6", 1, 6, 0, 0},
			false,
		},
		{
			"Fails to parse an invalid command",
			fields{Command: "1dork6"},
			rollSet{"1dork6", 0, 0, 0, 0},
			true,
		},
		{
			"Parses a command with addition",
			fields{Command: "3d12+123"},
			rollSet{"3d12+123", 3, 12, 123, 0},
			false,
		},
		{
			"Parses a command with keep",
			fields{Command: "3d12+123k2"},
			rollSet{"3d12+123k2", 3, 12, 123, 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &rollSet{
				Command: tt.fields.Command,
				Rolls:   tt.fields.Rolls,
				Size:    tt.fields.Size,
				Bonus:   tt.fields.Bonus,
				Keep:    tt.fields.Keep,
			}
			if err := rs.parseCommand(); (err != nil) != tt.wantErr {
				t.Errorf("rollSet.parseCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(*rs, tt.wantRs) {
				t.Errorf("parseCommand() = %v, want %v", *rs, tt.wantRs)
			}
			if err := rs.parseCommand(); (err != nil) != tt.wantErr {
				t.Errorf("rollSet.parseCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_rollSet_roll(t *testing.T) {
	type fields struct {
		Command string
		Rolls   int
		Size    int
		Bonus   int
		Keep    int
	}
	type args struct {
		seed int64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult int
		wantErr    bool
	}{
		{
			"it rolls...",
			fields{"1d6", 1, 6, 0, 0},
			args{1},
			6,
			false,
		},
		{
			"it rolls...",
			fields{"10d10+10", 10, 10, 10, 0},
			args{1},
			64,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := &rollSet{
				Command: tt.fields.Command,
				Rolls:   tt.fields.Rolls,
				Size:    tt.fields.Size,
				Bonus:   tt.fields.Bonus,
				Keep:    tt.fields.Keep,
			}
			gotResult, err := rs.roll(tt.args.seed)
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
