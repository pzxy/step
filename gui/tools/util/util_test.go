package util

import (
	"testing"
)

func TestCreateTarFile(t *testing.T) {
	type args struct {
		fileName  string
		directory string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"test1",
			args{
				fileName:  "../tedge.tar",
				directory: "../tedge/",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateTarFile(tt.args.fileName, tt.args.directory); (err != nil) != tt.wantErr {
				t.Errorf("CreateTarFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDownloadFile(t *testing.T) {
	//edgex_v1.5.0_Darwin_x86_64.tar.gz
	fileName := "edgex_v1.5.0_Linux_armv7.tar.gz"
	url := "https://github.com/pzxy/tuya-edge-driver-sdk-go/releases/download/1.5.0/edgex_v1.5.0_Darwin_x86_64.tar.gz"
	type args struct {
		fileOutputPath string
		url            string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1",
			args{
				fileOutputPath: fileName,
				url:            url,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DownloadFile(tt.args.fileOutputPath, tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("DownloadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
