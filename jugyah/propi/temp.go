package propi

import (
	"fmt"
	"shared"

	"github.com/playwright-community/playwright-go"
)

func Temp(page playwright.Page, fileName string) {
	page.Goto("https://www.propi.in/real-estate-consultants/3-10")

	shared.SaveContextToFile(page, fileName)

	page.Goto("https://www.propi.in/real-estate-consultants/4-10")

}

func IsLoggedIn(page playwright.Page) (bool, error) {
	loggedInElement := page.Locator(`a#UC_HeadingElementsNew_lnkLogin > span:text("Log in")`)
	if err := loggedInElement.WaitFor(
		playwright.LocatorWaitForOptions{State: playwright.WaitForSelectorStateAttached, Timeout: playwright.Float(2_000)},
	); err != nil {
		return true, nil
	}

	exists, err := loggedInElement.IsVisible()
	if err != nil {
		return true, fmt.Errorf("could not check if logged in: %v", err)
	}

	return !exists, nil
}
