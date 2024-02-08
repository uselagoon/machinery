package config

import (
	"reflect"
	"testing"
)

func TestConfig_GetContext(t *testing.T) {
	type args struct {
		context string
	}
	tests := []struct {
		name        string
		description string
		c           Config
		args        args
		want        *Context
		wantErr     bool
	}{
		{
			name:        "test1",
			description: "verify that the requested context is returned from the config",
			args: args{
				context: "lagoon",
			},
			c: readConfig("test-data/seed-config.yaml"),
			want: &Context{
				Name: "lagoon",
				ContextConfig: ContextConfig{
					Name:                   "lagoon",
					APIHostname:            "https://api.example.com",
					AuthenticationEndpoint: "https://keycloak.example.com",
					TokenHost:              "token.example.com",
					TokenPort:              22,
				},
				User: "user1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetContext(tt.args.context)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.GetContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.GetContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_NewContext(t *testing.T) {
	type args struct {
		context ContextConfig
		user    string
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
			description: "verify that the created context is returned from the config",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: ContextConfig{
					Name:                   "new-lagoon",
					APIHostname:            "https://newapi.example.com",
					AuthenticationEndpoint: "https://newkeycloak.example.com",
					TokenHost:              "newtoken.example.com",
					TokenPort:              22,
				},
				user: "user2",
			},
		},
		{
			name:        "test2",
			description: "verify that the created context returns an error if the defined user does not exist",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: ContextConfig{
					Name:                   "new-lagoon",
					APIHostname:            "https://newapi.example.com",
					AuthenticationEndpoint: "https://newkeycloak.example.com",
					TokenHost:              "newtoken.example.com",
					TokenPort:              22,
				},
				user: "user3",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.NewContext(tt.args.context, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Config.NewContext() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr != true {
				got, err := tt.c.GetContext(tt.args.context.Name)
				if (err != nil) != tt.wantErr {
					t.Errorf("Config.NewContext() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got.ContextConfig, tt.args.context) {
					t.Errorf("Config.NewContext() = %v, want %v", got.ContextConfig, tt.args.context)
				}
			}
		})
	}
}

func TestConfig_DeleteContext(t *testing.T) {
	type args struct {
		context string
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
			description: "verify that the deleted context is removed from the config",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: "lagoon",
			},
		},
		{
			name:        "test2",
			description: "verify that the deleted context returns an error if the defined context does not exist",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: "lagoon2",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.DeleteContext(tt.args.context); (err != nil) != tt.wantErr {
				t.Errorf("Config.DeleteContext() error = %v, wantErr %v", err, tt.wantErr)
			}
			_, err := tt.c.GetUser(tt.args.context)
			if err == nil {
				t.Errorf("Config.DeleteContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestConfig_GetDefaultContext(t *testing.T) {
	tests := []struct {
		name        string
		description string
		c           Config
		want        *Context
		wantErr     bool
	}{
		{
			name:        "test1",
			description: "verify that the default context is returned from the config",
			c:           readConfig("test-data/seed-config.yaml"),
			want: &Context{
				Name: "lagoon",
				ContextConfig: ContextConfig{
					Name:                   "lagoon",
					APIHostname:            "https://api.example.com",
					AuthenticationEndpoint: "https://keycloak.example.com",
					TokenHost:              "token.example.com",
					TokenPort:              22,
				},
				User: "user1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetDefaultContext()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.GetDefaultContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.GetDefaultContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_SetDefaultContext(t *testing.T) {
	type args struct {
		context string
	}
	tests := []struct {
		name        string
		description string
		c           Config
		args        args
		want        *Context
		wantErr     bool
	}{
		{
			name:        "test1",
			description: "verify that the default context is changed in the config",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: "another-lagoon",
			},
			want: &Context{
				Name: "another-lagoon",
				ContextConfig: ContextConfig{
					Name:                   "another-lagoon",
					APIHostname:            "https://api2.example.com",
					AuthenticationEndpoint: "https://keycloak2.example.com",
					TokenHost:              "token2.example.com",
					TokenPort:              22,
				},
				User: "user2",
			},
		},
		{
			name:        "test2",
			description: "verify that the default context returns an error if the requested context is not in the config",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: "fake-lagoon",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.SetDefaultContext(tt.args.context); (err != nil) != tt.wantErr {
				t.Errorf("Config.SetDefaultContext() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr != true {
				got, err := tt.c.GetDefaultContext()
				if (err != nil) != tt.wantErr {
					t.Errorf("Config.GetDefaultContext() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Config.GetDefaultContext() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestConfig_UpdateContext(t *testing.T) {
	type args struct {
		context ContextConfig
		user    string
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
			description: "verify that the updated context is returned from the config",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: ContextConfig{
					Name:                   "lagoon",
					APIHostname:            "https://newapi.example.com",
					AuthenticationEndpoint: "https://newkeycloak.example.com",
					TokenHost:              "newtoken.example.com",
					TokenPort:              22,
				},
				user: "user2",
			},
		},
		{
			name:        "test2",
			description: "verify that the updated context returns an error if the defined context does not exist",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: ContextConfig{
					Name:                   "new-lagoon",
					APIHostname:            "https://newapi.example.com",
					AuthenticationEndpoint: "https://newkeycloak.example.com",
					TokenHost:              "newtoken.example.com",
					TokenPort:              22,
				},
				user: "user3",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.UpdateContext(tt.args.context, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Config.UpdateContext() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr != true {
				got, err := tt.c.GetContext(tt.args.context.Name)
				if (err != nil) != tt.wantErr {
					t.Errorf("Config.UpdateContext() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got.ContextConfig, tt.args.context) {
					t.Errorf("Config.UpdateContext() = %v, want %v", got.ContextConfig, tt.args.context)
				}
			}
		})
	}
}
