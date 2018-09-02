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
