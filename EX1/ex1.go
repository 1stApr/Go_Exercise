package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Instance struct {
	Instance []Product `json:"Instances"`
}
type Product struct {
	Type   string `json:"type"`
	VCPU   int    `json:"vCPU"`
	VRam   int    `json:"vRam"`
	Counts int    `json:"counts"`
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

func readData(fileName string) *Instance {
	fileName = strings.TrimRight(fileName, "\r\n")
	jsonFile, err := os.Open(fileName)
	if err != nil {
		//fmt.Println("Error!")
		fmt.Println(err)
	}
	defer jsonFile.Close()
	//instance := make([]Instance, 3)
	var instance Instance
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &instance)
	return &instance
}

func printProduct(instance *Instance) {
	for i := 0; i < len(instance.Instance); i++ {
		fmt.Printf("Type: %-10v", instance.Instance[i].Type)
		fmt.Printf(" vCPU: %-10v", instance.Instance[i].VCPU)
		fmt.Printf(" vRam: %-10v", instance.Instance[i].VRam)
		fmt.Printf(" counts: %-10v", instance.Instance[i].Counts)
		fmt.Printf("\n")
	}
	fmt.Print("\n")
}

func printResult(oldInstance, newInstance *Instance) {
	for i := 0; i < len(newInstance.Instance); i++ {
		temp := newInstance.Instance[i].Counts - oldInstance.Instance[i].Counts
		if temp >= 0 {
			//fmt.Println("[\""+oldInstance.Instance[i].Type+"\"]"+" [provision] [", temp, "]")
			fmt.Printf("[\"%-9v\"]	[%-9v]	[%-1v]\n", oldInstance.Instance[i].Type, "provision", temp)
		} else {
			//fmt.Println("[\""+oldInstance.Instance[i].Type+"\"]"+" [delete] [", -temp, "]")
			fmt.Printf("[\"%-9v\"]	[%-9v]	[%-1v]\n", oldInstance.Instance[i].Type, "delete", -temp)
		}
	}

}

var check bool

func main() {
	check = false
	var oldInstance *Instance
	var newInstance *Instance
	//var quere []Instance
	reader := bufio.NewReader(os.Stdin)
	for exit := 1; exit != 2; {
		fmt.Print("Enter path: ")
		path, _ := reader.ReadString('\n')
		if isExit(path) {
			exit = 2
			break
		} else {
			if isJsonFile(path) {
				//readData(path)
				if check {
					oldInstance = newInstance
					newInstance = readData(path)
					printProduct(oldInstance)
					printProduct(newInstance)
					printResult(oldInstance, newInstance)
				} else {
					newInstance = readData(path)
					printProduct(newInstance)
					check = true
				}

			} else {
				fmt.Println("Invalid file! Enter path again: ")
			}
		}
	}

}
