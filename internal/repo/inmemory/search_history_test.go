package inmemory_test

import (
	"testing"

	"github.com/serdyanuk/cicd-jenkins/internal/repo/inmemory"
)

func Test_SearchHistoryRepo(t *testing.T) {
	repo := inmemory.NewSearchHistoryRepo()

	testData := "test"
	t.Run("Should add", func(t *testing.T) {
		repo.Add(testData)
		got := repo.List()
		if got[0] != "testData" {
			t.Errorf("expected: %s, got: %s", testData, got[0])
		}
	})
}
