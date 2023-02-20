package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/byebyebruce/pkg/chatgpt"
	"github.com/byebyebruce/pkg/util"
	"github.com/fatih/color"
)

var CHATGPT_KEY string

func main() {
	flag.Parse()
	if key := os.Getenv("CHATGPT_KEY"); len(key) > 0 {
		CHATGPT_KEY = key
	}
	if len(CHATGPT_KEY) == 0 {
		fmt.Println("set env CHATGPT_KEY=YOUR_API_KEY")
		os.Exit(1)
	}

	const (
		welcomeTip = `Welcome to ChatGPT! Ask your question.`
		you        = `Question:`
		chatGPT    = `Answer:`
	)
	client := chatgpt.New(CHATGPT_KEY, "", 0)
	fmt.Println(color.CyanString(welcomeTip))

	for {
		fmt.Println()
		fmt.Println(color.GreenString(you))

		q := ""
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			q = strings.TrimSpace(scanner.Text())
			if len(q) > 0 {
				break
			}
		}

		fmt.Println(color.YellowString(chatGPT))
		answer, err := util.AsyncTaskAndShowLoadingBar("Wait", func() (string, error) {
			return client.ChatWithContext(q)
		})
		if err != nil {
			fmt.Println(color.RedString("ERROR:%s", err.Error()))
		} else {
			fmt.Println(answer)
		}
	}
}
