package propi

import (
	"fmt"
	"shared"
	"strings"

	"github.com/playwright-community/playwright-go"
)

func LoadAgentUrlsMap(fileName string) AgentsUrlMap {
	var agentsUrl AgentsUrlMap
	shared.ReadJsonFile(fileName, &agentsUrl)

	return agentsUrl
}

func FetchAgentDetail(page playwright.Page, agentUrl string) {
	url := baseURL + agentUrl
	page.Goto(url)

	// googAfterNoonText := "text=Afternoon, Pravin!"
	googAfterNoonText := "text=Guest!"
	goodAfterNoonLocator := page.Locator(googAfterNoonText).Nth(0)
	if err := goodAfterNoonLocator.WaitFor(
		playwright.LocatorWaitForOptions{
			State:   playwright.WaitForSelectorStateAttached,
			Timeout: playwright.Float(1_000),
		},
	); err != nil {
		msg := fmt.Sprintf("could not find good afternoon text: user Not logged in: %v", err)
		// panic(msg)
		fmt.Println(msg)
	}

	// get div with class name "row"
	rowLocator := page.Locator(".row")
	rows, err := rowLocator.All()
	if err != nil {
		fmt.Println("could not get div with class name row:", err)
		return
	}

	// get div with class name "col-md-6 col-sm-6 col-lg-4"
	userDetail, err := rows[1].Locator(".col-md-6.col-sm-6.col-lg-4").All()
	if err != nil {
		fmt.Println("could not get div with class name col-md-6 col-sm-6 col-lg-4:", err)
		return
	}

	///////////
	MRCNText, err := userDetail[0].Locator("#ContentPlaceHolderMain_trRERA td.text-right").First().InnerText(
		playwright.LocatorInnerTextOptions{Timeout: playwright.Float(1_000)},
	)
	if err != nil {
		fmt.Println("could not get table:", err)
	}

	MRCNText = strings.TrimSpace(MRCNText)
	fmt.Println("Table:", MRCNText)

	/////////////////
	Location, err := userDetail[0].Locator("#ContentPlaceHolderMain_trAdd > td:nth-child(2) > div:nth-child(2)").First().InnerText(
		playwright.LocatorInnerTextOptions{Timeout: playwright.Float(1_000)},
	)
	if err != nil {
		fmt.Println("could not get Location:", err)
	}

	Location = strings.TrimSpace(Location)

	fmt.Println("Location:", Location)

	/////////////////
	specializations, err := userDetail[0].Locator("#ContentPlaceHolderMain_tr7 > td > span").All()
	if err != nil {
		fmt.Println("could not get specializations:", err)
		return
	}
	fmt.Println("Number of specializations:", len(specializations), "pppppppp")

	var specializations1 []string = []string{}
	for _, specialization := range specializations {
		text, err := specialization.InnerText(
			playwright.LocatorInnerTextOptions{Timeout: playwright.Float(1_000)},
		)
		if err != nil {
			fmt.Println("could not get specialization text:", err)
			return
		}
		specializations1 = append(specializations1, text)
	}

	fmt.Println("specializations1:", specializations1)

	////////////
	AreasOfOperation, err := userDetail[0].Locator("#ContentPlaceHolderMain_trAOP > td > span").All()
	if err != nil {
		fmt.Println("could not get AreasOfOperation:", err)
		return
	}

	var AreasOfOperation1 []string = []string{}
	for _, AreaOfOperation := range AreasOfOperation {
		text, err := AreaOfOperation.InnerText(
			playwright.LocatorInnerTextOptions{Timeout: playwright.Float(1_000)},
		)
		if err != nil {
			fmt.Println("could not get AreaOfOperation text:", err)
			return
		}
		AreasOfOperation1 = append(AreasOfOperation1, text)
	}

	fmt.Println("AreasOfOperation1:", AreasOfOperation1)

	////
	ActiveInBuildings, err := userDetail[0].Locator("#ContentPlaceHolderMain_tr1 > td > span").All()
	if err != nil {
		fmt.Println("could not get ActiveInBuildings:", err)
		return
	}

	var ActiveInBuildings1 []string = []string{}
	for _, ActiveInBuilding := range ActiveInBuildings {
		text, err := ActiveInBuilding.InnerText(
			playwright.LocatorInnerTextOptions{Timeout: playwright.Float(1_000)},
		)
		if err != nil {
			fmt.Println("could not get ActiveInBuilding text:", err)
			return
		}
		ActiveInBuildings1 = append(ActiveInBuildings1, text)
	}

	fmt.Println("ActiveInBuildings1:", ActiveInBuildings1)

	///
	ActiveInLocations, err := userDetail[0].Locator("#ContentPlaceHolderMain_tr5 > td > span").All()
	if err != nil {
		fmt.Println("could not get ActiveInLocations:", err)
		return
	}

	var ActiveInLocations1 []string = []string{}
	for _, ActiveInLocation := range ActiveInLocations {
		text, err := ActiveInLocation.InnerText(
			playwright.LocatorInnerTextOptions{Timeout: playwright.Float(1_000)},
		)
		if err != nil {
			fmt.Println("could not get ActiveInLocation text:", err)
			return
		}
		ActiveInLocations1 = append(ActiveInLocations1, text)
	}

	fmt.Println("ActiveInLocations1:", ActiveInLocations1)
	fmt.Println("/////////////////////////////=============================")

	//////////////////////
	/////////////////////
	getPropertiesTab(page)
	//

	/////////////////////
	/////////////////////
	// getProjectsTab(page)

	/////////////////////
	/////////////////////
	checkPaginationExistForPropertyTab(page)

	/////////////////////////
	fmt.Println("User is logged in")
}

