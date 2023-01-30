// Copyright (c) 2021 Veritas Technologies LLC. All rights reserved. IP63-2828-7171-04-15-9
package pm

import (
	"os"
	"testing"
)

func Test_getStatusColor(t *testing.T) {
	if os.Getenv("INTEGRATION_TEST") == "RUNNING" {
		t.Skip("Not applicable while running integration tests.")
		return
	}

	type args struct {
		status string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Start",
			args: args{status: dStatusStart},
			want: "blue",
		},
		{
			name: "Ok/Pass",
			args: args{status: dStatusOk},
			want: "green",
		},
		{
			name: "Fail",
			args: args{status: dStatusFail},
			want: "red",
		},
		{
			name: "Skip",
			args: args{status: dStatusSkip},
			want: "yellow",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStatusColor(tt.args.status); got != tt.want {
				t.Errorf("getStatusColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
