package jwt

import "testing"

func TestGenerateAdminToken(t *testing.T) {
	type args struct {
		secret        string
		audience      string
		subject       string
		issuer        string
		issued        int64
		expireSeconds int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test1 - 60 second token",
			args: args{
				secret:        "test-secret",
				audience:      "test",
				subject:       "test-subject",
				issuer:        "test",
				issued:        100000,
				expireSeconds: 60,
			},
			want: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJhdWQiOiJ0ZXN0IiwiZXhwIjoxMDAwNjAsImlhdCI6MTAwMDAwLCJpc3MiOiJ0ZXN0Iiwic3ViIjoidGVzdC1zdWJqZWN0In0.h1xd3L4sq7VHe3xN6794-jIq-XwuHPBhnKOoZhPqBIA",
		},
		{
			name: "test1 - 50 year token",
			args: args{
				secret:        "test-secret",
				audience:      "test",
				subject:       "test-subject",
				issuer:        "test",
				issued:        1665044322,
				expireSeconds: 1576800000,
			},
			want: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4iLCJhdWQiOiJ0ZXN0IiwiZXhwIjozMjQxODQ0MzIyLCJpYXQiOjE2NjUwNDQzMjIsImlzcyI6InRlc3QiLCJzdWIiOiJ0ZXN0LXN1YmplY3QifQ.-PBIVMxQUYyppodPobJ0ufqHvqJ8boZCuJ5mDtZPC1Q",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateAdminToken(tt.args.secret, tt.args.audience, tt.args.subject, tt.args.issuer, tt.args.issued, tt.args.expireSeconds)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateAdminToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateAdminToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
