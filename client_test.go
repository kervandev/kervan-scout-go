package kervanscout_test

import (
	"errors"
	"testing"

	kervanscout "github.com/kervandev/kervan-scout-go/v2"
)

func TestSendIssue(t *testing.T) {
	client := kervanscout.New(kervanscout.Config{
		ProjectToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0Ijoic2Vjb25kczoxNjg4MjU0NDkyIG5hbm9zOjc2Njk5NTg3NSIsImlkIjoiYjE1OWI0ODgtODIxNi00ZDAxLTlhYzYtMDkzODYzYjRhN2QxIn0.h11HK74p_GvZhJ0IxXy6qtStpYG7r88i6NU6msjNlic",
	})
	defer client.CatchPanicError("admin - panic error")

	var err error = errors.New("normal error")

	if err != nil {
		client.SendIssue(
			"admin - normal error",
			err.Error(),
		)
	}

	panic("panic error")
}
