package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	// ENV_LOCAL : carregando variáveis de ambiente à partir do arquivo local ENV_LOCAL.
	err := godotenv.Load("ENV_LOCAL")

	if err != nil {
		log.Fatal("Error => ", err)
	}

	logpath := os.Getenv("LOG_PATH")

	// TIME : obtendo hora e data no formato YYYY-MM-DD hh:mm:ss.
	currentTime := time.Now()
	timestamp := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println(timestamp)

	// READ : buffer e laço para streaming.
	file, err := os.Open(logpath)
	if err != nil {
		log.Fatal("Error => ", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var line []string
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}

	// FOR : Loop que que permanecerá em execução enquanto o arquivo de log existir.
	for _, err := os.Stat(logpath); !os.IsNotExist(err); {
		for _, eachline := range line {
			reqBody, err := json.Marshal(map[string]string{
				"Timestamp": timestamp,
				"Event":     eachline,
			})
			if err != nil {
				log.Fatal("Error => ", err)
			}
			res, err := http.Post("http://localhost:9200/golang/golang", "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				log.Fatal("Error => ", err)
			}
			fmt.Println(res)

		}
	}
}
