# Conversor LinkedIn JSON para CSV/XLSX

Este programa converte arquivos JSON exportados do LinkedIn para os formatos CSV ou XLSX, facilitando a análise dos dados em planilhas. Ele pode processar um único arquivo ou todos os arquivos `.json` de um diretório.

## Pré-requisitos
- Go 1.19 ou superior (apenas para compilar ou rodar via `go run`)
- Binários prontos para Mac, Windows e Linux (veja abaixo)

## Instalação e Uso Rápido

### 1. Baixe o binário do GitHub
Acesse a página de [Releases do projeto](https://github.com/HenriqueLuizz/chrome_linkedin/releases) e baixe o binário correspondente ao seu sistema operacional:
- **Mac:** `convert_darwin_amd64`
- **Windows:** `convert_windows_amd64.exe`
- **Linux:** `convert_linux_amd64`

### 2. Extraia e dê permissão de execução (Linux/Mac)
Se necessário, torne o arquivo executável:
```sh
chmod +x convert_darwin_amd64
chmod +x convert_linux_amd64
```

### 3. Execute o programa
#### **Mac/Linux**
```sh
./convert_darwin_amd64 -input=linkedin_posts_2025-05-30.json -format=csv
```

#### **Windows**
Para abrir o CMD no Windows e acessar a pasta onde está o executável, siga os passos abaixo:

1. Pressione `Win + R` para abrir a janela "Executar".
2. Digite `cmd` e pressione `Enter` para abrir o Prompt de Comando.
3. No Prompt de Comando, use o comando `cd` (change directory) para navegar até a pasta onde o executável está localizado. Por exemplo, se o executável estiver na pasta `C:\Users\SeuUsuario\Downloads`, você deve digitar:
   ```sh
   cd C:\Users\SeuUsuario\Downloads
   ```
4. Agora, você pode executar o programa conforme descrito acima.

5. No terminal (cmd ou PowerShell) na pasta do binário e execute:
```sh
convert_windows_amd64.exe -input=linkedin_posts_2025-05-30.json -format=xlsx
```
ou pode ser passado apenas o caminha da pasta
```sh
convert_windows_amd64.exe -input=C:\Users\SeuUsuario\Downloads\arquivos_json\ -format=xlsx
```

### 4. Exemplos detalhados
#### Converter um arquivo JSON para CSV
```sh
./convert_darwin_amd64 -input=linkedin_posts_2025-05-30.json -format=csv
```

#### Converter todos os arquivos JSON de um diretório para XLSX
```sh
./convert_darwin_amd64 -input=. -format=xlsx
```

#### No Windows (um arquivo)
```sh
convert_windows_amd64.exe -input=linkedin_posts_2025-05-30.json -format=csv
```

#### No Linux (diretório)
```sh
./convert_linux_amd64 -input=. -format=xlsx
```

### 5. Saída
O programa irá gerar arquivos `.csv` ou `.xlsx` com o mesmo nome do arquivo de entrada, na mesma pasta.

## Sobre a coluna `post_date`
O programa calcula automaticamente a data exata do post (coluna `post_date`) usando o campo `date` (ex: "Há 2 dias ...") e o campo `timestamp` (data/hora da extração). O resultado é uma data/hora precisa do post, no mesmo formato do timestamp.

## Compilando manualmente (opcional)
Se preferir compilar:
```sh
cd convert
# Para Mac
go build -o convert_darwin_amd64 convert.go
# Para Windows
go build -o convert_windows_amd64.exe convert.go
# Para Linux
go build -o convert_linux_amd64 convert.go
```

## Licença
GPL-3.0 