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

func FetchAgentDetail(page playwright.Page, agentUrl string) (map[string]interface{}, error) {

	url := baseURL + agentUrl
	page.Goto(url)

	var agentData map[string]interface{} = make(map[string]interface{})
	agentData["properties"] = make([]map[string]string, 0)
	agentData["projects"] = make([]map[string]string, 0)

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
		fmt.Println(msg)
		panic(msg)
	}

	rowLocator := page.Locator(".row")
	rows, err := rowLocator.All()
	if err != nil {
		fmt.Println("could not get div with class name row:", err)
		return nil, err
	}

	userDetail, err := rows[1].Locator(".col-md-6.col-sm-6.col-lg-4").All()
	if err != nil {
		fmt.Println("could not get div with class name col-md-6 col-sm-6 col-lg-4:", err)
		return nil, err
	}

	MRCNText, err := userDetail[0].Locator("#ContentPlaceHolderMain_trRERA td.text-right").First().InnerText(
		playwright.LocatorInnerTextOptions{Timeout: playwright.Float(1_000)},
	)
	if err != nil {
		fmt.Println("could not get table:", err)
	}
	MRCNText = strings.TrimSpace(MRCNText)
	agentData["MRCN"] = MRCNText

	/////////////////
	Location, err := userDetail[0].Locator("#ContentPlaceHolderMain_trAdd > td:nth-child(2) > div:nth-child(2)").First().InnerText(
		playwright.LocatorInnerTextOptions{Timeout: playwright.Float(1_000)},
	)
	if err != nil {
		fmt.Println("could not get Location:", err)
	}
	Location = strings.TrimSpace(Location)
	agentData["Location"] = Location

	/////////////////
	specializationsLocator, err := userDetail[0].Locator("#ContentPlaceHolderMain_tr7 > td > span").All()
	if err != nil {
		fmt.Println("could not get specializations:", err)
	}

	var specializations []string = []string{}
	for _, specialization := range specializationsLocator {
		text, err := specialization.InnerText(
			playwright.LocatorInnerTextOptions{Timeout: playwright.Float(1_000)},
		)
		if err != nil {
			fmt.Println("could not get specialization text:", err)
		}
		specializations = append(specializations, text)
	}
	agentData["specializations"] = specializations

	////////////
	AreasOfOperationLocator, err := userDetail[0].Locator("#ContentPlaceHolderMain_trAOP > td > span").All()
	if err != nil {
		fmt.Println("could not get AreasOfOperation:", err)
	}

	var AreasOfOperation []string = []string{}
	for _, AreaOfOperation := range AreasOfOperationLocator {
		text, err := AreaOfOperation.InnerText(
			playwright.LocatorInnerTextOptions{Timeout: playwright.Float(1_000)},
		)
		if err != nil {
			fmt.Println("could not get AreaOfOperation text:", err)
		}
		AreasOfOperation = append(AreasOfOperation, text)
	}
	agentData["AreasOfOperation"] = AreasOfOperation

	////
	ActiveInBuildingsLocator, err := userDetail[0].Locator("#ContentPlaceHolderMain_tr1 > td > span").All()
	if err != nil {
		fmt.Println("could not get ActiveInBuildings:", err)
	}

	var ActiveInBuildings []string = []string{}
	for _, ActiveInBuilding := range ActiveInBuildingsLocator {
		text, err := ActiveInBuilding.InnerText(
			playwright.LocatorInnerTextOptions{Timeout: playwright.Float(1_000)},
		)
		if err != nil {
			fmt.Println("could not get ActiveInBuilding text:", err)
		}
		ActiveInBuildings = append(ActiveInBuildings, text)
	}
	agentData["ActiveInBuildings"] = ActiveInBuildings

	///
	ActiveInLocationsLocator, err := userDetail[0].Locator("#ContentPlaceHolderMain_tr5 > td > span").All()
	if err != nil {
		fmt.Println("could not get ActiveInLocations:", err)
	}

	var ActiveInLocations []string = []string{}
	for _, ActiveInLocation := range ActiveInLocationsLocator {
		text, err := ActiveInLocation.InnerText(
			playwright.LocatorInnerTextOptions{Timeout: playwright.Float(1_000)},
		)
		if err != nil {
			fmt.Println("could not get ActiveInLocation text:", err)
		}
		ActiveInLocations = append(ActiveInLocations, text)
	}
	agentData["ActiveInLocations"] = ActiveInLocations

	err = handlePropertiesTab(page, agentData)
	if err != nil {
		fmt.Println("could not handle properties tab:", err)
	}

	err = handleProjectsTab(page, agentData)
	if err != nil {
		fmt.Println("could not handle projects tab:", err)
	}

	return agentData, nil
}

