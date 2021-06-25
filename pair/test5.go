package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct{
	Title string	`json:"title"`
	Year int		`json:"year"`
	Price int	`json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王",2000,10,[]string{"xingye","zhangbozhi"}}

	//结构体 -》json
	jsonStr,err := json.Marshal(movie)

	if err!= nil {
		fmt.Println("json marshal error",err)
		return
	}

	fmt.Printf("jsonStr = %s\n",jsonStr)



	// json ->结构体
	myMovie := Movie{}

	err = json.Unmarshal(jsonStr,&myMovie)
	if err != nil {
		fmt.Println("json marshal error",err)
		return
	}

	fmt.Println(myMovie)
}
