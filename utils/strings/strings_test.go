package strings

import (
	"reflect"
	"testing"
)

func TestContainsString(t *testing.T) {
	type args struct {
		slice []string
		s     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				slice: []string{"string-one", "string-two"},
				s:     "string-one",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsString(tt.args.slice, tt.args.s); got != tt.want {
				t.Errorf("ContainsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveString(t *testing.T) {
	type args struct {
		slice []string
		s     string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{
			name: "test1",
			args: args{
				slice: []string{"string-one", "string-two"},
				s:     "string-one",
			},
			wantResult: []string{"string-two"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := RemoveString(tt.args.slice, tt.args.s); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("RemoveString() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestStripNewLines(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				s: "string-one\n",
			},
			want: "string-one",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripNewLines(tt.args.s); got != tt.want {
				t.Errorf("StripNewLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReturnNonEmptyString(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				value: "string-one",
			},
			want: "string-one",
		},
		{
			name: "test2",
			args: args{
				value: "",
			},
			want: "-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReturnNonEmptyString(tt.args.value); got != tt.want {
				t.Errorf("ReturnNonEmptyString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlicesEqual(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				a: []string{"nginx", "solr", "redis"},
				b: []string{"solr", "nginx", "redis"},
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				a: []string{"nginx", "solr", "redis"},
				b: []string{"solr", "nginx", "redis", "mariadb"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SlicesEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("SlicesEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
