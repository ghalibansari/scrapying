package propi

import (
	"fmt"
	"math/rand"
	"shared"
	"strings"
	"sync"
	"time"

	"github.com/playwright-community/playwright-go"
)

var agentFirstPage = 1 // 1
var agentLastPage = 2  // 5587
// const email = "ghalib@jugyah.com"
// const password = "FYM6U6SJaxkgsc8"

func FetchedAgents() {
	print("Propi fetching Started \n")

	//var take input as 1
	var manualInput int
	fmt.Println("Enter 1 for manual input: ")
	fmt.Scanln(&manualInput)

	if manualInput == 1 {
		var startPageInput int
		var endPageInput int
		fmt.Println("Enter Start Page Number: ")
		fmt.Scanln(&startPageInput)
		fmt.Println("Enter End Page Number: ")
		fmt.Scanln(&endPageInput)

		//check if startPageInput is less than endPageInput
		if startPageInput > endPageInput {
			fmt.Println("Start Page Number should be less than End Page Number")
			return
		}

		if startPageInput >= 1 && startPageInput >= agentFirstPage && startPageInput < endPageInput {
			fmt.Println("overwriting agentFirstPage and agentLastPage")
			agentFirstPage = startPageInput
			agentLastPage = endPageInput
		}

		//print agentFirstPage and agentLastPage and startPageInput and endPageInput
		fmt.Println("startPageInput: ", startPageInput)
		fmt.Println("endPageInput: ", endPageInput)
		fmt.Println("agentFirstPage: ", agentFirstPage)
		fmt.Println("agentLastPage: ", agentLastPage)
	}

	var settings Settings
	err := shared.ReadJsonFile("settings", &settings)
	if err != nil {
		fmt.Println("Error reading settings json file:", err)
		return
	}

	var jsonData OldAgentsMap
	err = shared.ReadJsonFile("agent", &jsonData)
	if err != nil {
		fmt.Println("Error reading json file:", err)
		return
	}

	// shared.PrintJson(&jsonData)

	pw, browser, context, page, err := shared.InitPlaywrightPage()
	if err != nil {
		fmt.Println("Error initializing playwright page:", err)
		return
	}
	defer pw.Stop()
	defer browser.Close()
	defer context.Close()
	defer page.Close()

	page, err = Login(page, email, password)
	if err != nil {
		fmt.Println("could not login:", err)
		return
	}

	for i := agentFirstPage; i <= agentLastPage; i++ {
		sleepTime := time.Duration(rand.Intn(4_000))
		print("Sleeping for: ", sleepTime, ", page :- ", i, "\n\n")
		//convert sleepTime to milliseconds

		time.Sleep(sleepTime * time.Millisecond)
		page, err = shared.LoadContextFromFile(page, "loggedIn")
		if err != nil {
			fmt.Println("could not load context from file:", err)
			return
		}
		loggedOut, err := IsLoggedIn(page)
		if err != nil {
			fmt.Println("could not check if logged in:", err)
			return
		}
		if !loggedOut {
			fmt.Println("Not logged in")
			return
		}
		fetchPage(i, page, &jsonData)

	}
	// Temp(page, "loggedIn.json")
	time.Sleep(4 * time.Second)

}

func fetchPage(pageNumber int, page playwright.Page, jsonData *OldAgentsMap) error {
	url := baseURL + "real-estate-consultants/" + fmt.Sprintf("%d", pageNumber) + "-10"
	if _, err := page.Goto(url); err != nil {
		return fmt.Errorf("could not navigate to page %d: %w", pageNumber, err)
	}

	ChannelPartnerLocator := page.Locator(`text="Channel Partners"`).Nth(1)
	if err := ChannelPartnerLocator.WaitFor(
		playwright.LocatorWaitForOptions{State: playwright.WaitForSelectorStateAttached, Timeout: playwright.Float(1_000)},
	); err != nil {
		return fmt.Errorf("could not find Channel Partners text: %w", err)
	}
	fmt.Println("Channel Partners found")

	// check if logged in or not
	// if _, err := page.Goto(baseURL); err != nil {
	// 	fmt.Println("could not navigate to home page:", err)
	// 	return err
	// }

	cardsLocator := page.Locator(".card.blog-horizontal")
	count, err := cardsLocator.Count()
	if err != nil {
		return fmt.Errorf("could not count cards: %w", err)
	}
	fmt.Println("Number of cards:", count)

	cards, err := cardsLocator.All()
	if err != nil {
		return fmt.Errorf("could not get cards: %w", err)
	}

	var wg sync.WaitGroup
	var mu sync.RWMutex
	for idx, card := range cards {
		wg.Add(1)
		go handleCard(card, jsonData, &wg, &mu, idx, pageNumber)
	}

	_, err = shared.SaveContextToFile(page, "loggedIn")
	if err != nil {
		fmt.Println("could not save context to file:", err)
		return err
	}

	wg.Wait()
	fmt.Printf("\nPage fetched: %d \n", pageNumber)

	err = shared.WriteJsonFile("agent", jsonData)
	if err != nil {
		return fmt.Errorf("could not write json file: %w", err)
	}

	return nil
}

func handleCard(card playwright.Locator, jsonData *OldAgentsMap, wg *sync.WaitGroup, mu *sync.RWMutex, idx int, pageNumber int) {
	defer wg.Done()

	imgLocator := card.Locator(".img-fluid.rounded-circle.shadow-1.mb-2")
	imageLink, err := imgLocator.GetAttribute("src")
	if err != nil {
		fmt.Println("could not get src attribute:", err)
		return
	}

	nameLocator := card.Locator("h6").Nth(0)
	name, err := nameLocator.InnerText()
	if err != nil {
		fmt.Println("could not get name:", err)
		return
	}

	companyNameLocator := card.Locator(".font-size-sm.text-success.font-weight-semibold").Nth(0)
	companyName, err := companyNameLocator.InnerText()
	if err != nil {
		fmt.Println("could not get company name:", err)
		return
	}

	addressLocator := card.Locator("p > span")
	addressElements, err := addressLocator.All()
	if err != nil {
		fmt.Println("could not retrieve address elements:", err)
		return
	}

	var address string
	for _, element := range addressElements {
		text, err := element.InnerText()
		if err != nil {
			fmt.Println("could not get address element text:", err)
			continue
		}
		address += text + ", "
	}
	address = strings.TrimSuffix(address, ", ")

	contactNumber := GetPhoneNumber(card, 1)
	if contactNumber == "" {
		contactNumber = GetPhoneNumber(card, 2)
	}
	if contactNumber == "" {
		fmt.Println("could not get contact number for name: ", name, "companyName: ", companyName, "pageNumber: ", pageNumber)
	}

	id := shared.GenerateMD5Hash(name + companyName + address)

	agentData := OldAgentNestedData{
		Id:            id,
		Name:          name,
		Address:       address,
		MainImgLink:   imageLink,
		CompanyName:   companyName,
		ContactNumber: contactNumber,
		PageNumber:    pageNumber,
	}

	mu.Lock()
	(*jsonData)[id] = agentData
	mu.Unlock()

	fmt.Println("Card fetched:", idx)
}
