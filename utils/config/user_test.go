package config

import (
	"reflect"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

func TestConfig_GetUser(t *testing.T) {
	type args struct {
		user string
	}
	tests := []struct {
		name        string
		description string
		c           Config
		args        args
		want        *User
		wantErr     bool
	}{
		{
			name:        "test1",
			description: "verify that the user is returned from the config",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				user: "user1",
			},
			want: &User{
				Name: "user1",
				UserConfig: UserConfig{
					Name:   "user1",
					SSHKey: "/home/user1/.ssh/id_ed25519",
					Grant: &oauth2.Token{
						AccessToken:  "blah",
						TokenType:    "grant",
						RefreshToken: "blah",
						Expiry:       time.Time{},
					},
				},
			},
		},
		{
			name:        "test2",
			description: "verify that an error is returned if the user does not exist in the config",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				user: "user3",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_NewUser(t *testing.T) {
	type args struct {
		user UserConfig
	}
	tests := []struct {
		name        string
		description string
		c           Config
		args        args
		wantErr     bool
	}{
		{
			name:        "test1",
			description: "verify that the created user is returned from the config",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				user: UserConfig{
					Name:   "user3",
					SSHKey: "/home/user3/.ssh/id_ed25510",
					Grant: &oauth2.Token{
						AccessToken:  "blah",
						TokenType:    "grant",
						RefreshToken: "blah",
						Expiry:       time.Time{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.NewUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Config.NewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr != true {
				got, err := tt.c.GetUser(tt.args.user.Name)
				if (err != nil) != tt.wantErr {
					t.Errorf("Config.NewUser() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got.UserConfig, tt.args.user) {
					t.Errorf("Config.NewUser() = %v, want %v", got.UserConfig, tt.args.user)
				}
			}
		})
	}
}

func TestConfig_DeleteUser(t *testing.T) {
	type args struct {
		user string
	}
	tests := []struct {
		name        string
		description string
		c           Config
		args        args
		wantErr     bool
	}{
		{
			name:        "test1",
			description: "verify that the deleted user is removed from the config",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				user: "usera",
			},
		},
		{
			name:        "test2",
			description: "verify that the deleted user returns an error if the defined user is attached to a context",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				user: "user1",
			},
			wantErr: true,
		},
		{
			name:        "test3",
			description: "verify that the deleted user returns an error if the defined user does not exist",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				user: "user3",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.DeleteUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Config.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr != true {
				_, err := tt.c.GetUser(tt.args.user)
				if err == nil {
					t.Errorf("Config.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
		})
	}
}

func TestConfig_UpdateUser(t *testing.T) {
	type args struct {
		user UserConfig
	}
	tests := []struct {
		name        string
		description string
		c           Config
		args        args
		wantErr     bool
	}{
		{
			name:        "test1",
			description: "verify that the updated user is returned from the config",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				user: UserConfig{
					Name:   "user2",
					SSHKey: "/home/user2/.ssh/id_ed25511",
					Grant: &oauth2.Token{
						AccessToken:  "blah2",
						TokenType:    "grant",
						RefreshToken: "blah2",
						Expiry:       time.Time{},
					},
				},
			},
		},
		{
			name:        "test2",
			description: "verify that the updated user returns an error if the user is not in the config",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				user: UserConfig{
					Name:   "user4",
					SSHKey: "/home/user4/.ssh/id_ed25511",
					Grant: &oauth2.Token{
						AccessToken:  "blah",
						TokenType:    "grant",
						RefreshToken: "blah",
						Expiry:       time.Time{},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.UpdateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Config.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr != true {
				got, err := tt.c.GetUser(tt.args.user.Name)
				if (err != nil) != tt.wantErr {
					t.Errorf("Config.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got.UserConfig, tt.args.user) {
					t.Errorf("Config.UpdateUser() = %v, want %v", got.UserConfig, tt.args.user)
				}
			}
		})
	}
}
