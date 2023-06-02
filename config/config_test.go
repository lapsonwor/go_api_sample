package config

import (
	"reflect"
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		args    args
		wantNot *Configuration
	}{
		{
			args:    args{name: "config.json"},
			wantNot: &Configuration{},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := LoadConfiguration(tt.args.name)
			if reflect.DeepEqual(got, tt.wantNot) {
				t.Errorf("LoadConfiguration() = %+v, wantNot %+v", got, tt.wantNot)
			}
		})
	}
}
