package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type singleDatabase struct {
	capitals map[string]int
}

func (db *singleDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

var instance *singleDatabase
var once sync.Once

// both init and sync.Once are thread-safe
// but only sync.Once is lazy
// init() — we could, but it's not lazy
func GetSingletonDatabase() *singleDatabase {
	once.Do(func() {
		caps, e := readData("./capitals.txt")
		db := singleDatabase{caps}
		if e == nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}

// func GetTotalPopulation(cities []string) int {
// 	result := 0
// 	for _, city := range cities {
// 		result += GetSingletonDatabase().GetPopulation(city)
// 	}//        1. 라이브 데이터 베이스에 직접 의존하고 있어 좋지 않다.
// 	return result
// }

// 2. Database를 인터페이스로
type Database interface {
	GetPopulation(city string) int
}

// 3. DB를 인자로 넘긴다.
func GetTotalPopulation(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}

// 4. DummyDatabase를 만든다.
type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3,
		}
	}
	return d.dummyData[name]
}

func main() {
	// db := GetSingletonDatabase()
	// pop := db.GetPopulation("Seoul")
	// fmt.Println("Pop of Seoul = ", pop)
	// cities := []string{"Seoul", "Mexico City"}
	// tp := GetTotalPopulation(cities)
	// // Huge problem - test with live data if database changes, test will be fail
	// ok := tp == (17500000 + 17400000)
	names := []string{"alpha", "gamma"}
	tp := GetTotalPopulation(&DummyDatabase{}, names)
	fmt.Println(tp == 4)
}

func readData(path string) (map[string]int, error) {
	abs, _ := filepath.Abs(path)
	fmt.Println(abs)
	file, err := os.Open(abs)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}
