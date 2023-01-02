/*
*	Copyright (C) 2023 Kendall Tauser
*
*	This program is free software; you can redistribute it and/or modify
*	it under the terms of the GNU General Public License as published by
*	the Free Software Foundation; either version 2 of the License, or
*	(at your option) any later version.
*
*	This program is distributed in the hope that it will be useful,
*	but WITHOUT ANY WARRANTY; without even the implied warranty of
*	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*	GNU General Public License for more details.
*
*	You should have received a copy of the GNU General Public License along
*	with this program; if not, write to the Free Software Foundation, Inc.,
*	51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

package pkg

import (
	"fmt"
	"runtime"
	"testing"
)

func TestNewVersionString(t *testing.T) {
	type args struct {
		componentName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "morfic apiserver",
			args: args{
				componentName: "morfic-apiserver",
			},
			want: fmt.Sprintf("%s - Version: %s\nCommit: %s\nGo Build: %s\nBuild Date: %sOS: %s\nArchitecture: %s\nLicense: %s\n\n",
				"morfic-apiserver",
				"unknown",
				"unknown",
				runtime.Version(),
				"unknown",
				runtime.GOOS,
				runtime.GOARCH,
				"GPL V2.0",
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVersionString(tt.args.componentName); got != tt.want {
				t.Errorf("NewVersionString() = %v, want %v", got, tt.want)
			}
		})
	}
}
