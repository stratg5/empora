package main

import (
	"empora/address"
	"empora/entities"
	"flag"

	"github.com/smartystreets/smartystreets-go-sdk/wireup"
)

func main() {
	var csvPath = flag.String("csvPath", "./addresses.csv", "the path to the csv, defaults to ./addresses.csv")
	var id = flag.String("apiID", "", "the api ID")
	var token = flag.String("apiToken", "", "the api token")

	println(*csvPath)
	println(*id)
	println(*token)
	// parse the csv with fields Street, City and Zip Code
	// account for missing or incorrect fields
	// gather the 3 fields into an object
	// Send the fields to the endpoint for verification
	// Generate the output as a csv or as console output (flag for options)

	// TODO check for the id and token in entities file, if not present then prompt
	
	client := address.NewClient(wireup.BuildUSStreetAPIClient(wireup.SecretKeyCredential(entities.ID, entities.Token)))

}

// func main() {

	// 	Create a command-line program that validates a US address against the following API and
	// outputs either the corrected address or Invalid Address

	// https://www.smarty.com/products/us-address-verification

	// 1. The tool must run in a console window/terminal.

	// 2. The input should be piped in or read from a file: e.g. cat file.csv | node
	// program.js or ruby program.rb file.csv

	// 3. The input format is CSV with the following fields: Street, City, and Zip Code

	// 4. The output should include the original address and the corrected address joined with ,
	// and separated by a -> (see below).

	// 5. The free trial of the API provides a limited set of API checks. Please do not pay for
	// additional checks. A testing suite can quickly exhaust all 1000 checks, but there are
	// several testing strategies that can help mitigate this issue - one potential strategy is stub
	// or mock.
// }