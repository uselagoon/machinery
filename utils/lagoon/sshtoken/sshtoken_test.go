package sshtoken

import "testing"

func Test_tokenValidOrExpired(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1 - expired token",
			args: args{
				token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJhdWQiOiJ0ZXN0IiwiZXhwIjoxMDAwNjAsImlhdCI6MTAwMDAwLCJpc3MiOiJ0ZXN0Iiwic3ViIjoidGVzdC1zdWJqZWN0In0.h1xd3L4sq7VHe3xN6794-jIq-XwuHPBhnKOoZhPqBIA",
			},
			want: true,
		},
		{
			name: "test2 - valid long lived token",
			args: args{
				token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJhdWQiOiJ0ZXN0IiwiZXhwIjozMjQxODQ0MzIyLCJpYXQiOjE2NjUwNDQzMjIsImlzcyI6InRlc3QiLCJzdWIiOiJ0ZXN0LXN1YmplY3QifQ.-PBIVMxQUYyppodPobJ0ufqHvqJ8boZCuJ5mDtZPC1Q",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tokenValidOrExpired(tt.args.token); got != tt.want {
				t.Errorf("tokenValidOrExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}
