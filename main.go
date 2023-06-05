package main

import (
	"promptProgrammingGo/practices"
	"context"
	"fmt"
	"os"
	openai "github.com/sashabaranov/go-openai"
)

func executeFunction(prompt string) {
	token := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}

func main(){
	//executeFunction(practices.ParsePractice(practices.OnlyDelimiters()))
	//executeFunction(practices.ParsePractice(practices.StructuredOutput()))
	//executeFunction(practices.ParsePractice(practices.AreConditionsSatisfied()))
	//executeFunction(practices.ParsePractice(practices.FewShotPrompting()))
	//executeFunction(practices.ParsePractice(practices.SpecifySteps()))
	//executeFunction(practices.ParsePractice(practices.WorkOwnSolution()))
	//executeFunction(practices.ParsePractice(practices.Hallucination()))

	// review
	//executeFunction(practices.ParsePractice(practices.IdentifyEmotions()))
	executeFunction(practices.ParsePractice(practices.ExtractInformation()))

	
}