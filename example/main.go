package main

import (
	"fmt"

	"github.com/xm-chentl/gochatgpt"
)

func main() {
	c, err := gochatgpt.NewClient(gochatgpt.Config{
		API:          "https://api.openai.com/v1",
		APIKey:       "", // openai api key
		Organization: "", // openai organization
	})
	if err != nil {
		panic(err)
	}

	// res, err := c.Models()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("model_id: ", res.Data[0].ID)
	// modelRes, err := c.Model(res.Data[0].ID)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(modelRes)

	// chatRes, err := c.Chat("能帮我画一张图吗?")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(chatRes)

	img := c.GetImage()
	imgRes, err := img.Create(gochatgpt.RequestCreateImage{
		Prompt: "帮我画一只中国龙",
		Number: 1,
		Size:   gochatgpt.X512,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("img >>> ", imgRes)
}
