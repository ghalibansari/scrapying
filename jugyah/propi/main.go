package propi

import (
	"fmt"
	"math/rand"
	"shared"
	"time"
)

const baseURL = "https://www.propi.in/"
const email = "pravinjoshi61996@gmail.com"
const password = "pravinjoshi61996@gmail.com"

func Main() {
	pw, browser, context, page, err := shared.InitPlaywrightPage()
	if err != nil {
		fmt.Println("Error initializing playwright page:", err)
		return
	}
	defer pw.Stop()
	defer browser.Close()
	defer context.Close()
	defer page.Close()

	agentsUrl := LoadAgentUrlsMap("agentsUrl")
	fmt.Println(len(agentsUrl), "agents found in agentsUrl")

	//
	page.SetViewportSize(1920, 1080)

	for key, agent := range agentsUrl {
		if !agent.Fetched {
			// random sleep
			randomSeconds := rand.Intn(11) // generates a random number between 0 and 5
			time.Sleep(time.Duration(randomSeconds) * time.Second)

			data, err := FetchAgentDetail(page, agent.Url)
			if err != nil {
				fmt.Println("Error fetching agent detail:", err)
				continue
			}
			agent.Data = data
			fmt.Println("Data:", agent.Fetched)

			agent.Fetched = true
			agentsUrl[key] = agent

			fmt.Println("Fetched agent detail for", agent.Url)
			fmt.Println("Data:", agent.Fetched)

			//save
			shared.WriteJsonFile("agentsUrl1", agentsUrl)
		}

		// FetchAgentDetail(page, "/profile/padmanand-birajdar-buddhay-royal-enterprises/27869")
		// FetchAgentDetail(page, "/profile/ravindra-gupta-seabury-realtors/109410")
	}

	//length
	fmt.Println(len(agentsUrl), "agents found in agentsUrl")

	//sleep
	time.Sleep(1 * time.Second)

}

// func FetchedAgentDetailFromAgentUrlsMap(){}
