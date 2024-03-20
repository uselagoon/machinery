package config

import (
	"testing"
)

func TestConfig_SetGlobalFeature(t *testing.T) {
	type args struct {
		context string
		prefix  string
		feature string
		state   bool
	}
	tests := []struct {
		name        string
		description string
		c           Config
		args        args
		want        bool
		wantErr     bool
	}{
		{
			name:        "test1",
			description: "verify that updating a feature changes its value",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: "lagoon",
				prefix:  "cli",
				feature: "feature1",
				state:   false,
			},
			want: false,
		},
		{
			name:        "test2",
			description: "verify that updating a feature that doesn't exist sets is value",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: "lagoon",
				prefix:  "cli",
				feature: "feature2",
				state:   true,
			},
			want: true,
		},
		{
			name:        "test3",
			description: "verify that updating a feature with no feature name returns an error",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: "lagoon",
				prefix:  "cli",
				feature: "",
				state:   true,
			},
			wantErr: true,
		},
		{
			name:        "test4",
			description: "verify that updating a feature with no prefix returns an error",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: "lagoon",
				prefix:  "",
				feature: "feature3",
				state:   true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.SetGlobalFeature(tt.args.prefix, tt.args.feature, tt.args.state); (err != nil) != tt.wantErr {
				t.Errorf("Config.SetGlobalFeature() error = %v, wantErr %v", err, tt.wantErr)
			}
			after, _ := tt.c.GetFeature(tt.args.context, tt.args.prefix, tt.args.feature)
			if after != tt.want {
				t.Errorf("Config.SetGlobalFeature() feature '%s': got = %v, want %v", featurePrefix(tt.args.prefix, tt.args.feature), after, tt.want)
			}
		})
	}
}

func TestConfig_SetContextFeature(t *testing.T) {
	type args struct {
		context string
		prefix  string
		feature string
		state   bool
	}
	tests := []struct {
		name        string
		description string
		c           Config
		args        args
		want        bool
		wantErr     bool
	}{
		{
			name:        "test1",
			description: "verify that updating a feature sets its value",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: "lagoon",
				prefix:  "cli",
				feature: "anotherfeature",
				state:   true,
			},
			want: true,
		},
		{
			name:        "test2",
			description: "verify that updating a feature with no feature name returns an error",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: "lagoon",
				prefix:  "cli",
				feature: "",
				state:   true,
			},
			wantErr: true,
		},
		{
			name:        "test3",
			description: "verify that updating a feature with no prefix returns an error",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: "lagoon",
				prefix:  "",
				feature: "anotherfeature",
				state:   true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.SetContextFeature(tt.args.context, tt.args.prefix, tt.args.feature, tt.args.state); (err != nil) != tt.wantErr {
				t.Errorf("Config.SetContextFeature() error = %v, wantErr %v", err, tt.wantErr)
			}
			after, _ := tt.c.GetFeature(tt.args.context, tt.args.prefix, tt.args.feature)
			if after != tt.want {
				t.Errorf("Config.SetContextFeature() feature '%s': got = %v, want %v", featurePrefix(tt.args.prefix, tt.args.feature), after, tt.want)
			}
		})
	}
}

func TestConfig_GetFeature(t *testing.T) {
	type args struct {
		context string
		prefix  string
		feature string
	}
	tests := []struct {
		name        string
		description string
		c           Config
		args        args
		want        bool
		wantErr     bool
	}{
		{
			name:        "test1",
			description: "verify that getting a feature that doesn't exist returns an error",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: "lagoon",
				prefix:  "cli",
				feature: "featuredoesntexist",
			},
			want: false,
		},
		{
			name:        "test2",
			description: "verify that getting a feature that exists on the context and global returns the context value",
			c:           readConfig("test-data/seed-config.2.yaml"),
			args: args{
				context: "lagoon",
				prefix:  "cli",
				feature: "mycontextfeature",
			},
			want: true,
		},
		{
			name:        "test3",
			description: "verify that getting a feature that exists on the global returns the global value",
			c:           readConfig("test-data/seed-config.2.yaml"),
			args: args{
				context: "lagoon",
				prefix:  "cli",
				feature: "myglobalfeature",
			},
			want: true,
		},
		{
			name:        "test4",
			description: "verify that getting a feature for a context that doesn't exist returns an error",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: "differentlagoon",
				prefix:  "cli",
				feature: "featuredoesntexist",
			},
			wantErr: true,
		},
		{
			name:        "test5",
			description: "verify that getting a feature for a feature that doesn't exist returns an error",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: "lagoon",
				prefix:  "cli",
				feature: "",
			},
			wantErr: true,
		},
		{
			name:        "test6",
			description: "verify that getting a feature for a prefix that doesn't exist returns an error",
			c:           readConfig("test-data/seed-config.yaml"),
			args: args{
				context: "lagoon",
				prefix:  "",
				feature: "myglobalfeature",
			},
			wantErr: true,
		},
		{
			name:        "test7",
			description: "verify that getting a feature that exists on the context and global returns the context value",
			c:           readConfig("test-data/seed-config.2.yaml"),
			args: args{
				context: "lagoon",
				prefix:  "cli",
				feature: "mycontextfeature2",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetFeature(tt.args.context, tt.args.prefix, tt.args.feature)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.GetFeature() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr != true {
				if got != tt.want {
					t.Errorf("Config.GetFeature() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
