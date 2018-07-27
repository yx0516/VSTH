package json

import (
	"encoding/json"
	"strings"

	"gpc/util/deep"
)

//-----------------------------------------------------------------------------------------------------------//

func Copy(src, dst interface{}) error {
	if jsonByte, err := json.Marshal(src); err == nil {
		d := json.NewDecoder(strings.NewReader(string(jsonByte)))
		d.UseNumber()
		return d.Decode(dst)
	} else {
		return err
	}
}

func CopyValue(old interface{}) (interface{}, error) {
	return deep.Copy(old, Copy)
}

//-----------------------------------------------------------------------------------------------------------//
