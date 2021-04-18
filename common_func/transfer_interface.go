package common_func

import (
	"bytes"
	"encoding/json"
	"reflect"
)

//dst必须传指针
func TransferUserNumber(dst interface{}, src ...interface{}) error {
	for _, data := range src {
		srcBytes := []byte{}
		var err error
		if reflect.TypeOf(data).Kind() != reflect.String {
			srcBytes, err = json.Marshal(data)
			if err != nil {
				return err
			}
		} else {
			srcBytes = []byte(data.(string))
		}

		buffer := bytes.NewBuffer(srcBytes)
		decoder := json.NewDecoder(buffer)
		decoder.UseNumber()
		if err := decoder.Decode(dst); err != nil {
			return err
		}
	}

	return nil
}