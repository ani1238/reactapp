package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
	"time"
    //"encoding/json"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"

    "github.com/aws/aws-lambda-go/lambda"
)

type JSONResponse struct {
    Value1 string `json:"key1"`
    Value2 string `json:"key2"`
}


func handler() (JSONResponse,error) {
    url := "https://search-demo3-hr7kns5ajpuethjdyq4zcmvkni.us-east-2.es.amazonaws.com/my-index/_delete_by_query"
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
    
    body1, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    
    

    jsonResponse := JSONResponse{
        Value1: "Test value 1",
        Value2: string(body1),
    }

    fmt.Printf("The struct returned before marshalling\n\n")
    fmt.Printf("%+v\n\n\n\n", jsonResponse)
    return jsonResponse,nil
}

func main() {
    
    lambda.Start(handler)
}
