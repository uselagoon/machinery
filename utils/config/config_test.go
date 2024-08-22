package config

import (
	"os"
	"reflect"
	"testing"
	"time"

	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"
)

func readConfig(file string) Config {
	rawYAML, err := os.ReadFile(file)
	c := &Config{}
	if err != nil {
		return *c
	}
	err = yaml.Unmarshal(rawYAML, c)
	if err != nil {
		return *c
	}
	return *c
}

func TestConfig_writeConfig(t *testing.T) {
	type args struct {
		file string
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
			description: "verify that writing a config file writes as expected",
			args: args{
				file: "test-data/config.1.yaml",
			},
			c: readConfig("test-data/seed-config.yaml"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.writeConfig(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("Config.writeConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
			seed := readConfig("test-data/config.1.yaml")
			if !reflect.DeepEqual(seed, tt.c) {
				t.Errorf("Config.GetContext() = %v, want %v", seed, tt.c)
			}
		})
	}
}

func TestConfig_readConfig(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name        string
		description string
		args        args
		wantErr     bool
	}{
		{
			name:        "test1",
			description: "verify that reading a config file reads as expected",
			args: args{
				file: "test-data/config.2.yaml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{}
			if err := c.readConfig(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("Config.readConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
			seed := readConfig("test-data/config.2.yaml")
			if !reflect.DeepEqual(&seed, c) {
				t.Errorf("Config.GetContext() = %v, want %v", &seed, c)
			}
		})
	}
}

func TestConfig_checkUserExists(t *testing.T) {
	type args struct {
		user string
	}
	tests := []struct {
		name        string
		description string
		c           Config
		args        args
		want        int
		want1       bool
	}{
		{

			name:        "test1",
			description: "verify that a user exists within the config",
			args: args{
				user: "user1",
			},
			c:     readConfig("test-data/seed-config.yaml"),
			want:  0,
			want1: true,
		},
		{

			name:        "test2",
			description: "verify that a user exists within the config",
			args: args{
				user: "user2",
			},
			c:     readConfig("test-data/seed-config.yaml"),
			want:  1,
			want1: true,
		},
		{

			name:        "test3",
			description: "verify that a user does not exist within the config",
			args: args{
				user: "user3",
			},
			c:     readConfig("test-data/seed-config.yaml"),
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.c.checkUserExists(tt.args.user)
			if got != tt.want {
				t.Errorf("Config.checkUserExists() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Config.checkUserExists() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestConfig_checkContextExists(t *testing.T) {
	type args struct {
		context string
	}
	tests := []struct {
		name        string
		description string
		c           Config
		args        args
		want        int
		want1       bool
	}{
		{

			name:        "test1",
			description: "verify that a context exists within the config",
			args: args{
				context: "lagoon",
			},
			c:     readConfig("test-data/seed-config.yaml"),
			want:  0,
			want1: true,
		},
		{

			name:        "test2",
			description: "verify that a context exists within the config",
			args: args{
				context: "another-lagoon",
			},
			c:     readConfig("test-data/seed-config.yaml"),
			want:  1,
			want1: true,
		},
		{

			name:        "test3",
			description: "verify that a context does not exist within the config",
			args: args{
				context: "someother-lagoon",
			},
			c:     readConfig("test-data/seed-config.yaml"),
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.c.checkContextExists(tt.args.context)
			if got != tt.want {
				t.Errorf("Config.checkContextExists() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Config.checkContextExists() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestConfig_GetUserContext(t *testing.T) {
	type args struct {
		context string
	}
	tests := []struct {
		name        string
		description string
		c           Config
		args        args
		want        *UserConfig
		want1       *ContextConfig
		wantErr     bool
	}{
		{
			name:        "test1",
			description: "verify that the requested context will return the context and associated user from the config",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: "lagoon",
			},
			want: &UserConfig{
				Name:   "user1",
				SSHKey: "/home/user1/.ssh/id_ed25519",
				Grant: &oauth2.Token{
					AccessToken:  "blah",
					TokenType:    "grant",
					RefreshToken: "blah",
					Expiry:       time.Time{},
				},
				PublicKeyIdentities: []string{},
			},
			want1: &ContextConfig{
				Name:                   "lagoon",
				APIHostname:            "https://api.example.com",
				AuthenticationEndpoint: "https://keycloak.example.com",
				TokenHost:              "token.example.com",
				TokenPort:              22,
				SSHHost:                "ssh.example.com",
				SSHPort:                22,
			},
		},
		{
			name:        "test1",
			description: "verify that the requested context will return an error if the requested context does not exist in the config",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: "not-a-lagoon",
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.c.GetUserContext(tt.args.context)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.GetUserContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.GetUserContext() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Config.GetUserContext() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
