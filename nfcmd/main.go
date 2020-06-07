package main

import (
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ffmiyo/gummary"
	json "github.com/json-iterator/go"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	url := request.QueryStringParameters["q"]
	sel := "p"
	rawText, err := gummary.Scrape(sel, url)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 503,
			Body:       "",
		}, err
	}
	ranked := gummary.RankText(rawText)
	summary := strings.Join(ranked, " ")
	resp, err := json.Marshal(map[string]interface{}{
		"item": summary,
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 503,
			Body:       "",
		}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
		Body:       string(resp),
	}, nil
}

func main() {
	lambda.Start(handler)
}
