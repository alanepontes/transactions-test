package main

import (
	"encoding/csv"
	"io/ioutil"
	"os"
	"strings"
	"fmt"
	"github.com/shopspring/decimal"
)

const fatorOutput float64 = 100

func FormatToMoney(value int) decimal.Decimal {
	money := decimal.New(int64(value), -2)
    return money
}

func formatToOutput(value decimal.Decimal) string {
	output := value.Mul(decimal.NewFromFloat(fatorOutput))
    return output.StringFixed(0)
}

func CSVnify(accountId int, value decimal.Decimal) string {
	return fmt.Sprintf("%d,%s",accountId, formatToOutput(value))
}

func HandleArguments(args []string) (error, []string) {
	if len(args[1:]) != 2 {
        return &CustomError{args[0], "Required syntax: nameFile01.csv nameFile02.csv\n"}, nil
	}
    return nil, args[1:]
}

func ProcessCSV(nameFile string) (error, string) {
    _, err := os.Stat(nameFile)
	if err != nil {
		return err, ""
	} else {
		data, err2 := ioutil.ReadFile(nameFile)
		if err2 != nil {
			return err2, ""
		} else {
			return nil, string(data)
		}
	}
}

func HandleCSV(data string) (error, [][]string){
    reader := csv.NewReader(strings.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		return err, [][]string{}
	} else {
		return nil, records
	}
}

