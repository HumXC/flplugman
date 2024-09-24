package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePlugin(t *testing.T) {
	nfo1, err := os.Open("test/A1StereoControl.nfo")
	if err != nil {
		t.Fatal(err)
	}
	nfo2, err := os.Open("test/OTT.nfo")
	if err != nil {
		t.Fatal(err)
	}
	_ = nfo2
	want1 := Plugin{
		Name:           "A1StereoControl",
		PresetFilename: "%FLPluginDBPath%\\Installed\\Effects\\VST\\A1StereoControl.fst",
		Files:          1,
		File: []PluginFile{
			{
				Name:       "A1StereoControl",
				Filename:   "D:\\DAW\\VST\\A1StereoControl-x64.dll",
				PlugClass:  0,
				Type:       1,
				BitSize:    64,
				Arch:       "Intel",
				Magic:      1093751651,
				GUID:       "{56535441-3153-6361-3173-746572656F63}",
				Date:       4676012264064203365,
				Size:       5290496,
				ScanFlags:  1,
				Vendorname: "A1AUDIO.de",
				Category:   "Effect",
			},
		},
	}
	want2 := Plugin{
		Name:           "OTT",
		PresetFilename: "%FLPluginDBPath%\\Installed\\Effects\\VST3\\OTT.fst",
		Files:          1, File: []PluginFile{
			{
				Name:       "OTT",
				Filename:   "C:\\Program Files\\Common Files\\VST3\\OTT.vst3",
				PlugClass:  7,
				Type:       1,
				BitSize:    64,
				Arch:       "Intel",
				GUID:       "{56534558-6654-546F-7474-000000000000}",
				Date:       4676374293729955603,
				Size:       3379200,
				ScanFlags:  1,
				Vendorname: "Xfer Records",
				Category:   "Fx|Dynamics",
			},
		},
	}
	got1, err := ParsePlugin(nfo1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, got1, want1)
	got2, err := ParsePlugin(nfo2)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, got2, want2)
}
func BenchmarkParsePlugin(b *testing.B) {
	data, err := os.ReadFile("test/A1StereoControl.nfo")
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		nfo1 := bytes.NewReader(data)
		ParsePlugin(nfo1)
	}
}
