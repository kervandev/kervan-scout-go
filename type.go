package kervanscout

import (
	"fmt"
	"strings"
)

// IssueType is the type of issue
type IssueType string

const (
	// IssueTypeUnknown is the default issue type
	IssueTypeUnknown IssueType = "unknown"
	// IssueTypeExecution is an execution issue
	IssueTypeExecution = "execution"
	// IssueTypeBug is a bug
	IssueTypeBug = "bug"
	// IssueTypeFeature is a feature
	IssueTypeFeature = "feature"
	// IssueTypeEnhancement is an enhancement
	IssueTypeEnhancement = "enhancement"
	// IssueTypeQuestion is a question
	IssueTypeQuestion = "question"
	// IssueTypeHelp is a help issue
	IssueTypeHelp = "help"
)

// String returns the string representation of the issue type
func (it IssueType) String() string {
	return string(it)
}

// ParseIssueType parses the issue type from a string
func ParseIssueType(s string) (IssueType, error) {
	switch strings.ToLower(s) {
	case "bug":
		return IssueTypeBug, nil
	case "execution":
		return IssueTypeExecution, nil
	case "feature":
		return IssueTypeFeature, nil
	case "enhancement":
		return IssueTypeEnhancement, nil
	case "question":
		return IssueTypeQuestion, nil
	case "help":
		return IssueTypeHelp, nil
	case "unknown":
		return IssueTypeUnknown, nil
	default:
		return "", fmt.Errorf("unknown issue type: %s", s)
	}
}
