package propi

import (
	"fmt"
	"regexp"

	"github.com/playwright-community/playwright-go"
)

func GetPhoneNumber(card playwright.Locator, Nth int) string {
	contactNumberLocator := card.Locator("ul > li").Nth(Nth).Locator("button")
	err := contactNumberLocator.WaitFor(
		playwright.LocatorWaitForOptions{State: playwright.WaitForSelectorStateAttached, Timeout: playwright.Float(1_000)},
	)
	if err != nil {
		return ""
	}

	onClick, err := contactNumberLocator.GetAttribute("onclick")
	if err != nil {
		fmt.Println("ERROR - could not get onclick attribute:", err)
		return ""
	}

	re := regexp.MustCompile(`waMsg\('([^']*)'`)
	matches := re.FindStringSubmatch(onClick)
	if len(matches) < 2 {
		fmt.Printf("could not find %d Nth parameter", Nth)
		return ""
	}

	return matches[1]
}
