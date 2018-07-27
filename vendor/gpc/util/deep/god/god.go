package god

import (
	"bytes"
	"encoding/gob"

	"gpc/util/deep"
)

//-----------------------------------------------------------------------------------------------------------//

func Copy(src, dst interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func CopyValue(old interface{}) (interface{}, error) {
	return deep.Copy(old, Copy)
}

//-----------------------------------------------------------------------------------------------------------//
