package kervanscout_test

import (
	"errors"
	"testing"

	kervanscout "github.com/kervandev/kervan-scout-go/v2"
)

func TestSendIssue(t *testing.T) {
	client := kervanscout.New(kervanscout.Config{
		ProjectToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0Ijoic2Vjb25kczoxNjg2Nzg5MjcyICBuYW5vczo0OTMyNjQxMDQiLCJpZCI6ImIxNTliNDg4LTgyMTYtNGQwMS05YWM2LTA5Mzg2M2I0YTdkMSJ9.6sNYAmoczxqHoWp_SVEJUdSr_w99SK-ndzfKQWvrUD0",
	})
	defer client.CatchPanicError("admin - panic error")

	var err error = errors.New("normal error")

	if err != nil {
		client.SendIssue(&kervanscout.Issue{
			Title:   "admin - normal error",
			Message: err.Error(),
		})
	}

	panic("panic error")
}
