package variables

import (
	"os"
	"reflect"
	"testing"
)

func TestGetLagoonFeatureFlags(t *testing.T) {
	tests := []struct {
		name string
		set  map[string]string
		want map[string]string
	}{
		{
			name: "test1",
			set: map[string]string{
				"LAGOON_FEATURE_FLAG_ABC": "enabled",
				"LAGOON_FEATURE_FLAG_123": "enabled",
				"OTHERS":                  "path noises",
			},
			want: map[string]string{
				"LAGOON_FEATURE_FLAG_ABC": "enabled",
				"LAGOON_FEATURE_FLAG_123": "enabled",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for e, v := range tt.set {
				os.Setenv(e, v)
			}
			if got := GetLagoonFeatureFlags(); !reflect.DeepEqual(got, tt.want) {
				for e := range tt.set {
					os.Unsetenv(e)
				}
				t.Errorf("GetLagoonFeatureFlags() = %v, want %v", got, tt.want)
			}
			for e := range tt.set {
				os.Unsetenv(e)
			}
		})
	}
}
