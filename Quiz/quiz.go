package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func csvReader(name string) {
	//open the file
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		log.Fatalln("An error encountered ::", err)
		return
	}

	//init the csv reader
	reader := csv.NewReader(file)

	//read the records of the csv file
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("An error encountered ::", err)
		}
		fmt.Printf("Questio: %s Answer: %s\n", record[0], record[1])
	}
}
func main() {
	csvReader("problems.csv")
}
