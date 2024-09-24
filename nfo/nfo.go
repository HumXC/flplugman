package nfo

import (
	"bufio"
	"bytes"
	"reflect"
	"strconv"
	"strings"
)

type Plugin struct {
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

func Unmarshal(nfo []byte) (Plugin, error) {
	var p Plugin
	m := make(map[string]string)
	scanner := bufio.NewScanner(bytes.NewReader(nfo))
	for scanner.Scan() {
		line := scanner.Text()
		eqIndex := strings.Index(line, "=")
		key := line[:eqIndex]
		value := line[eqIndex+1:]
		m[key] = value
	}
	val := reflect.ValueOf(&p).Elem()
	typ := reflect.TypeOf(p)

	for i := 0; i < val.NumField(); i++ {
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("nfo")
		v := m["ps_"+tag]
		if val.Field(i).Type() == reflect.TypeOf(make([]PluginFile, 0)) {
			for j := 0; j < p.Files; j++ {
				pf, err := unmarshalPluginFile(m, j)
				if err != nil {
					return p, err
				}
				slice := val.Field(i)
				newSlice := reflect.MakeSlice(slice.Type(), slice.Len()+1, slice.Cap()+1)
				reflect.Copy(newSlice, slice)
				newSlice.Index(slice.Len()).Set(reflect.ValueOf(pf))
				val.Field(i).Set(newSlice)
			}
		} else {
			err := setVal(val.Field(i), v)
			if err != nil {
				return p, err
			}
		}
	}
	return p, nil
}
func Marshal(p Plugin) []byte {
	result := bytes.NewBufferString("fileversion=2\n")

	val := reflect.ValueOf(p)
	typ := reflect.TypeOf(p)
	for i := 0; i < val.NumField(); i++ {
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("nfo")
		if tag != "file" {
			key := "ps_" + tag
			value := getVal(val.Field(i))
			result.WriteString(key + "=" + value + "\n")
		} else {
			for j := 0; j < val.Field(i).Len(); j++ {
				marshalPluginFile(val.Field(i).Index(j), j, result)
			}
		}
	}
	return result.Bytes()
}
func unmarshalPluginFile(m map[string]string, index int) (PluginFile, error) {
	var p PluginFile
	val := reflect.ValueOf(&p).Elem()
	typ := reflect.TypeOf(p)
	for i := 0; i < val.NumField(); i++ {
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("nfo")
		v := m["ps_file_"+tag+"_"+strconv.Itoa(index)]
		err := setVal(val.Field(i), v)
		if err != nil {
			return p, err
		}
	}
	return p, nil
}
func marshalPluginFile(v reflect.Value, index int, w *bytes.Buffer) {
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldType := t.Field(i)
		tag := fieldType.Tag.Get("nfo")
		key := "ps_file_" + tag + "_" + strconv.Itoa(index)
		value := getVal(v.Field(i))
		if tag == "magic" && v.Field(i).Int() == 0 {
			continue
		}
		w.WriteString(key + "=" + value + "\n")
	}
}

func getVal(field reflect.Value) string {
	switch field.Kind() {
	case reflect.String: // string
		return field.String()
	case reflect.Int: // int
		return strconv.Itoa(int(field.Int()))
	case reflect.Int64: //int64
		return strconv.Itoa(int(field.Int()))
	}
	return ""
}
func setVal(field reflect.Value, value string) error {
	if value == "" {
		return nil
	}
	switch field.Kind() {
	case reflect.String: // string
		field.SetString(value)
	case reflect.Int: // int
		i, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		field.SetInt(int64(i))
	case reflect.Int64: //int64
		i, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		field.SetInt(int64(i))
	}
	return nil
}
