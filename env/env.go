package env

import (
	"../php2go"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	envRawData map[string]string
	Data       *envReflect
)

/**
 * 初始化
 */
func init() {
	envRawData = make(map[string]string)
	if Data == nil {
		// load .env file
		content, err := php2go.FileGetContents(".env")
		if err != nil {
			panic(err)
		}
		split := php2go.Explode("\n", content)
		for _, envItem := range split {
			if len(envItem) > 0 {
				itemSpilt := strings.Split(envItem, "=")
				itemKey := itemSpilt[0]
				itemKeyArr := strings.Split(itemKey, "_")
				for k, v := range itemKeyArr {
					itemKeyArr[k] = php2go.Ucfirst(strings.ToLower(v))
				}
				itemKey = php2go.Implode("", itemKeyArr)
				envRawData[itemKey] = itemSpilt[1]
			}
		}
		Data = new(envReflect)
		err = build(Data)
		if err != nil {
			panic(err)
		}
	}
}

func parseBool(v string) (bool, error) {
	if v == "" {
		return false, nil
	}
	return strconv.ParseBool(v)
}

func parse(key string, f reflect.Value, sf reflect.StructField) error {
	df := sf.Tag.Get("default")
	ev := envRawData[key]
	if ev == "" && df != "" {
		ev = df
	}
	switch f.Kind() {
	case reflect.String:
		f.SetString(ev)
	case reflect.Int:
		iv, err := strconv.ParseInt(ev, 10, 32)
		if err != nil {
			return fmt.Errorf("%s:%s", key, err)
		}
		f.SetInt(iv)
	case reflect.Int64:
		if f.Type().String() == "time.Duration" {
			t, err := time.ParseDuration(ev)
			if err != nil {
				return fmt.Errorf("%s:%s", key, err)
			}
			f.Set(reflect.ValueOf(t))
		} else {
			iv, err := strconv.ParseInt(ev, 10, 64)
			if err != nil {
				return fmt.Errorf("%s:%s", key, err)
			}
			f.SetInt(iv)
		}
	case reflect.Uint:
		uiv, err := strconv.ParseUint(ev, 10, 32)
		if err != nil {
			return fmt.Errorf("%s:%s", key, err)
		}
		f.SetUint(uiv)
	case reflect.Uint64:
		uiv, err := strconv.ParseUint(ev, 10, 64)
		if err != nil {
			return fmt.Errorf("%s:%s", key, err)
		}
		f.SetUint(uiv)
	case reflect.Float32:
		f32, err := strconv.ParseFloat(ev, 32)
		if err != nil {
			return fmt.Errorf("%s:%s", key, err)
		}
		f.SetFloat(f32)
	case reflect.Float64:
		f64, err := strconv.ParseFloat(ev, 64)
		if err != nil {
			return fmt.Errorf("%s:%s", key, err)
		}
		f.SetFloat(f64)
	case reflect.Bool:
		b, err := parseBool(ev)
		if err != nil {
			return fmt.Errorf("%s:%s", key, err)
		}
		f.SetBool(b)
	case reflect.Slice:
		vals := strings.Split(ev, ",")
		switch f.Type() {
		case reflect.TypeOf([]string{}):
			f.Set(reflect.ValueOf(vals))
		case reflect.TypeOf([]int{}):
			t := make([]int, len(vals))
			for i, v := range vals {
				val, err := strconv.ParseInt(v, 10, 32)
				if err != nil {
					return fmt.Errorf("%s:%s", key, err)
				}
				t[i] = int(val)
			}
		case reflect.TypeOf([]int64{}):
			t := make([]int64, len(vals))
			for i, v := range vals {
				val, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s:%s", key, err)
				}
				t[i] = val
			}
		case reflect.TypeOf([]uint{}):
			t := make([]uint, len(vals))
			for i, v := range vals {
				val, err := strconv.ParseUint(v, 10, 32)
				if err != nil {
					return fmt.Errorf("%s:%s", key, err)
				}
				t[i] = uint(val)
			}
		case reflect.TypeOf([]uint64{}):
			t := make([]uint64, len(vals))
			for i, v := range vals {
				val, err := strconv.ParseUint(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s:%s", key, err)
				}
				t[i] = val
			}
		case reflect.TypeOf([]float32{}):
			t := make([]float32, len(vals))
			for i, v := range vals {
				val, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s:%s", key, err)
				}
				t[i] = float32(val)
			}
		case reflect.TypeOf([]float64{}):
			t := make([]float64, len(vals))
			for i, v := range vals {
				val, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s:%s", key, err)
				}
				t[i] = val
			}
		case reflect.TypeOf([]bool{}):
			t := make([]bool, len(vals))
			for i, v := range vals {
				val, err := parseBool(v)
				if err != nil {
					return fmt.Errorf("%s:%s", key, err)
				}
				t[i] = val
			}
		}
	}
	return nil
}

func fill(ind reflect.Value) error {
	for i := 0; i < ind.NumField(); i++ {
		f := ind.Type().Field(i)
		key := f.Name
		err := parse(key, ind.Field(i), f)
		if err != nil {
			return err
		}
	}
	return nil
}

func build(v interface{}) error {
	ind := reflect.Indirect(reflect.ValueOf(v))
	if reflect.ValueOf(v).Kind() != reflect.Ptr || ind.Kind() != reflect.Struct {
		return fmt.Errorf("only the pointer to a struct is supported")
	}
	// get env data's key name
	err := fill(ind)
	if err != nil {
		return err
	}
	return nil
}
