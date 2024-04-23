package propi

import (
	"fmt"
	"shared"
)

// 7045811014
// const email = "pravinjoshi61996@gmail.com"
// const password = "pravinjoshi61996@gmail.com"

func FetchProfileUrl() {

	pw, browser, context, page, err := shared.InitPlaywrightPage()
	if err != nil {
		fmt.Println("Error initializing playwright page:", err)
		return
	}
	defer pw.Stop()
	defer browser.Close()
	defer context.Close()
	defer page.Close()

	const startPage = 2422
	const endPage = 5590

	var agents AgentsUrlMap

	err = shared.ReadJsonFile("agentsUrl", &agents)
	if err != nil {
		fmt.Println("could not read agentsUrl file:", err)
		return
	}

	for i := startPage; i <= endPage; i++ {

		url := baseURL + "real-estate-consultants/" + fmt.Sprintf("%d", i) + "-10"
		if _, err := page.Goto(url); err != nil {
			fmt.Printf("could not navigate to page %d: %v", i, err)
			return
		}

		_, err := page.Goto(url)
		if err != nil {
			fmt.Println("could not navigate to page:", err)
			return
		}

		cardsLocator := page.Locator(".btn.btn-teal.mt-1.mr-1.mb-2")
		count, err := cardsLocator.Count()
		if err != nil {
			fmt.Printf("could not count cards: %v", err)
			return
		}
		fmt.Println("Number of cards:", count, "for page:", i)
		if count == 0 {
			fmt.Println("No cards found for page:", i)
			break
		}

		links, err := cardsLocator.All()
		if err != nil {
			fmt.Printf("could not get cards: %v", err)
			return
		}

		for _, link := range links {
			url, err := link.GetAttribute("href")
			if err != nil {
				fmt.Println("could not get href attribute:", err)
				return
			}

			if url == "" {
				fmt.Println("url is empty for page: ", i)
				break
			}

			hash := shared.GenerateMD5Hash(url)
			agentsUrl := AgentUrlDetail{
				Url:        url,
				PageNumber: i,
			}
			agents[hash] = agentsUrl
		}

		shared.WriteJsonFile("agentsUrl", agents)
	}
}

// func fetchProfileUrlDetail() {}
