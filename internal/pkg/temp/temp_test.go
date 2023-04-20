package temp

import (
	"testing"
)

// TODO need to write tests.

func TestTemplate_GetTemplateByAlias(t *testing.T) {
	type args struct {
		fileName string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			t := Template{}
			got, err := t.GetTemplateByAlias(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetTemplateByAlias() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("GetTemplateByAlias() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTemplate_getTemplateDir(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "test",
			want:    "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			te := Template{}
			got, err := te.getTemplateDir()
			if (err != nil) != tt.wantErr {
				t.Errorf("getTemplateDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getTemplateDir() got = %v, want %v", got, tt.want)
			}
		})
	}
}
