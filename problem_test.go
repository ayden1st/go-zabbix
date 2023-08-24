package zabbix

import (
	"testing"
)

func TestProblem(t *testing.T) {
	session := GetTestSession(t)
	problemParam := ProblemGetParams{
		Recent:     true,
		SelectTags: "extend",
	}

	problems, err := session.GetProblems(problemParam)
	if err != nil {
		t.Fatalf("Error getting problems: %v", err)
	}

	t.Logf("Validated %d problems", len(problems))
}
