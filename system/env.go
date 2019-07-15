package system

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	defaultSep = "_"
	envRef     *envReflect
)

func init() {
	if envRef == nil {
		// load .env file
		content, err := GetFileContent(CurrentPath() + "/.env")
		if err != nil {
			panic(err)
		}
		spilt := strings.Split(content, "\r\n")
		for _, envItem := range spilt {
			if len(envItem) > 0 {
				itemSpilt := strings.Split(envItem, "=")
				err := os.Setenv("ENVREFLECT_"+itemSpilt[0], itemSpilt[1])
				if err != nil {
					panic(err)
				}
			}
		}
		// get env
		envRef = new(envReflect)
		err = build(envRef)
		if err != nil {
			panic(err)
		}
	}
}

func upper(v string) string {
	return strings.ToUpper(v)
}

func parseBool(v string) (bool, error) {
	if v == "" {
		return false, nil
	}
	return strconv.ParseBool(v)
}

func parse(key string, f reflect.Value, sf reflect.StructField) error {
	df := sf.Tag.Get("default")
	isRequire, err := parseBool(sf.Tag.Get("require"))
	if err != nil {
		return fmt.Errorf("the value of %s is not a valid `member` of bool ，only "+
			"[1 0 t f T F true false TRUE FALSE True False] are supported", key)
	}
	ev, exist := os.LookupEnv(key)
	Dump(key)
	Dump(ev)
	if !exist && isRequire {
		return fmt.Errorf("%s is required, but has not been set", key)
	}
	if !exist && df != "" {
		ev = df
	}
	// log.Print("ev:", ev)
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
		sep := ";"
		s, exist := sf.Tag.Lookup("slice_sep")
		if exist && s != "" {
			sep = s
		}
		vals := strings.Split(ev, sep)
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

func combine(p, n string, sep string, ok bool) string {
	if p == "" {
		return n
	}
	if !ok {
		return p + defaultSep + n
	}
	return p + sep + n
}

func fill(pf string, ind reflect.Value) error {
	for i := 0; i < ind.NumField(); i++ {
		f := ind.Type().Field(i)
		name := f.Name
		envName, exist := f.Tag.Lookup("env")
		if exist {
			name = envName
		}
		s, exist := f.Tag.Lookup("sep")
		p := combine(pf, upper(name), s, exist)
		err := parse(p, ind.Field(i), f)
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
	dataKey := upper(ind.Type().Name())
	err := fill(dataKey, ind)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 获取环境数据
 */
func EnvFetch() *envReflect {
	return envRef
}
