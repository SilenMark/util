package file

import "testing"

var exist = "isexist.go"
var notExist = "someNotExistStuff.SomeFormatFile"

func TestIsExist(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"fileExist", args{exist}, true},
		{"fileNotExist", args{notExist}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsExist(tt.args.name); got != tt.want {
				t.Errorf("IsExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
