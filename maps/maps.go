package main

import "fmt"

func main(){
    dict :=make(map[string]string)
    dict["go"] = "Golang"
    dict["cs"] = "CSharp"
    dict["rb"] = "Ruby"
    dict["py"] = "Python"
    dict["js"] = "JavaScript"
    for k,v := range dict {
	fmt.Printf("Key: %s Value %s\n", k,v)
    }
    
}