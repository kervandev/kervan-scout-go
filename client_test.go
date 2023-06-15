package kervanscout_test

import (
	"testing"

	kervanscout "github.com/kervandev/kervan-scout-go"
)

func TestSendIssue(t *testing.T) {
	client := kervanscout.New(&kervanscout.Config{
		Host:         "http://localhost:3002",
		ProjectToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiYjE3MGNkMjctNGE0Zi00ZTgwLWFlZTQtYjZkZWI5ODNlMDZkIiwiY3JlYXRlZF9hdCI6InNlY29uZHM6MTY4NTI4Nzk4MyAgbmFub3M6NDU0MzYzODAiLCJpZCI6IjNlYWI0YjM3LTE0MjctNDk3Ni1iMWUyLTMxYjBhMmNmNzZmMyJ9.fzR49XZmqG4G_TzfwIHTFuEtDDXba_4ZTuk1__sNdtI",
	})

	client.SendIssue("erenkeeee", "KKEKKW")
}
