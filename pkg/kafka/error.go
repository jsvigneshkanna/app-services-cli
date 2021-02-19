package kafka

import (
	"errors"
	"fmt"

	"github.com/bf2fc6cc711aee1a0c2a/cli/internal/localizer"
)

func init() {
	localizer.LoadMessageFiles("kafka")
}

type Error struct {
	Err error
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func ErrorNotFound(id string) *Error {
	return &Error{
		Err: errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "kafka.common.error.notFoundErrorById",
			TemplateData: map[string]interface{}{
				"ID": id,
			},
		})),
	}
}