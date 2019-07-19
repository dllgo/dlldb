package main

import (
	"encoding/json"
	"fmt"

	"github.com/dllgo/dlldb"
	"github.com/dllgo/dlldb/example/model"
)

func main() {
	fmt.Println("123")
	iserv := dlldb.NewServices()

	var result []model.Menu
	page, err := iserv.List(&result, "id desc", 10, 0, "")
	if err != nil {

		//not found
		if err != nil {
			// has error
			fmt.Println(err.Error())
		}

	}
	data := map[string]interface{}{}
	data["lists"] = result
	data["page"] = page
	b4, err := json.Marshal(data)
	if err != nil {
		fmt.Println("%v\n", err)
	}
	fmt.Println(string(b4))
}
