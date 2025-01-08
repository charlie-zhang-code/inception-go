package field

import (
	"reflect"
	"time"
)

func TimeFields(src, dst interface{}) error {
	// 获取源和目标类型的反射值
	srcVal := reflect.ValueOf(src).Elem()
	dstVal := reflect.ValueOf(dst).Elem()

	// 遍历目标结构体的所有字段
	for i := 0; i < dstVal.NumField(); i++ {
		dstField := dstVal.Type().Field(i)
		srcField := srcVal.FieldByName(dstField.Name)

		// 检查是否是 time.Time 类型并且在源结构体中存在对应字段
		if srcField.IsValid() && srcField.Type() == reflect.TypeOf(time.Time{}) {
			// 格式化时间并设置到目标字段
			timeStr := ""
			if !srcField.Interface().(time.Time).IsZero() {
				timeStr = srcField.Interface().(time.Time).Format(time.DateTime)
			}
			dstVal.Field(i).SetString(timeStr)
		}
	}

	return nil
}
