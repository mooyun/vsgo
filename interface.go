package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"


)

type User struct {
	Id   int    `json:"id" xml:"id,attr"`
	Name string `json:"name" xml :"name,attr"`
}
type JsonParse string

func (p JsonParse) Serialize(v interface{}){
	if bytes,err:=json.Marshal(v);err!=nil{
		panic(err)
	}else{
		fmt.Printf("%s\n",bytes)
	}
}

type XmlParse string

func (p XmlParse) Serialize(v interface{}){
	if bytes,err:=xml.Marshal(v);err!=nil{
		panic(err)
	}else{
		fmt.Printf("%s",bytes)
	}
}



func main() {
	fmt.Println("接口")
	user:=User{
		Id: 101,
		Name: "小李",
	}

	 var p1 JsonParse
	// p1.Serialize(user)
	
	var p2 XmlParse
	// p2.Serialize(user)

	fmt.Println("接口之后")
	printAny(user,p1)
	fmt.Println()
	printAny(user,p2)
}

type TokenParse interface{
	Serialize(v interface{})
}

func printAny(v interface{}, p TokenParse) {
	p.Serialize(v)
}