package propi

import (
	"fmt"
	"shared"
	"time"

	"github.com/playwright-community/playwright-go"
)

func Login(page playwright.Page, username string, password string) (playwright.Page, error) {
	fmt.Println("Logging in started")

	if _, err := page.Goto("https://www.propi.in/Login"); err != nil {
		fmt.Println("could not navigate to login page:", err)
		return nil, err
	}

	usernameLocator := page.Locator(`input[name="email"]`)
	if err := usernameLocator.Fill(username); err != nil {
		fmt.Println("could not fill username:", err)
		return nil, err
	}

	passwordLocator := page.Locator(`input[name="password"]`)
	if err := passwordLocator.Fill(password); err != nil {
		return nil, err
	}

	loginButtonLocator := page.Locator(`#buttLogin`)
	time.Sleep(1 * time.Second)
	if err := loginButtonLocator.Click(); err != nil {
		fmt.Println("could not click login button:", err)
		return nil, err
	}
	time.Sleep(1 * time.Second)
	if err := loginButtonLocator.Click(); err != nil {
		fmt.Println("could not click login button:", err)
		return nil, err
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Logged in successfully")

	_, err := shared.SaveContextToFile(page, "loggedIn")
	if err != nil {
		fmt.Println("could not save context to file:", err)
		return nil, err
	}

	return page, nil
}
