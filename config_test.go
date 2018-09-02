// Package config contains helper methods to simplify working with command line arguments.
package config

import (
	"reflect"
	"testing"
)

func TestSplitArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    []string
		wantErr bool
	}{
		{name: "t1",
			args: []string{
				"-a -long-var=3",
				"-c",
				"",
				"-debug=true -k",
				"-o=out.txt"},
			want: []string{
				"-a", "-long-var", "3",
				"-c",
				"-debug", "true", "-k",
				"-o", "out.txt"},
			wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SplitArgs(tt.args)

			if (err != nil) != tt.wantErr {
				t.Errorf("SplitArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_discoverArgumentValue(t *testing.T) {
	type args struct {
		args         []string
		longName     string
		shortName    string
		defaultValue string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "short",
			args: args{
				args:         []string{"-k -c short.ini"},
				longName:     "",
				shortName:    "c",
				defaultValue: "default.ini"},
			want:    "short.ini",
			wantErr: false},
		{
			name: "long",
			args: args{
				args:         []string{"-k --config config.ini"},
				longName:     "config",
				shortName:    "",
				defaultValue: "default.ini"},
			want:    "config.ini",
			wantErr: false},
		{
			name: "default",
			args: args{
				args:         []string{"-k --noop"},
				longName:     "config",
				shortName:    "c",
				defaultValue: "default.ini"},
			want:    "default.ini",
			wantErr: false},
		{
			name: "duplicate_mix",
			args: args{
				args:         []string{"-k -c short.ini --config long.ini"},
				longName:     "config",
				shortName:    "c",
				defaultValue: "default.ini"},
			want:    "default.ini",
			wantErr: true},
		{
			name: "duplicate_short",
			args: args{
				args:         []string{"-k -c short.ini -f -c short2.ini"},
				longName:     "config",
				shortName:    "c",
				defaultValue: "default.ini"},
			want:    "default.ini",
			wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := discoverArgumentValue(tt.args.args, tt.args.longName, tt.args.shortName, tt.args.defaultValue)

			if (err != nil) != tt.wantErr {
				t.Errorf("discoverArgumentValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("discoverArgumentValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
