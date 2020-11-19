package main

import (
    "fmt"
    "net/http"
    "strings"
	"time"
    //"encoding/json"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"

    "github.com/aws/aws-lambda-go/lambda"
)



func handler() {
    url := "https://search-demo3-hr7kns5ajpuethjdyq4zcmvkni.us-east-2.es.amazonaws.com/modify/_delete_by_query"
    query := `{
      "query": {
        "match_all": {
        
        }
      }
    }`
    
    region := "us-east-2"
    service := "es"
    
    credentials := credentials.NewEnvCredentials()
	signer := v4.NewSigner(credentials)

    client := &http.Client{}

    body := strings.NewReader(query)
        
    req, err := http.NewRequest("POST", url, body)
    
    if err != nil {
        fmt.Println(err)
    }
    
    req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Allow-Origin","*")

    signer.Sign(req, body, service, region, time.Now())

    resp, err := client.Do(req)
    if err!=nil{
        fmt.Println(err)
    }
    fmt.Println(resp)

    
    
}

func main() {
    
    lambda.Start(handler)
}

