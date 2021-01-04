#Ex1
|            |
|------------|
| <img src="https://github.com/1stApr/Go_Exercise/blob/master/GoImage.png" width="750"> |
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Product struct {
	//Type   string  `json:"type"`
	VCPU   float64 `json:"vCPU"`
	VRam   float64 `json:"vRam"`
	Counts float64 `json:"counts"`
}

func isJsonFile(fileName string) bool {
	temp := fileName[len(fileName)-7:]
	temp = strings.TrimRight(temp, "\r\n")
	if strings.Compare(temp, ".json") == 0 {
		return true
	} else {
		return false
	}
}

func isExit(fileName string) bool {
	temp := fileName[len(fileName)-5:]
	if strings.Compare(temp, "Exit") == 0 {
		return true
	} else {
		return false
	}
}

// func readData(fileName string) *Instance {
// 	fileName = strings.TrimRight(fileName, "\r\n")
// 	jsonFile, err := os.Open(fileName)
// 	if err != nil {
// 		//fmt.Println("Error!")
// 		fmt.Println(err)
// 	}
// 	defer jsonFile.Close()
// 	//instance := make([]Instance, 3)
// 	var instance Instance
// 	raw, err := ioutil.ReadFile(fileName)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		os.Exit(1)
// 	}
// 	json.Unmarshal(raw, &instance)
// 	return &instance
// }

// func printProduct(instance *Instance) {
// 	for i := 0; i < len(instance.Instance); i++ {
// 		fmt.Printf("Type: %-10v", instance.Instance[i].Type)
// 		fmt.Printf(" vCPU: %-10v", instance.Instance[i].VCPU)
// 		fmt.Printf(" vRam: %-10v", instance.Instance[i].VRam)
// 		fmt.Printf(" counts: %-10v", instance.Instance[i].Counts)
// 		fmt.Printf("\n")
// 	}
// 	fmt.Print("\n")
// }

// func printResult(oldInstance, newInstance *Instance) {
// 	for i := 0; i < len(newInstance.Instance); i++ {
// 		temp := newInstance.Instance[i].Counts - oldInstance.Instance[i].Counts
// 		if temp >= 0 {
// 			//fmt.Println("[\""+oldInstance.Instance[i].Type+"\"]"+" [provision] [", temp, "]")
// 			fmt.Printf("[\"%-9v\"]	[%-9v] [%-1v]\n", oldInstance.Instance[i].Type, "provision", temp)
// 		} else {
// 			//fmt.Println("[\""+oldInstance.Instance[i].Type+"\"]"+" [delete] [", -temp, "]")
// 			fmt.Printf("[\"%-9v\"]	[%-9v] [%-1v]\n", oldInstance.Instance[i].Type, "delete", -temp)
// 		}
// 	}

// }

var check bool

func readData(filePath string) map[string]Product {

	data, _ := ioutil.ReadFile(filePath)

	var result map[string]interface{}
	rs := make(map[string]Product)
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("Error\n")
		log.Fatal(err)
	}
	//fmt.Println(result)
	for _, v := range result["Instances"].([]interface{}) {
		instance := v.(map[string]interface{})
		fmt.Printf("V [%T]:= %v\n", v, v)
		fmt.Printf("instance [%T]:= %v\n", instance, instance)
		//rs[instance["Type"].(string)] = Product{instance["VCPU"].(float64), instance["VRAM"].(float64), instance["Counts"].(float64)}

	}
	return rs

}
func main() {

	fmt.Println(readData("C:/Users/Admin/Desktop/Go_Exercise/config.json"))

	// check = false
	// var oldInstance *Instance
	// var newInstance *Instance
	// //var quere []Instance
	// reader := bufio.NewReader(os.Stdin)
	// for exit := 1; exit != 2; {
	// 	fmt.Print("Enter path: ")
	// 	path, _ := reader.ReadString('\n')
	// 	if isExit(path) {
	// 		exit = 2
	// 		break
	// 	} else {
	// 		if isJsonFile(path) {
	// 			//readData(path)
	// 			if check {
	// 				oldInstance = newInstance
	// 				newInstance = readData(path)
	// 				printProduct(oldInstance)
	// 				printProduct(newInstance)
	// 				printResult(oldInstance, newInstance)
	// 			} else {
	// 				newInstance = readData(path)
	// 				printProduct(newInstance)
	// 				check = true
	// 			}

	// 		} else {
	// 			fmt.Println("Invalid file! Enter path again: ")
	// 		}
	// 	}
	// }

}
