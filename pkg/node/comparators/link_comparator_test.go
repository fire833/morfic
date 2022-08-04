/*
*	Copyright (C) 2022  Kendall Tauser
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

package comparator

import (
	"testing"

	api "github.com/fire833/morfic/pkg/ipcapi/v1alpha1"
)

func TestCompareLink(t *testing.T) {
	type args struct {
		link1 *api.Link
		link2 *api.Link
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareLink(tt.args.link1, tt.args.link2); got != tt.want {
				t.Errorf("CompareLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareLinkAttributes(t *testing.T) {
	type args struct {
		attr1 *api.LinkAttributes
		attr2 *api.LinkAttributes
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareLinkAttributes(tt.args.attr1, tt.args.attr2); got != tt.want {
				t.Errorf("CompareLinkAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}
