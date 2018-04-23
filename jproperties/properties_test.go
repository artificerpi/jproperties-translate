package jproperties

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_readProps(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantP   *Properties
		wantErr bool
	}{
		{
			"test_read_props",
			args{strings.NewReader("language=zh_CN\r\nname=jproperties\r\n")},
			&Properties{dict: map[string]string{
				"language": "zh_CN",
				"name":     "jproperties"}},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, err := readProps(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("readProps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("readProps() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}

func Test_writeProps(t *testing.T) {
	type args struct {
		p *Properties
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
		wantErr    bool
	}{
		{
			"test_write_props",
			args{
				&Properties{dict: map[string]string{
					"language": "zh_CN"},
					description: "sample properties file"},
			},
			"sample properties file\nlanguage=zh_CN\n",
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			if err := writeProps(writer, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("writeProps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("writeProps() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
