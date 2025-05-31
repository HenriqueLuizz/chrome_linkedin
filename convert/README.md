# Conversor LinkedIn JSON para CSV/XLSX

Este programa converte arquivos JSON exportados do LinkedIn para os formatos CSV ou XLSX, facilitando a análise dos dados em planilhas. Ele pode processar um único arquivo ou todos os arquivos `.json` de um diretório.

## Pré-requisitos
- Go 1.19 ou superior (apenas para compilar ou rodar via `go run`)
- Binários prontos para Mac, Windows e Linux (veja abaixo)

## Instalação

### Usando binário pronto
- **Mac:** use o arquivo `convert_mac`
- **Windows:** use o arquivo `convert.exe`
- **Linux:** use o arquivo `convert_linux`

Coloque o binário desejado em uma pasta de sua preferência e dê permissão de execução (Linux/Mac):
```sh
chmod +x convert_mac
```

### Compilando manualmente
Se preferir compilar:
```sh
cd convert
# Para Mac
GOOS=darwin GOARCH=amd64 go build -o convert_mac convert.go
# Para Windows
GOOS=windows GOARCH=amd64 go build -o convert.exe convert.go
# Para Linux
GOOS=linux GOARCH=amd64 go build -o convert_linux convert.go
```

## Como usar

### Parâmetros
- `-input` : Caminho do arquivo JSON ou diretório contendo arquivos `.json`
- `-format`: Formato de saída: `csv` ou `xlsx`

### Exemplos

#### Converter um arquivo JSON para CSV
```sh
./convert_mac -input=linkedin_posts_2025-05-30.json -format=csv
```

#### Converter todos os arquivos JSON de um diretório para XLSX
```sh
./convert_mac -input=. -format=xlsx
```

#### No Windows
```sh
convert.exe -input=linkedin_posts_2025-05-30.json -format=csv
```

#### No Linux
```sh
./convert_linux -input=linkedin_posts_2025-05-30.json -format=xlsx
```

## Sobre a coluna `post_date`
O programa calcula automaticamente a data exata do post (coluna `post_date`) usando o campo `date` (ex: "Há 2 dias ...") e o campo `timestamp` (data/hora da extração). O resultado é uma data/hora precisa do post, no mesmo formato do timestamp.

## Licença
GPL-3.0 