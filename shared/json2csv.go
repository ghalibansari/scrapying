package shared

import (
	"fmt"
	"strings"
)

func convertJsonToCsv() {
	fmt.Println("Convert json to csv")
	// read json file
	type Agents map[string]map[string]interface{}
	var agents Agents
	err := ReadJsonFile("agent", &agents)
	if err != nil {
		fmt.Println("Error reading json file:", err)
		return
	}

	// convert json to csv, get header and data
	if len(agents) == 0 {
		fmt.Println("No agents found")
		return
	}
	fmt.Println("Agents found:", len(agents))

	// get one agents key
	// var agentKey string
	// mainImgLink companyName contactNumber pageNumber id name address
	var header []string = []string{
		"id", "name", "contactNumber", "companyName", "pageNumber", "address", "mainImgLink",
	}

	// for k := range agents {
	// 	agentKey = k
	// 	break
	// }

	// get header
	// agent := agents[agentKey]
	// fmt.Println("Agent:", agent)

	// for k := range agent {
	// 	header = append(header, k)
	// }

	// fmt.Println("Header:", header)

	// get data [][]string from agents in sequence of header
	var data [][]string
	for _, agent := range agents {
		var row []string
		for _, v := range header {
			// header is address , value is string replce "," with " - "
			if v == "address" {
				agent[v] = strings.Join(strings.Split(agent[v].(string), ","), " - ")
			}
			val := agent[v]
			// if val is nil, append empty string
			if val == nil {
				val = ""
			}
			row = append(row, fmt.Sprintf("%v", val))

		}

		data = append(data, row)
	}

	fmt.Println("Header length:", len(header))
	fmt.Println("Data length:", len(data))

	// print last 5 data
	fmt.Println("Last 5 data:")
	for i := len(data) - 5; i < len(data); i++ {
		fmt.Println(data[i])
	}

	err = WriteCsvFile("igr", header, data)
	if err != nil {
		fmt.Println("Error writing csv file:", err)
		return
	}

}
