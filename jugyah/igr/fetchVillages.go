package igr

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

func fetchVillages(page playwright.Page) {
	res, err := page.Goto("https://freesearchigrservice.maharashtra.gov.in/")
	if err != nil {
		fmt.Println("could not navigate to login page:", err)
		return //nil, err
	}

	// print res
	fmt.Printf("%v", res)
}
