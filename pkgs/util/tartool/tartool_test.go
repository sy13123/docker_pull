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
				inputDirPath := "/home/git_work/go_pull_http/tmp_controller_v1.2.0"
				TarGz( targetFilePath,  inputDirPath )
			  }()
		})
	}
}
