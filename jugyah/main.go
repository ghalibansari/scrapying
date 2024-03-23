package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"shared"
	"strings"
	"sync"
	"time"

	"github.com/playwright-community/playwright-go"
)

// https://www.propi.in/real-estate-consultants/2-10
const baseURL = "https://www.propi.in/"
const agentFirstPage = 1824
const agentLastPage = 5587
const email = "ghalib@jugyah.com"
const password = "FYM6U6SJaxkgsc8"

type BuildingNestedData struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	ShortBhk     string `json:"shortBhk"`
	DetailLink   string `json:"detailLink"`
	MainImgLink  string `json:"mainImgLink"`
	ShortAddress string `json:"shortAddress"`
}
type Buildings map[string]BuildingNestedData

type AgentNestedData struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	MainImgLink   string `json:"mainImgLink"`
	CompanyName   string `json:"companyName"`
	ContactNumber string `json:"contactNumber"`
	PageNumber    int    `json:"pageNumber"`
}
type Agents map[string]AgentNestedData

func main() {
	print("Main Code Started \n")

	var jsonData Agents
	err := shared.ReadJsonFile("agent.json", &jsonData)
	if err != nil {
		fmt.Println("Error reading json file:", err)
		return
	}

	shared.PrintJson(&jsonData)

	// pw, browser, page, err := initPlaywrightPage()
	// if err != nil {
	// 	fmt.Println("Error initializing playwright page:", err)
	// 	return
	// }
	// defer pw.Stop()
	// defer browser.Close()
	// defer page.Close()

	// for i := agentFirstPage; i <= agentLastPage; i++ {
	// 	time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
	// 	fetchPage(i, page, &jsonData)
	// }

}

func initPlaywrightPage() (*playwright.Playwright, playwright.Browser, playwright.Page, error) {
	pw, err := playwright.Run()
	if err != nil {
		return nil, nil, nil, err
	}

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		// Headless: playwright.Bool(false),
		// SlowMo:   playwright.Float(1_0),
	})
	if err != nil {
		return nil, nil, nil, err
	}

	page, err := browser.NewPage()
	if err != nil {
		return nil, nil, nil, err
	}

	page, err = login(page, email, password)
	if err != nil {
		return nil, nil, nil, err
	}

	return pw, browser, page, nil
}

// func readJsonFile() (Agents, error) {
// 	jsonFile, err := os.Open("agent.json")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer jsonFile.Close()

// 	byteValue, _ := io.ReadAll(jsonFile)

// 	var data Agents
// 	json.Unmarshal(byteValue, &data)

// 	return data, nil
// }

func writeJsonFile(data *Agents) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("agent.json", file, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Data written to file")

	return nil
}

func generateMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func login(page playwright.Page, username string, password string) (playwright.Page, error) {
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

	return page, nil
}

func fetchPage(pageNumber int, page playwright.Page, jsonData *Agents) {
	url := baseURL + "real-estate-consultants/" + fmt.Sprintf("%d", pageNumber) + "-10"
	if _, err := page.Goto(url); err != nil {
		fmt.Println("could not navigate to page:", err)
		return
	}

	ChannelPartnerLocator := page.Locator(`text="Channel Partners"`).Nth(1)
	if err := ChannelPartnerLocator.WaitFor(
		playwright.LocatorWaitForOptions{State: playwright.WaitForSelectorStateAttached, Timeout: playwright.Float(1_000)},
	); err != nil {
		fmt.Println("could not find Channel Partners:", err)
		return
	}
	fmt.Println("Channel Partners found")

	cardsLocator := page.Locator(".card.blog-horizontal")
	count, err := cardsLocator.Count()
	if err != nil {
		fmt.Println("could not find cards:", err)
		return
	}
	fmt.Println("Number of cards:", count)

	cards, err := cardsLocator.All()
	if err != nil {
		fmt.Println("could not retrieve rows:", err)
		return
	}

	var wg sync.WaitGroup
	var mu sync.RWMutex
	for idx, card := range cards {
		wg.Add(1)
		go handleCard(card, jsonData, &wg, &mu, idx, pageNumber)
	}

	wg.Wait()
	fmt.Printf("\n Page fetched: %d \n", pageNumber)

	err = writeJsonFile(jsonData)
	if err != nil {
		fmt.Println("Error writing json file:", err)
		return
	}
}

func handleCard(card playwright.Locator, jsonData *Agents, wg *sync.WaitGroup, mu *sync.RWMutex, idx int, pageNumber int) {
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

	contactNumber := getPhoneNumber(card, 1)
	if contactNumber == "" {
		contactNumber = getPhoneNumber(card, 2)
	}
	if contactNumber == "" {
		fmt.Println("could not get contact number for name: ", name, "companyName: ", companyName, "pageNumber: ", pageNumber)
	}

	id := generateMD5Hash(name + companyName + contactNumber)

	agentData := AgentNestedData{
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

func getPhoneNumber(card playwright.Locator, Nth int) string {
	contactNumberLocator := card.Locator("ul > li").Nth(Nth).Locator("button")
	err := contactNumberLocator.WaitFor(
		playwright.LocatorWaitForOptions{State: playwright.WaitForSelectorStateAttached, Timeout: playwright.Float(1_000)},
	)
	if err != nil {
		// fmt.Println("could not find contact number:", err)
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
		fmt.Println("could not find %d Nth parameter", Nth)
		return ""
	}

	return matches[1]
}
