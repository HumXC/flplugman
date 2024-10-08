package nfo

import (
	"bufio"
	"bytes"
	"reflect"
	"strconv"
	"strings"
)

type Plugin struct {
	FileVersion int    `nfo:"fileversion"`
	Tip         string `nfo:"tip"`
	Color       string `nfo:"color"`
	Bitmap      string `nfo:"Bitmap"`
	IconIndex   int    `nfo:"iconindex"`
	PS          PS     `nfo:"ps"`
}
type PS struct {
	Name           string       `nfo:"name"`
	PresetFilename string       `nfo:"presetfilename"`
	Files          int          `nfo:"files"`
	File           []PluginFile `nfo:"file"`
}

type PluginFile struct {
	Name       string `nfo:"name"`
	Filename   string `nfo:"filename"`
	PlugClass  int    `nfo:"plugclass"`
	Type       int    `nfo:"type"`
	BitSize    int    `nfo:"bitsize"`
	Arch       string `nfo:"arch"`
	Magic      int    `nfo:"magic"`
	GUID       string `nfo:"guid"`
	Date       int64  `nfo:"date"`
	Size       int    `nfo:"size"`
	ScanFlags  int    `nfo:"scanflags"`
	Vendorname string `nfo:"vendorname"`
	Category   string `nfo:"category"`
}

func unmarshal(m map[string]any, val reflect.Value) error {
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("nfo")
		if tag == "" {
			continue
		}
		v := m[tag]
		f := val.Field(i)
		switch v := v.(type) {
		case map[string]any:
			if err := unmarshal(v, f); err != nil {
				return err
			}
		case []map[string]any:
			length := len(v)
			slice := reflect.MakeSlice(f.Type(), length, length)
			for j := 0; j < length; j++ {
				err := unmarshal(v[j], slice.Index(j))
				if err != nil {
					return err
				}
			}
			f.Set(slice)
		case string:
			switch f.Kind() {
			case reflect.String: // string
				f.SetString(v)
			case reflect.Int: // int
				i, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				f.SetInt(int64(i))
			case reflect.Int64: //int64
				i, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				f.SetInt(int64(i))
			}
		}
	}
	return nil
}
func Unmarshal(nfo []byte) (Plugin, error) {
	var p Plugin
	m := make(map[string]any)
	scanner := bufio.NewScanner(bytes.NewReader(nfo))
	for scanner.Scan() {
		line := scanner.Text()
		eqIndex := strings.Index(line, "=")
		keys := line[:eqIndex]
		value := line[eqIndex+1:]
		var parent map[string]any = m
		ks := strings.Split(keys, "_")
		index := -1
		indexKey := ""
		if num, err := strconv.Atoi(ks[len(ks)-1]); err == nil {
			index = num
			indexKey = ks[len(ks)-3]
		}
		for i, key := range ks {
			if i == len(ks)-1 {
				parent[key] = value
				break
			}
			if index != -1 && key == indexKey {
				_, ok := parent[key]
				if !ok {
					ps := m["ps"].(map[string]any)
					lengthStr := ps["files"].(string)
					length, _ := strconv.Atoi(lengthStr)
					parent[key] = make([]map[string]any, length)
				}
				arr := parent[key].([]map[string]any)
				if arr[index] == nil {
					arr[index] = make(map[string]any)
				}
				arr[index][ks[i+1]] = value
				break
			}
			_, ok := parent[key]
			if !ok {
				parent[key] = make(map[string]any)
			}
			parent = parent[key].(map[string]any)
		}
	}
	return p, unmarshal(m, reflect.ValueOf(&p).Elem())
}
func marshal(prefix, suffix string, val reflect.Value, w *bytes.Buffer) {
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("nfo")
		if tag == "" {
			continue
		}
		key := prefix + tag + suffix
		f := val.Field(i)
		value := ""
		switch f.Kind() {
		case reflect.String: // string
			value = f.String()
		case reflect.Int, reflect.Int64: // int
			num := int(f.Int())
			if num == 0 {
				break
			}
			value = strconv.Itoa(num)
		case reflect.Struct:
			marshal(key+"_", "", f, w)
			continue
		case reflect.Slice:
			for j := 0; j < f.Len(); j++ {
				marshal(key+"_", "_"+strconv.Itoa(j), f.Index(j), w)
			}
			continue
		}
		if value == "" {
			continue
		}
		w.WriteString(key + "=" + value + "\n")
	}

}
func Marshal(p Plugin) []byte {
	val := reflect.ValueOf(p)
	result := bytes.NewBuffer(nil)
	marshal("", "", val, result)
	return result.Bytes()
}
