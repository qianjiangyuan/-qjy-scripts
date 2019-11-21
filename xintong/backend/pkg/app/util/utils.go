package util

import "reflect"

func CopyStruct(src, dst interface{}) {
  sval := reflect.ValueOf(src).Elem()
  dval := reflect.ValueOf(dst).Elem()

  for i := 0; i < sval.NumField(); i++ {
    value := sval.Field(i)
    name := sval.Type().Field(i).Name

    dvalue := dval.FieldByName(name)
    if dvalue.IsValid() == false {
      continue
    }
    dvalue.Set(value)
  }
}
