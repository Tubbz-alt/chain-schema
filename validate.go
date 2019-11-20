package chain_schema

import (
	"fmt"
	"github.com/golang/protobuf/descriptor"
	//"github.com/uber/prototool"
)

func (e Enum) Validate() error {
	for _, v := range e.Values {
		if v.Value == 0 {
			return nil
		}
	}
	return fmt.Errorf("enum must have zero value")
}

func Test()  {
	x, y := descriptor.ForMessage(&Struct{})
}
