package scaffold

import (
	"testing"
)

// MockGenerator is a stub for testing the Generator interface
type MockGenerator struct{}

func (m *MockGenerator) Generate(cfg Config) ([]File, error) {
	return []File{
		{Path: "README.md", Content: "# " + cfg.ProjectName},
	}, nil
}

func TestGeneratorInterface(t *testing.T) {
	// Verify MockGenerator satisfies Generator interface
	var _ Generator = (*MockGenerator)(nil)

	tests := []struct {
		name    string
		cfg     Config
		want    []File
		wantErr bool
	}{
		{
			name: "generates files",
			cfg:  Config{ProjectName: "test-project"},
			want: []File{
				{Path: "README.md", Content: "# test-project"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &MockGenerator{}
			got, err := g.Generate(tt.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("Generate() got %d files, want %d", len(got), len(tt.want))
			}
			if got[0].Path != tt.want[0].Path {
				t.Errorf("Generate() path = %v, want %v", got[0].Path, tt.want[0].Path)
			}
			if got[0].Content != tt.want[0].Content {
				t.Errorf("Generate() content = %v, want %v", got[0].Content, tt.want[0].Content)
			}
		})
	}
}
