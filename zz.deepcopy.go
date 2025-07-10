package podinfo

import (
	"bytes"
	"encoding/json"
)

func DeepCopyJSON[T any](src T, dst T) error {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(src); err != nil {
		return err
	}
	if err := json.NewDecoder(buf).Decode(&dst); err != nil {
		return err
	}
	return nil
}

func (in *HttpRuntimeResponse) DeepCopyInto(out *HttpRuntimeResponse) {
	DeepCopyJSON(in, out)
}
