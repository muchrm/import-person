package personsql

import (
	"reflect"
	"testing"
	"time"
)

func Test_converDate(t *testing.T) {
	type args struct {
		s string
	}
	date, _ := time.Parse(time.RFC3339, "1992-06-15T00:00:00Z")
	date2, _ := time.Parse(time.RFC3339, "2560-06-15T00:00:00Z")
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name:    "case #1",
			args:    args{"1992-06-15"},
			want:    date,
			wantErr: false,
		},
		{
			name:    "case #2",
			args:    args{"2560-06-15"},
			want:    date2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := converDate(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("converDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("converDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
