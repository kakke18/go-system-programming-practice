package main

import (
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	question1()
	question2()
	question3()
}

func question1() {
	file, err := os.Create("sample.txt")
	if err != nil {
		fmt.Print(err)
		return
	}
	defer func() {
		if err = file.Close(); err != nil {
			fmt.Print(err)
		}
	}()

	if _, err = fmt.Fprintf(file, "数字:%d 小数点:%f 文字列:%s", 1, 1.2, "hoge"); err != nil {
		fmt.Print(err)
		return
	}
}

func question2() {
	writer := csv.NewWriter(os.Stdout)
	if err := writer.Write([]string{"value1", "value2"}); err != nil {
		fmt.Print(err)
		return
	}
	if err := writer.Write([]string{"value3", "value4"}); err != nil {
		fmt.Print(err)
		return
	}
	writer.Flush()
}

func question3() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Content-Type", "application/json")

		// json化する元のデータ
		source := map[string]string{
			"hello": "world",
		}

		// ここ移行にコードを書く
		writer := gzip.NewWriter(w)
		defer func() {
			if err := writer.Close(); err != nil {
				fmt.Print(err)
			}
		}()
		writer.Header.Name = "test.txt"

		multiWriter := io.MultiWriter(writer, os.Stdout)
		encoder := json.NewEncoder(multiWriter)
		encoder.SetIndent("", "   ")
		if err := encoder.Encode(source); err != nil {
			fmt.Print(err)
			return
		}

		if err := writer.Flush(); err != nil {
			fmt.Print(err)
			return
		}
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Print(err)
		return
	}
}
