package main

import "scrapJD/propi"

func main() {
	print("Main Code Started \n")

	// igr.Main()
	// propi.Main()
	propi.Main()
	// updateAgentHas()
}

// func updateAgentHas() {
// 	// read json file
// 	var jsonData propi.Agents
// 	err := shared.ReadJsonFile("agent1", &jsonData)
// 	if err != nil {
// 		fmt.Println("Error reading json file:", err)
// 		return
// 	}

// 	newAgents := make(propi.Agents)

// 	for _, agent := range jsonData {
// 		// get id, name, company from agent..
// 		name := agent.Name
// 		company := agent.CompanyName
// 		Address := agent.Address

// 		hash := shared.GenerateMD5Hash(name + company + Address)

// 		// check if hash already exists
// 		if old, ok := newAgents[hash]; ok {
// 			fmt.Println("Hash already exists:", hash, old.Name, old.CompanyName, old.Address)
// 			fmt.Println("Hash already exists:", hash, name, company, Address)
// 		}

// 		// if agent.ContactNumber is empty,
// 		if agent.ContactNumber != "" {
// 			agent.Id = hash
// 			newAgents[hash] = agent
// 		}
// 	}

// 	// write json file
// 	err = shared.WriteJsonFile("agent", newAgents)
// 	if err != nil {
// 		fmt.Println("Error writing json file:", err)
// 		return
// 	}
// }
