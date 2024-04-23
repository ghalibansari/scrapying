package propi

import (
	"fmt"
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

	// for _, agent := range agentsUrl {
	// FetchAgentDetail(page, "/profile/padmanand-birajdar-buddhay-royal-enterprises/27869")
	FetchAgentDetail(page, "/profile/ravindra-gupta-seabury-realtors/109410")
	// }

	//length
	fmt.Println(len(agentsUrl), "agents found in agentsUrl")

	//sleep
	time.Sleep(1 * time.Second)

}

// func FetchedAgentDetailFromAgentUrlsMap(){}
