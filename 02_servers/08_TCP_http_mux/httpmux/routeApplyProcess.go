package httpmux

import (
	"fmt"
	"io"
)


type applyProcessStruct struct {

}


func NewApplyProcess() handleRoute {
	return &applyProcessStruct{}
}

func (*applyProcessStruct) String() string {
	return fmt.Sprintf("Apply Process handler")
}

func (*applyProcessStruct) Name() string {
		return fmt.Sprintf("ApplyProcess")
}

func (*applyProcessStruct) Body(w io.Writer) error {
	return nil
}