# LinkedIn Post Extractor

🚀 **LinkedIn Post Extractor** é uma extensão para o Google Chrome que permite capturar posts do feed do LinkedIn e exportá-los em **JSON** para análise posterior.

## 📌 Funcionalidades

- 📥 Captura posts do feed do LinkedIn automaticamente.
- 💾 Exportação dos dados em **JSON** ou **CSV**.
- 📅 Armazena a data e hora da extração.
- 🛠️ Código **open-source** sob a licença **GPL-3.0**.

## 🛠️ Instalação

### 1️⃣ Instalando a extensão manualmente:

1. **Baixe o código-fonte** deste repositório clicando em **Code > Download ZIP** ou clonando via Git:

    ```sh
    git clone https://github.com/HenriqueLuizz/chrome_linkedin.git
    ```

2. Extraia os arquivos do ZIP para uma pasta local.
3. Acesse o Chrome e ative o modo desenvolvedor:
    - Vá até chrome://extensions/
    - Ative o Modo do desenvolvedor no canto superior direito.
4. Carregue a extensão:
    - Clique em Carregar sem compactação e selecione a pasta do projeto.

Pronto! A extensão está instalada no seu Chrome. 🎉

## 📋 Como Usar

1. Acesse o LinkedIn e vá até a página de perfil de um usuário.
2. Clique no botão "Capturar Posts" na extensão.
3. Escolha o formato de exportação: JSON.
4. O arquivo será baixado automaticamente para o seu computador.

## 📂 Estrutura do Projeto

```sh
📂 linkedin-post-extractor
 ├── 📁 icons/         # Ícones da extensão
 ├── 📁 scripts/       # Código JavaScript principal
 ├── 📄 manifest.json  # Configuração da extensão
 ├── 📄 popup.html     # Interface do usuário
 ├── 📄 popup.js       # Lógica do botão de captura
 ├── 📄 styles.css     # Estilos do popup
 └── 📄 README.md      # Este arquivo
````

## 📜 Licença

Este projeto é licenciado sob a GNU General Public License v3.0.
Isso significa que qualquer projeto que utilizar este código também deverá ser open-source.

📜 Leia mais sobre a licença em GNU GPL-3.0.

## 🤝 Contribuindo

Contribuições são bem-vindas! Para contribuir:

1. Faça um fork do projeto.
2. Crie uma branch (git checkout -b minha-feature).
3. Faça um commit (git commit -m 'Adiciona nova funcionalidade').
4. Faça um push (git push origin minha-feature).
5. Abra um Pull Request.