func getPropertiesTab(page playwright.Page) {
	// get properties of tab
	const tabId = "a[href=\"#Properties\"]"
	// const tabId = "a[href=\"#Requirements\"]"
	propertiesLink := page.Locator(tabId).First()

	//print properties html
	propertiesHtml, err := propertiesLink.InnerHTML()
	if err != nil {
		fmt.Println("could not get properties html:", err)
		return
	}
	fmt.Println("properties html:xxxxxxxxxxxxxx", propertiesHtml, "kkkkkkkkkkkkkkk")

	propertiesLink = page.Locator(tabId).First()

	err = propertiesLink.Click()
	if err != nil {
		fmt.Println("could not click on properties link:", err)
		return
	}

	// properties table
	// propertiesTableLocator := page.Locator("#Properties.tab-pane.fade > .card-body > table > tbody > tr > td").First()

	//-- table header

	// propertiesTable, err := propertiesTableLocator.First().InnerHTML()

	// if err != nil {
	// 	fmt.Println("could not get properties table:", err)
	// 	return
	// }
	// fmt.Println("properties table:", propertiesTable)

	// // get first div of properties table locator
	// propertyTableHeaderLocator := propertiesTableLocator.Locator("div").First()

	// // print properties table header html
	// propertyTableHeaderHtml, err := propertyTableHeaderLocator.InnerHTML()
	// if err != nil {
	// 	fmt.Println("could1111111 not get properties table header html:", err)
	// 	return
	// }

	propertiesTableBody := page.Locator("#Properties.tab-pane.fade > .card-body > table > tbody > tr > td> .dxgvCSD > table > tbody").First()

	propertiesTableDatas, err := propertiesTableBody.Locator("tr").All()
	if err != nil {
		fmt.Println("could not get properties table datas:", err)
		return
	}

	fmt.Println("Number of properties table datas:", len(propertiesTableDatas))

	// for _, propertiesTableData := range propertiesTableDatas {
	// 	propertiesTableDataHtml, err := propertiesTableData.InnerHTML()
	// 	if err != nil {
	// 		fmt.Println("could not get properties table data html:", err)
	// 		return
	// 	}
	// 	fmt.Println("properties table data html:", propertiesTableDataHtml)
	// }
	// get first tr of properties table datas == skip 3
	propertiesTableData := propertiesTableDatas[4] // 4 is row number
	data, _ := propertiesTableData.Locator("td").All()

	// len data
	fmt.Println("Number of properties table data:222222222222222222222222", len(data))
	firstProperties := data[0]

	firstPropertiesLink, err := firstProperties.Locator("a").First().GetAttribute("href")
	if err != nil {
		fmt.Println("could not get first properties link:", err)
		return
	}
	fmt.Println("first properties link:", firstPropertiesLink)

	firstPropertiesText, err := firstProperties.Locator("a").First().TextContent()
	if err != nil {
		fmt.Println("could not get first properties html:", err)
		return
	}

	// trim first properties text
	fmt.Println("first properties html:", strings.TrimSpace(firstPropertiesText), "kkkkkkkkkkkkkkk")

	locationData := data[1]
	locationText, err := locationData.InnerText()
	if err != nil {
		fmt.Println("could not get other data text:", err)
		return
	}
	locationUrl, err := locationData.Locator("a").First().GetAttribute("href")
	if err != nil {
		fmt.Println("could not get other data text:", err)
		return
	}
	fmt.Println("other data text:333333333333333", locationUrl)

	fmt.Println("other data text:444444444444444", locationText)

	priceData := data[2]
	priceText, err := priceData.InnerText()
	if err != nil {
		fmt.Println("could not get other data text:", err)
		return
	}
	fmt.Println("other data text:555555555555555", priceText)

	buildingData := data[3]
	buildingText, err := buildingData.InnerText()
	if err != nil {
		fmt.Println("could not get other data text:", err)
		return
	}
	fmt.Println("other data text:666666666666666", buildingText)

	floorData := data[4]
	floorText, err := floorData.InnerText()
	if err != nil {
		fmt.Println("could not get other data text:", err)
		return
	}
	fmt.Println("other data text:777777777777777", floorText)

	addedData := data[5]
	addedText, err := addedData.InnerText()
	if err != nil {
		fmt.Println("could not get other data text:", err)
		return
	}
	fmt.Println("other data text:888888888888888", addedText)

	updatedData := data[6]
	updatedText, err := updatedData.InnerText()
	if err != nil {
		fmt.Println("could not get other data text:", err)
		return
	}
	fmt.Println("other data text:999999999999999", updatedText)
}

