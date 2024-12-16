package e2e

import (
	"testing"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestE2ERounds(t *testing.T) {
	t.Run("Should successfully submit answer", func(t *testing.T) {
		t.Cleanup(ResetBrowserContexts)
		hostPlayerPage := pages[0]
		otherPlayerPage := pages[1]

		err := startGame(hostPlayerPage, otherPlayerPage)
		require.NoError(t, err)

		err = hostPlayerPage.GetByRole("button", playwright.PageGetByRoleOptions{Name: "Close"}).Click()
		require.NoError(t, err)

		err = hostPlayerPage.Locator("#submit_answer_form").
			Locator(`input[name="answer"]`).
			Fill("this is a test answer")
		require.NoError(t, err)

		err = hostPlayerPage.GetByRole("button", playwright.PageGetByRoleOptions{Name: "Submit Answer"}).Click()
		require.NoError(t, err)
	})

	t.Run("Should successfully submit vote for another player", func(t *testing.T) {
		t.Cleanup(ResetBrowserContexts)
		hostPlayerPage := pages[0]
		otherPlayerPage := pages[1]

		err := startGame(hostPlayerPage, otherPlayerPage)
		require.NoError(t, err)

		err = hostPlayerPage.GetByRole("button", playwright.PageGetByRoleOptions{Name: "Close"}).Click()
		require.NoError(t, err)

		err = hostPlayerPage.Locator("#submit_answer_form").
			Locator(`input[name="answer"]`).
			Fill("this is a test answer")
		require.NoError(t, err)

		err = hostPlayerPage.GetByRole("button", playwright.PageGetByRoleOptions{Name: "Submit Answer"}).Click()
		require.NoError(t, err)

		err = hostPlayerPage.GetByRole("button", playwright.PageGetByRoleOptions{Name: "Not Ready"}).Click()
		require.NoError(t, err)

		err = otherPlayerPage.GetByRole("button", playwright.PageGetByRoleOptions{Name: "Close"}).Click()
		require.NoError(t, err)

		err = otherPlayerPage.Locator("#submit_answer_form").
			Locator(`input[name="answer"]`).
			Fill("this is a another answer")
		require.NoError(t, err)

		err = otherPlayerPage.GetByRole("button", playwright.PageGetByRoleOptions{Name: "Submit Answer"}).Click()
		require.NoError(t, err)

		err = otherPlayerPage.GetByRole("button", playwright.PageGetByRoleOptions{Name: "Not Ready"}).Click()
		require.NoError(t, err)

		votesText := otherPlayerPage.GetByText("Votes:")
		expect.Locator(votesText).ToBeVisible()

		votesText = hostPlayerPage.GetByText("Votes:")
		expect.Locator(votesText).ToBeVisible()

		err = hostPlayerPage.GetByRole("button", playwright.PageGetByRoleOptions{Name: "Submit Vote"}).Click()
		require.NoError(t, err)

		votesText = hostPlayerPage.GetByText("Votes: 1")
		expect.Locator(votesText).ToBeVisible()

		err = hostPlayerPage.GetByRole("button", playwright.PageGetByRoleOptions{Name: "Not Ready"}).Click()
		require.NoError(t, err)

		b, err := hostPlayerPage.GetByRole("button", playwright.PageGetByRoleOptions{Name: "Ready"}).IsVisible()
		require.NoError(t, err)
		assert.True(t, b)

		votedFor := hostPlayerPage.GetByText("You all voted for")
		expect.Locator(votedFor).ToBeVisible()
	})
}
