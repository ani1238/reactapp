package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-lambda-go/lambda"
)

var dynamo *dynamodb.DynamoDB

type Cupcakes struct {
	ID    int
	Month string
	Cc    string
}

type CupDate struct {
	ID int
	ud string
}

const Table_Name = "cupcakes"

func init() {
	dynamo = connectDynamo()
}

func connectDynamo() (db *dynamodb.DynamoDB) {
	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	})))
}

func CreateTable() {
	_, err := dynamo.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("ID"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
		TableName: aws.String(Table_Name),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
	}
}

func PutItem(cupcakes Cupcakes) {
	_, err := dynamo.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String(strconv.Itoa(cupcakes.ID)),
			},
			"Month": {
				S: aws.String(cupcakes.Month),
			},
			"Cupcakes": {
				S: aws.String(cupcakes.Cc),
			},
		},
		TableName: aws.String(Table_Name),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
	}
}

func UpdateItem(cupcakes Cupcakes) {
	_, err := dynamo.UpdateItem(&dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string{
			"#M": aws.String("Month"),
			"#C": aws.String("Cupcakes"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":Month": {
				S: aws.String(cupcakes.Month),
			},
			":Cupcakes": {
				S: aws.String(cupcakes.Cc),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String(strconv.Itoa(cupcakes.ID)),
			},
		},
		TableName:        aws.String(Table_Name),
		UpdateExpression: aws.String("SET #M = :Month, #C = :Cupcakes"),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
	}
}

func UpdateCurrentDateItem(cupcakes CupDate) {
	_, err := dynamo.UpdateItem(&dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string{
			"#u": aws.String("update_time"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":update_time": {
				S: aws.String(cupcakes.ud),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String(strconv.Itoa(cupcakes.ID)),
			},
		},
		TableName:        aws.String(Table_Name),
		UpdateExpression: aws.String("SET #u = :update_time"),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
	}
	fmt.Printf("Updated id %d\n", cupcakes.ID)
}

func DeleteItem(id int) {
	_, err := dynamo.DeleteItem(&dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String(strconv.Itoa(id)),
			},
		},
		TableName: aws.String(Table_Name),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
	}

}

func GetItem(id int) (cupcake Cupcakes) {
	result, err := dynamo.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String(strconv.Itoa(id)),
			},
		},
		TableName: aws.String(Table_Name),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &cupcake)
	if err != nil {
		panic(err)
	}

	return cupcake

}

func createWholeTable() {
	csvfile, err := os.Open("multiTimeline.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records

	//CreateTable()

	var i = 1
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var m = record[0]
		var c = record[1]
		var cc Cupcakes = Cupcakes{
			ID:    i,
			Month: m,
			Cc:    c,
		}
		i++
		PutItem(cc)
		fmt.Printf("Month: %s Cupcake %s\n", record[0], record[1])
	}
}

func main() {
	//createWholeTable()

	//demoupdate
	lambda.Start(handler)

}

func handler() {
	for i := 0; i < 100; i++ {
		var r = rand.Intn(201) + 2
		var cur_time = time.Now()
		var cc CupDate = CupDate{
			ID: r,
			ud: cur_time.String(),
		}
		UpdateCurrentDateItem(cc)
	}
}
