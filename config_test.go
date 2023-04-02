package main

import (
	"reflect"
	"testing"
)

func TestParseArgs(t *testing.T) {
	type args struct {
		argv []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test#1", args{[]string{"-v", "-i"}}, false},
		{"test#1", args{[]string{"-o"}}, true},
		{"test#1", args{[]string{"-otest"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := ParseArgs(tt.args.argv); (err != nil) != tt.wantErr {
				t.Errorf("ParseArgs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResolveConfig(t *testing.T) {
	type args struct {
		argv []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ResolveConfig(tt.args.argv)
		})
	}
}

func Test_getParameter(t *testing.T) {
	type args struct {
		flag string
		argv []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getParameter(tt.args.flag, tt.args.argv)
			if (err != nil) != tt.wantErr {
				t.Errorf("getParameter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getParameter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_param_SetVal(t *testing.T) {
	type fields struct {
		letter    string
		lname     string
		desc      ParamType
		valBool   bool
		valString string
	}
	type args struct {
		argv []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Param{
				letter:    tt.fields.letter,
				lname:     tt.fields.lname,
				desc:      tt.fields.desc,
				valBool:   tt.fields.valBool,
				valString: tt.fields.valString,
			}
			got, err := p.SetVal(tt.args.argv)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetVal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetVal() got = %v, want %v", got, tt.want)
			}
		})
	}
}
