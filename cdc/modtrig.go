package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"

	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, e events.DynamoDBEvent) {

	domain := "https://search-demo3-hr7kns5ajpuethjdyq4zcmvkni.us-east-2.es.amazonaws.com" // e.g. https://my-domain.region.es.amazonaws.com
	index := "modify"
	//id := "1"

	region := "us-east-2" // e.g. us-east-1
	service := "es"
	count := 0

	for _, record := range e.Records {
		fmt.Println(record)

		endpoint := domain + "/" + index + "/" + "_doc" + "/" + record.EventID

		if record.EventName == "MODIFY" || record.EventName == "INSERT" {
			count++
			month := record.Change.NewImage["Month"].String()
			cupcakes := record.Change.NewImage["Cupcakes"].String()
			ut := record.Change.NewImage["update_time"].String()
			document := `{ "Month": "`+month+`", "Cupcakes" : "`+cupcakes+`", "update_time" :"`+ut+`" }`
			//fmt.Println(document.String())
			
			
			body := strings.NewReader(document)

			credentials := credentials.NewEnvCredentials()
			signer := v4.NewSigner(credentials)

			// An HTTP client for sending the request
			client := &http.Client{}
			req, err := http.NewRequest(http.MethodPut, endpoint, body)
			if err != nil {
				fmt.Print(err)
			}

			// You can probably infer Content-Type programmatically, but here, we just say that it's JSON
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Access-Control-Allow-Origin","*")


			// Sign the request, send it, and print the response
			signer.Sign(req, body, service, region, time.Now())
			resp, err := client.Do(req)
			if err != nil {
				fmt.Print(err)
			}
			fmt.Print(resp)
			
		}

		// Print new values for attributes of type String
		/*for name, value := range record.Change.NewImage {
			if value.DataType() == events.DataTypeString {
				fmt.Printf("Attribute name: %s, value: %s\n", name, value.String())
			}
		}*/
	}

	// Sample JSON document to be included as the request body

	// Get credentials from environment variables and create the AWS Signature Version 4 signer

	// Form the HTTP request

}

func main() {
	lambda.Start(handleRequest)
}

//https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-aws-integrations.html#es-aws-integrations-dynamodb-es

//https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-request-signing.html#es-request-signing-go
