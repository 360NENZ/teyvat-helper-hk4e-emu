package asset

import (
	"reflect"
	"testing"
)

func TestNewAbilityNameHash(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want AbilityNameHash
	}{{
		name: "Dafault",
		args: args{name: "Dafault"},
		want: 1178079449,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAbilityNameHash(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAbilityNameHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
