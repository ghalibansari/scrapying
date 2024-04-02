package shared

import "github.com/playwright-community/playwright-go"

func InitPlaywrightPage() (*playwright.Playwright, playwright.Browser, playwright.BrowserContext, playwright.Page, error) {
	pw, err := playwright.Run()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		// Headless: playwright.Bool(false),
		// SlowMo:   playwright.Float(2_000),
	})
	if err != nil {
		return nil, nil, nil, nil, err
	}

	context, err := browser.NewContext()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	page, err := context.NewPage()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return pw, browser, context, page, nil
}
