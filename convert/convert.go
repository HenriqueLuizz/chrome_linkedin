package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/tealeg/xlsx"
)

type Post struct {
	Content     string `json:"content" csv:"content"`
	Date        string `json:"date" csv:"date"`
	Description string `json:"description" csv:"description"`
	Link        string `json:"link" csv:"link"`
	Name        string `json:"name" csv:"name"`
	Reactions   string `json:"reactions" csv:"reactions"`
	Reposts     string `json:"reposts" csv:"reposts"`
	Timestamp   string `json:"timestamp" csv:"timestamp"`
	PostDate    string `csv:"post_date"`
	Owner       string `csv:"owner"`
}

func main() {
	inputPath := flag.String("input", "", "Arquivo ou diretório de entrada (JSON)")
	outputFormat := flag.String("format", "csv", "Formato de saída: csv ou xlsx")
	flag.Parse()

	if *inputPath == "" {
		fmt.Println("Por favor, informe o caminho do arquivo ou diretório com a flag -input")
		os.Exit(1)
	}

	info, err := os.Stat(*inputPath)
	if err != nil {
		fmt.Printf("Erro ao acessar o caminho: %v\n", err)
		os.Exit(1)
	}

	if info.IsDir() {
		files, err := ioutil.ReadDir(*inputPath)
		if err != nil {
			fmt.Printf("Erro ao ler diretório: %v\n", err)
			os.Exit(1)
		}
		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".json") {
				convertFile(filepath.Join(*inputPath, file.Name()), *outputFormat)
			}
		}
	} else {
		convertFile(*inputPath, *outputFormat)
	}
}

func parsePostDate(dateStr, timestamp string) string {
	// timestamp: "2025-05-30 19:23:10"
	layout := "2006-01-02 15:04:05"
	refTime, err := time.Parse(layout, timestamp)
	if err != nil {
		return ""
	}
	// Regex para pegar o número e unidade
	re := regexp.MustCompile(`Há (\d+) (minuto|minutos|hora|horas|dia|dias|semana|semanas|mês|meses|ano|anos)`) //
	match := re.FindStringSubmatch(dateStr)
	if len(match) < 3 {
		return ""
	}
	value, _ := strconv.Atoi(match[1])
	unit := match[2]
	delta := time.Duration(0)
	switch unit {
	case "minuto", "minutos":
		delta = time.Duration(value) * time.Minute
	case "hora", "horas":
		delta = time.Duration(value) * time.Hour
	case "dia", "dias":
		delta = time.Duration(value) * 24 * time.Hour
	case "semana", "semanas":
		delta = time.Duration(value) * 7 * 24 * time.Hour
	case "mês", "meses":
		// Aproximação: 30 dias por mês
		delta = time.Duration(value) * 30 * 24 * time.Hour
	case "ano", "anos":
		// Aproximação: 365 dias por ano
		delta = time.Duration(value) * 365 * 24 * time.Hour
	}
	postTime := refTime.Add(-delta)
	return postTime.Format(layout)
}

// Função para tratar duplicação no campo name
func cleanName(name string) string {
	if len(name) == 0 {
		return name
	}
	// Se o nome tem pelo menos 2 caracteres e a segunda metade é igual à primeira
	mid := len(name) / 2
	if len(name) >= 2 && name[:mid] == name[mid:] {
		return name[:mid]
	}
	return name
}

// Função para encontrar o nome que mais se repete
func findMostFrequentName(posts []Post) string {
	nameCount := make(map[string]int)
	for _, post := range posts {
		nameCount[post.Name]++
	}

	var mostFrequentName string
	maxCount := 0
	for name, count := range nameCount {
		if count > maxCount {
			maxCount = count
			mostFrequentName = name
		}
	}
	return mostFrequentName
}

func convertFile(jsonPath, format string) {
	data, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		fmt.Printf("Erro ao ler arquivo %s: %v\n", jsonPath, err)
		return
	}
	var posts []Post
	if err := json.Unmarshal(data, &posts); err != nil {
		fmt.Printf("Erro ao fazer parse do JSON em %s: %v\n", jsonPath, err)
		return
	}
	// Preencher o campo PostDate e limpar o campo Name
	for i := range posts {
		posts[i].PostDate = parsePostDate(posts[i].Date, posts[i].Timestamp)
		posts[i].Name = cleanName(posts[i].Name)
	}

	// Encontrar o nome mais frequente para a coluna Owner
	mostFrequentName := findMostFrequentName(posts)
	for i := range posts {
		posts[i].Owner = mostFrequentName
	}

	baseName := strings.TrimSuffix(jsonPath, filepath.Ext(jsonPath))
	if format == "csv" {
		csvFile, err := os.Create(baseName + ".csv")
		if err != nil {
			fmt.Printf("Erro ao criar arquivo CSV: %v\n", err)
			return
		}
		defer csvFile.Close()
		if err := gocsv.MarshalFile(&posts, csvFile); err != nil {
			fmt.Printf("Erro ao exportar CSV: %v\n", err)
		}
		fmt.Printf("Arquivo CSV gerado: %s.csv\n", baseName)
	} else if format == "xlsx" {
		file := xlsx.NewFile()
		sheet, err := file.AddSheet("Posts")
		if err != nil {
			fmt.Printf("Erro ao criar planilha XLSX: %v\n", err)
			return
		}
		headers := []string{"content", "date", "description", "link", "name", "reactions", "reposts", "timestamp", "post_date", "owner"}
		headerRow := sheet.AddRow()
		for _, h := range headers {
			headerRow.AddCell().Value = h
		}
		for _, post := range posts {
			row := sheet.AddRow()
			row.AddCell().Value = post.Content
			row.AddCell().Value = post.Date
			row.AddCell().Value = post.Description
			row.AddCell().Value = post.Link
			row.AddCell().Value = post.Name
			row.AddCell().Value = post.Reactions
			row.AddCell().Value = post.Reposts
			row.AddCell().Value = post.Timestamp
			row.AddCell().Value = post.PostDate
			row.AddCell().Value = post.Owner
		}
		if err := file.Save(baseName + ".xlsx"); err != nil {
			fmt.Printf("Erro ao salvar XLSX: %v\n", err)
		} else {
			fmt.Printf("Arquivo XLSX gerado: %s.xlsx\n", baseName)
		}
	} else {
		fmt.Println("Formato não suportado. Use 'csv' ou 'xlsx'.")
	}
}
