package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Record struct {
	Name       string
	Suname     string
	Number     string
	LastAccess string
}

var data = []Record{}
var index map[string]int

const CSVFILE string = "/home/dzrise/PhpstormProjects/education/go/ChapterThree/input.csv"

func readCsvFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}

	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		temp := Record{
			Name:       line[0],
			Suname:     line[1],
			Number:     line[2],
			LastAccess: line[3],
		}
		data = append(data, temp)
	}

	return nil
}

func saveCscFile(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer csvfile.Close()

	csvWriter := csv.NewWriter(csvfile)
	csvWriter.Comma = ','
	for _, row := range data {
		temp := []string{row.Name, row.Suname, row.Number, row.LastAccess}
		_ = csvWriter.Write(temp)
	}
	csvWriter.Flush()
	return nil
}

func search(key string) *Record {
	println(index)
	i, ok := index[key]
	if !ok {
		return nil
	}

	data[i].LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
	for i, v := range data {
		if v.Number == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}

	fmt.Println()
}

func createIndex() error {
	index = make(map[string]int)
	fmt.Println(data)
	for i, k := range data {
		key := k.Number
		index[key] = i
	}
	fmt.Println(index)
	return nil
}

func matchTel(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func initS(s1 string, s2 string, s3 string) *Record {
	if s3 == "" || s2 == "" {
		return nil
	}
	LastAccess := strconv.FormatInt(time.Now().Unix(), 10)
	return &Record{Name: s1, Suname: s2, Number: s3, LastAccess: LastAccess}

}

func deleteRecord(key string) error {
	i, ok := index[key]
	if !ok {
		return fmt.Errorf("%s not found", key)
	}

	data = append(data[:i], data[i+1:]...)
	delete(index, key)
	err := saveCscFile(CSVFILE)
	if err != nil {
		return err
	}

	return nil
}

func insert(record *Record) error {
	_, ok := index[(*record).Number]
	if ok {
		return fmt.Errorf("Number %s already exists", record.Number)
	}

	data = append(data, *record)
	_ = createIndex()
	err := saveCscFile(CSVFILE)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Usage: insert|delete|search|list <argumnts>")
		return
	}

	_, err := os.Stat(CSVFILE)

	if err != nil {
		fmt.Println("Createing", CSVFILE)
		f, err := os.Create(CSVFILE)
		if err != nil {
			f.Close()
			fmt.Println(err)
			return
		}
		f.Close()
	}

	fileinfo, err := os.Stat(CSVFILE)
	mode := fileinfo.Mode()

	if !mode.IsRegular() {
		fmt.Println(CSVFILE, "is not a regular file")
		return
	}

	err = readCsvFile(CSVFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = createIndex()
	if err != nil {
		fmt.Println("Can't create index")
		return
	}

	arguments := os.Args
	switch arguments[1] {
	case "insert":
		if len(os.Args) != 5 {
			fmt.Println("Usage: insert <name> <suname> <number>")
			return
		}
		t := strings.ReplaceAll(arguments[4], "-", "")
		if !matchTel(t) {
			fmt.Println("Invalid number: ", t)
			return
		}

		temp := initS(arguments[2], arguments[3], arguments[4])
		if temp != nil {
			err := insert(temp)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		fmt.Println("Inserted")

	case "delete":
		if len(arguments) != 3 {
			fmt.Println("Usage: delete number")
			return
		}
		t := strings.ReplaceAll(arguments[2], "-", "")
		if !matchTel(t) {
			fmt.Println("Invalid number: ", t)
			return
		}

		err := deleteRecord(t)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Deleted")
	case "search":
		if len(os.Args) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}

		result := search(os.Args[2])

		if result == nil {
			fmt.Println("Entery not found: ", os.Args[2])
			return
		}

		fmt.Println(*result)

	case "list":

		list()

	default:
		fmt.Println("Not a valid option")
	}

}