func handlePropertiesTab(page playwright.Page, agentData map[string]interface{}) error {
	// get properties of tab
	const tabId = "a[href=\"#Properties\"]"
	// const tabId = "a[href=\"#Requirements\"]"
	propertiesLink := page.Locator(tabId).First()

	err := propertiesLink.Click()
	if err != nil {
		fmt.Println("could not click on properties link:", err)
		return nil
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

	getPropertyTableData(page, agentData)

	checkPaginationExistForPropertyTab(page, agentData)

	return nil

}

func getPropertyTableData(page playwright.Page, agentData map[string]interface{}) {
	propertiesTableBody := page.Locator("#Properties.tab-pane.fade > .card-body > table > tbody > tr > td> .dxgvCSD > table > tbody").First()

	propertiesTableTableRows, err := propertiesTableBody.Locator("tr").All()
	if err != nil {
		fmt.Println("could not get properties table data:", err)
		return
	}

	handlePropertiesTabRowsData(propertiesTableTableRows, agentData)
}

func handlePropertiesTabRowsData(rows []playwright.Locator, agentData map[string]interface{}) []map[string]string {

	fmt.Println("Number of properties ", len(rows))

	// array for map[string]string
	var properties []map[string]string

	for i, row := range rows {
		// start at 4th row
		fmt.Println("row number:", i)
		if i > 3 {

			property, err := handlePropertyTabColumnData(row)
			if err != nil {
				fmt.Println("could not get property data:", err)
			}
			properties = append(properties, property)
		}
	}

	// print properties
	// fmt.Println("properties:", properties)

	// property := properties[0]
	// var header []string
	// for key := range property {
	// 	header = append(header, key)
	// }

	// var data [][]string
	// for _, value := range properties {
	// 	var row []string
	// 	for _, val := range header {
	// 		row = append(row, value[val])
	// 	}
	// 	data = append(data, row)
	// }
	// shared.WriteCsvFile("properties1", header, data)

	var oldProperties []map[string]string = agentData["properties"].([]map[string]string)
	var newProperties []map[string]string = append(oldProperties, properties...)
	agentData["properties"] = newProperties
	return properties
}

func handlePropertyTabColumnData(row playwright.Locator) (map[string]string, error) {
	columns, err := row.Locator("td").All()
	if err != nil {
		fmt.Println("could not get columns in property table: ", err)
		return nil, err
	}

	propertyColumn := columns[0]
	propertyLink, err := propertyColumn.Locator("a").First().GetAttribute("href")
	if err != nil {
		fmt.Println("could not get property link:", err)
		return nil, err
	}
	propertyName, err := propertyColumn.Locator("a").First().TextContent()
	if err != nil {
		fmt.Println("could not get property name:", err)
		return nil, err
	}
	propertyName = strings.Join(strings.Fields(propertyName), " ")
	propertyName = strings.TrimSpace(propertyName)

	locationColumn := columns[1]
	locationLink, err := locationColumn.Locator("a").First().GetAttribute("href")
	if err != nil {
		fmt.Println("could not get location link:", err)
		return nil, err
	}
	location, err := locationColumn.InnerText()
	if err != nil {
		fmt.Println("could not get location:", err)
		return nil, err
	}
	location = strings.TrimSpace(location)

	priceColumn := columns[2]
	price, err := priceColumn.InnerText()
	if err != nil {
		fmt.Println("could not get price:", err)
		return nil, err
	}

	buildingColumn := columns[3]
	building, err := buildingColumn.InnerText()
	if err != nil {
		fmt.Println("could not get building:", err)
		return nil, err
	}

	floorColumn := columns[4]
	floor, err := floorColumn.InnerText()
	if err != nil {
		fmt.Println("could not get floor:", err)
		return nil, err
	}

	createdAtColumn := columns[5]
	createdAt, err := createdAtColumn.InnerText()
	if err != nil {
		fmt.Println("could not get created at:", err)
		return nil, err
	}

	updatedAtColumn := columns[6]
	updatedAt, err := updatedAtColumn.InnerText()
	if err != nil {
		fmt.Println("could not get updated at:", err)
		return nil, err
	}

	property := map[string]string{
		"propertyLink": propertyLink,
		"propertyName": propertyName,
		"locationLink": locationLink,
		"location":     location,
		"price":        price,
		"building":     building,
		"floor":        floor,
		"createdAt":    createdAt,
		"updatedAt":    updatedAt,
	}

	return property, nil
}

func checkPaginationExistForPropertyTab(page playwright.Page, agentData map[string]interface{}) {
	checkPaginationExistsLocator := page.Locator("#Properties.tab-pane.fade > .card-body > table > tbody > tr > td > .dxgvPagerBottomPanel_Material")

	count, err := checkPaginationExistsLocator.Count()
	if err != nil {
		fmt.Println("could not get paginationExists:", err)
		return
	}

	if count == 0 {
		fmt.Println("No pagination exists")
		return
	}

	clickOnNextButtonForPropertyTab(page, agentData)
}

func clickOnNextButtonForPropertyTab(page playwright.Page, agentData map[string]interface{}) {

	paginationNextButtonLocator := page.Locator("#ContentPlaceHolderMain_MyProps_DXPagerBottom_PBN").First()

	isDisabled, err := paginationNextButtonLocator.IsDisabled()
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
	if isDisabled {
		fmt.Println("Next button is disabled")
		return
	}

	err = paginationNextButtonLocator.Click()
	if err != nil {
		fmt.Println("could not click on next button:", err)
		return
	}

	getPropertyTableData(page, agentData)
	clickOnNextButtonForPropertyTab(page, agentData)
}

func handleProjectsTab(page playwright.Page, agentData map[string]interface{}) error {
	const tabId = "a[href=\"#Projects\"]"

	projectTab := page.Locator(tabId).First()

	err := projectTab.Click()
	if err != nil {
		fmt.Println("could not click on projects link:", err)
		return nil
	}

	getProjectTableData(page, agentData)

	checkPaginationExistForProjectTab(page, agentData)

	return nil
}

func getProjectTableData(page playwright.Page, agentData map[string]interface{}) {
	projectsTableBody := page.Locator("#Projects.tab-pane.fade > .card-body > table > tbody > tr > td> .dxgvCSD > table > tbody").First()

	projectsTableRows, err := projectsTableBody.Locator("tr").All()
	if err != nil {
		fmt.Println("could not get projects table data:", err)
		return
	}

	handleProjectsTabRowsData(projectsTableRows, agentData)
}

func handleProjectsTabRowsData(rows []playwright.Locator, agentData map[string]interface{}) []map[string]string {

	fmt.Println("Number of projects ", len(rows))

	// array for map[string]string
	var projects []map[string]string

	for i, row := range rows {
		// start at 4th row
		fmt.Println("row number:", i)
		if i > 3 {

			project, err := handleProjectTabColumnData(row)
			if err != nil {
				fmt.Println("could not get project data:", err)
			}
			projects = append(projects, project)
		}
	}

	var oldProjects []map[string]string = agentData["projects"].([]map[string]string)
	var newProjects []map[string]string = append(oldProjects, projects...)
	agentData["projects"] = newProjects
	return projects
}

func handleProjectTabColumnData(row playwright.Locator) (map[string]string, error) {
	columns, err := row.Locator("td").All()
	if err != nil {
		fmt.Println("could not get columns in property table: ", err)
		return nil, err
	}

	projectColumn := columns[0]
	projectLink, err := projectColumn.Locator("a").First().GetAttribute("href")
	if err != nil {
		fmt.Println("could not get project link:", err)
		return nil, err
	}

	projectText, err := projectColumn.Locator("a").First().TextContent()
	if err != nil {
		fmt.Println("could not get project name:", err)
		return nil, err
	}

	developerColumnText, err := columns[1].InnerText()
	if err != nil {
		fmt.Println("could not get developer:", err)
		return nil, err
	}

	locationColumnText, err := columns[2].InnerText()
	if err != nil {
		fmt.Println("could not get location:", err)
		return nil, err
	}

	mAHAReRaNoColumnText, err := columns[3].InnerText()
	if err != nil {
		fmt.Println("could not get mAHAReRaNo:", err)
		return nil, err
	}

	project := map[string]string{
		"projectLink": projectLink,
		"projectText": projectText,
		"developer":   developerColumnText,
		"location":    locationColumnText,
		"mAHAReRaNo":  mAHAReRaNoColumnText,
	}

	return project, nil
}

func checkPaginationExistForProjectTab(page playwright.Page, agentData map[string]interface{}) {
	checkPaginationExistsLocator := page.Locator("#Projects.tab-pane.fade > .card-body > table > tbody > tr > td > .dxgvPagerBottomPanel_Material")

	count, err := checkPaginationExistsLocator.Count()
	if err != nil {
		fmt.Println("could not get paginationExists:", err)
		return
	}

	if count == 0 {
		fmt.Println("No pagination exists")
		return
	}

	clickOnNextButtonForProjectTab(page, agentData)
}

func clickOnNextButtonForProjectTab(page playwright.Page, agentData map[string]interface{}) {

	paginationNextButtonLocator := page.Locator("#ContentPlaceHolderMain_ASPxGridViewProjs_DXPagerBottom_PBN").First()

	err := paginationNextButtonLocator.WaitFor(
		playwright.LocatorWaitForOptions{
			Timeout: playwright.Float(2_000),
		},
	)
	if err != nil {
		fmt.Println("could not wait for next button for project tab:", err)
		return
	}

	isDisabled, err := paginationNextButtonLocator.GetAttribute("aria-disabled")
	if err != nil {
		fmt.Println("could not get paginationNextButton html:", err)
		return
	}

	if isDisabled == "true" || isDisabled != "" {
		fmt.Println("Next button is disabled")
		return
	}

	err = paginationNextButtonLocator.Click()
	if err != nil {
		fmt.Println("could not click on next button:", err)
		return
	}

	getProjectTableData(page, agentData)
	clickOnNextButtonForProjectTab(page, agentData)
}
