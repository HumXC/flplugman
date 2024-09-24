package nfo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	nfo1, err := os.ReadFile("../test/A1StereoControl.nfo")
	if err != nil {
		t.Fatal(err)
	}
	nfo2, err := os.ReadFile("../test/OTT.nfo")
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
	got1, err := Unmarshal(nfo1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, got1, want1)
	got2, err := Unmarshal(nfo2)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, got2, want2)
}
func BenchmarkUnmarshal(b *testing.B) {
	nfo1, err := os.ReadFile("../test/A1StereoControl.nfo")
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		Unmarshal(nfo1)
	}
}

func TestMarshal(t *testing.T) {
	want1 := `fileversion=2
ps_name=A1StereoControl
ps_presetfilename=%FLPluginDBPath%\Installed\Effects\VST\A1StereoControl.fst
ps_files=1
ps_file_name_0=A1StereoControl
ps_file_filename_0=D:\DAW\VST\A1StereoControl-x64.dll
ps_file_plugclass_0=0
ps_file_type_0=1
ps_file_bitsize_0=64
ps_file_arch_0=Intel
ps_file_magic_0=1093751651
ps_file_guid_0={56535441-3153-6361-3173-746572656F63}
ps_file_date_0=4676012264064203365
ps_file_size_0=5290496
ps_file_scanflags_0=1
ps_file_vendorname_0=A1AUDIO.de
ps_file_category_0=Effect
`
	want2 := `fileversion=2
ps_name=OTT
ps_presetfilename=%FLPluginDBPath%\Installed\Effects\VST3\OTT.fst
ps_files=1
ps_file_name_0=OTT
ps_file_filename_0=C:\Program Files\Common Files\VST3\OTT.vst3
ps_file_plugclass_0=7
ps_file_type_0=1
ps_file_bitsize_0=64
ps_file_arch_0=Intel
ps_file_guid_0={56534558-6654-546F-7474-000000000000}
ps_file_date_0=4676374293729955603
ps_file_size_0=3379200
ps_file_scanflags_0=1
ps_file_vendorname_0=Xfer Records
ps_file_category_0=Fx|Dynamics
`
	p1 := Plugin{
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
	p2 := Plugin{
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
	got1 := Marshal(p1)
	assert.Equal(t, want1, string(got1))
	got2 := Marshal(p2)
	assert.Equal(t, want2, string(got2))
}
func BenchmarkMarshal(b *testing.B) {
	p1 := Plugin{
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
	for i := 0; i < b.N; i++ {
		Marshal(p1)
	}
}
