package tartool

import "testing"

func TestTartool(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			func() {
				targetFilePath := "/tmp/test.tar.gz"
				inputDirPath := "/tmp/test"
				TarGz( targetFilePath,  inputDirPath )
			  }()
		})
	}
}