func getProjectsTab(page playwright.Page) {
	const tabId = "a[href=\"#Projects\"]"

	projectTab := page.Locator(tabId).First()

	err := projectTab.Click()
	if err != nil {
		fmt.Println("could not click on projects link:", err)
		return
	}

	projectsTableBody := page.Locator("#Projects.tab-pane.fade > .card-body > table > tbody > tr > td> .dxgvCSD > table > tbody").First()

	projectsTableRows, err := projectsTableBody.Locator("tr").All()
	if err != nil {
		fmt.Println("could not get projects table datas:", err)
		return
	}

	fmt.Println("Number of projects table datas:ccccccccccc", len(projectsTableRows))

	row := projectsTableRows[3] // row

	columns, _ := row.Locator("td").All()

	// projectColumn := row.Locator("td > a").First()
	projectColumn := columns[0].Locator("a").First()

	projectLink, err := projectColumn.GetAttribute("href")
	if err != nil {
		fmt.Println("could not get project link:", err)
		return
	}
	fmt.Println("project link:11111111111111xxxxxxxxxxxxxx", projectLink)

	projectText, err := projectColumn.TextContent()
	if err != nil {
		fmt.Println("could not get project html:", err)
		return
	}
	fmt.Println("project html:xxxxxxxxxxxxxx", projectText, "kkkkkkkkkkkkkkk")

	developerColumn, _ := columns[1].InnerText()
	fmt.Println("developerColumn:xxxxxxxxxxxxxx", developerColumn, "oooooooooooooooo")

	locationColumn, _ := columns[2].InnerText()
	fmt.Println("locationColumn:xxxxxxxxxxxxxx", locationColumn, "oooooooooooooooo")

	mAHAReRaNoText, _ := columns[3].InnerText()
	fmt.Println("mAHAReRaNoText:xxxxxxxxxxxxxx", mAHAReRaNoText, "oooooooooooooooo")
}

func checkPaginationExistForPropertyTab(page playwright.Page) {

	checkPaginationExistsLocator := page.Locator("#Properties.tab-pane.fade > .card-body > table > tbody > tr > td > .dxgvPagerBottomPanel_Material")

	count, err := checkPaginationExistsLocator.Count()
	if err != nil {
		fmt.Println("could not get paginationExists:", err)
		return
	}

	fmt.Println("paginationExists:", "Number of paginationExists:", count)

	// return paginationExists, len(checkPaginationExists)
	getPaginationCount(page)
}

func getPaginationCount(page playwright.Page) {
	paginationNextButtonLocator := page.Locator("#ContentPlaceHolderMain_MyProps_DXPagerBottom_PBN").First()

	err := paginationNextButtonLocator.Click()
	if err != nil {
		fmt.Println("could not click on next button:", err)
		return
	}

}

func clickOnNextButtonForPropertyTab(paginationNextButtonLocator playwright.Locator) {
	isEnabled, err := paginationNextButtonLocator.IsEnabled()
	if err != nil {
		fmt.Println("could not get next button is enabled:", err)
		return
	}

	err = paginationNextButtonLocator.WaitFor(
		playwright.LocatorWaitForOptions{
			Timeout: playwright.Float(2_000),
		},
	)
	if err != nil {
		fmt.Println("could not wait for next button:", err)
		return
	}
	if !isEnabled {
		fmt.Println("Next button is not enabled")
		return
	}

	err = paginationNextButtonLocator.Click()
	if err != nil {
		fmt.Println("could not click on next button:", err)
		return
	}

	clickOnNextButtonForPropertyTab(paginationNextButtonLocator)
}
