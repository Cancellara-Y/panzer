package panzer

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type CheckData struct {
	GroupName string
	OwnerName string
	RobotName string
}

func ReadCsv1(fileName string) []*CheckData {
	//准备读取文件

	fs, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("can not open the file, err is %+v \n", err)
	}
	defer fs.Close()

	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件

	var list []*CheckData
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			fmt.Printf("can not read, err is %+v\n", err)
		}

		if err == io.EOF {
			break
		}

		if len(row) < 3 {
			fmt.Println("格式错误：", row)
			continue
		}
		fmt.Println("first: ", row[2])

		ss := strings.Split(row[2], "助理")
		for _, v := range ss {
			fmt.Printf("#%v#\n", v)
		}
		fmt.Println("second: ", strings.Trim(strings.Trim(row[2], " "), ";"))

		list = append(list, &CheckData{
			GroupName: row[0],
			OwnerName: row[1],
			RobotName: strings.Trim(row[2], ";"),
		})

	}
	return list
}

func ReadFile(fileName string) []string {
	//准备读取文件

	fs, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("can not open the file, err is %+v \n", err)
	}
	defer fs.Close()

	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件

	var list []string
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			fmt.Printf("can not read, err is %+v\n", err)
		}

		if err == io.EOF {
			break
		}

		if len(row) != 1 {
			fmt.Println("格式错误：", row)
			continue
		}

		fmt.Println(row[0])

	}
	return list
}

func ReadLine(fileName string) []string {

	f, err := os.Open(fileName)
	if err != nil {
		return []string{}
	}
	buf := bufio.NewReader(f)
	var result []string
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil && err != io.EOF {
			fmt.Printf("can not read, err is %+v\n", err)
		}

		if err == io.EOF {
			break
		}
		result = append(result, line)
	}
	return result
}

func goFunc() {
	arr := make(map[int][]int)
	for i := 1; i < 4; i++ {
		arr[i] = []int{i * 10, i*10 + 1, i*10 + 2, i*10 + 3}
	}

	pr := func(a int, list []int) {
		for _, v := range list {
			fmt.Println(a, v)
		}
	}

	var waitGroup sync.WaitGroup

	for k, list := range arr {

		waitGroup.Add(1)
		a := k
		b := list
		go func() {
			pr(a, b)
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

}
