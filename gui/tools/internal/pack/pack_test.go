package pack

import (
	"testing"
)

func Test_createScriptFile(t *testing.T) {
	type args struct {
		filePath         string
		downloadFileName string
		version          string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"name 1111",
			args{
				filePath:         "../" + PackDeployFileNamePath,
				downloadFileName: "edgex_v1.5.0_Darwin_x86_64.tar.gz",
				version:          "v1.5.0",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createScriptFile(tt.args.filePath, tt.args.downloadFileName, tt.args.version); (err != nil) != tt.wantErr {
				t.Errorf("createScriptFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
