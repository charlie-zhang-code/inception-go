package field

import (
	"fmt"
	"github.com/jinzhu/copier"
	"reflect"
)

func Lists(src interface{}, dstPtr interface{}) error {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dstPtr)

	// Check if the destination is a pointer and can be dereferenced
	if dstVal.Kind() != reflect.Ptr || dstVal.IsNil() {
		return fmt.Errorf("destination must be a non-nil pointer")
	}
	dstVal = dstVal.Elem()

	// Ensure the source and destination are slices
	if srcVal.Kind() != reflect.Slice || dstVal.Kind() != reflect.Slice {
		return fmt.Errorf("source and destination must be slices")
	}

	// Iterate over the source slice and convert each item to the destination type
	for i := 0; i < srcVal.Len(); i++ {
		srcItem := srcVal.Index(i).Interface()
		newDstItem := reflect.New(dstVal.Type().Elem()).Interface()

		// Use copier.Copy for field copying
		if err := copier.Copy(newDstItem, srcItem); err != nil {
			return err
		}

		// Handle special fields like time fields if necessary
		if err := TimeFields(srcItem, newDstItem); err != nil {
			return err
		}

		// Append the new item to the destination slice
		dstVal.Set(reflect.Append(dstVal, reflect.ValueOf(newDstItem).Elem()))
	}

	return nil
}
