package shared

import (
	"github.com/playwright-community/playwright-go"
)

func SaveContextToFile(page playwright.Page, fileName string) (*playwright.StorageState, error) {
	err := FileExtensionNotAllowed(fileName)
	if err != nil {
		return nil, err
	}

	context := page.Context()

	state, err := context.StorageState(fileName + ".json")
	if err != nil {
		return nil, err
	}

	return state, nil
}

func LoadContextFromFile(page playwright.Page, fileName string) (playwright.Page, error) {
	err := FileExtensionNotAllowed(fileName)
	if err != nil {
		return nil, err
	}

	var state playwright.StorageState

	err = ReadJsonFile(fileName, &state)
	if err != nil {
		return nil, err
	}

	optionalCookies := make([]playwright.OptionalCookie, len(state.Cookies))
	for i, cookie := range state.Cookies {

		optionalCookies[i] = playwright.OptionalCookie{
			Name:     cookie.Name,
			Value:    cookie.Value,
			Domain:   &cookie.Domain,
			Path:     &cookie.Path,
			Expires:  &cookie.Expires,
			HttpOnly: &cookie.HttpOnly,
			Secure:   &cookie.Secure,
			SameSite: cookie.SameSite,
		}
	}

	optionalState := &playwright.OptionalStorageState{
		Cookies: optionalCookies,
		Origins: state.Origins,
	}

	context := page.Context()
	browser := context.Browser()
	context, err = browser.NewContext(playwright.BrowserNewContextOptions{
		StorageState: optionalState,
	})
	if err != nil {
		return nil, err
	}

	page.Close()

	page, err = context.NewPage()
	if err != nil {
		return nil, err
	}

	return page, nil
}
