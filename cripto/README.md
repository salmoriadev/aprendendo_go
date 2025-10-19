# Aprendendo Go: Pequena aplicação de certificação digital

Este projeto faz parte da minha jornada de estudos em Go. O objetivo principal é implementar um fluxo criptográfico básico de ICP, aplicando conceitos de design patterns e boas práticas de arquitetura de software.

O código demonstra como estruturar uma aplicação Go de forma modular, permitindo trocar algoritmos em tempo de execução sem alterar a lógica principal do sistema.

---

## Funcionalidades

O projeto simula um cenário real de assinatura digital de documentos, executando as etapas abaixo:

1. **Gera chaves** — cria pares de chaves RSA (pública e privada).  
2. **Cria certificados** — gera uma Autoridade Certificadora (CA) raiz e emite um certificado de usuário assinado por ela.  
3. **Prepara um documento** — cria um arquivo `.txt` com uma mensagem e o converte para `.pdf`.  
4. **Assina e verifica**:
   - Gera um resumo criptográfico (hash) do PDF.
   - Assina digitalmente esse resumo usando a chave privada do usuário.
   - Verifica a validade da assinatura usando a chave pública do usuário.

Todo o resultado (chaves, certificados, PDF e arquivos de assinatura) é salvo na pasta `arquivos_gerados/`.

---

## Arquitetura e boas práticas

Para manter o código limpo, testável e flexível, o projeto é dividido em pacotes com responsabilidades únicas:

- **`package criptografia`**
  - Contém apenas a lógica pura de criptografia (opera em `[]byte`).
  - Não tem dependência de I/O ou arquivos.
  - Funções típicas: `GerarChavePrivada`, `GerarCertificadoAutoassinado`, `ChavePrivadaParaPEM`, etc.

- **`package servicos`**
  - Responsável por todo o I/O (leitura/escrita de `.pem`, `.pdf`, `.txt`).
  - Orquestra o fluxo: chama `criptografia` para obter dados puros e persiste no disco.
  - Funções típicas: `ExecucaoChaves`, `ExecucaoCertificados`, `AssinarDocumentoPDF`.

- **`package main`**
  - Faz a "ligação" do sistema e injeta as dependências.

O núcleo do estudo foi desacoplar a lógica de negócio da implementação de algoritmos específicos — por isso a criação de interfaces no pacote `criptografia` que permitem trocar estratégias.

---

## Dependências

### Bibliotecas padrão do Go
- `crypto/rsa`  
- `crypto/x509`  
- `crypto/sha256`  
- `encoding/pem`  
- `os`  
- `log`  
- `fmt`

### Bibliotecas de terceiros
- `github.com/jung-kurt/gofpdf` — utilizada para criação do arquivo PDF.

> Observação: `go mod` gerencia dependências automaticamente; não é necessário instalar manualmente os pacotes.

---
1. **Clone o repositório:**
```bash
git clone https://github.com/salmoriadev/aprendendo_go.git

cd aprendendo_go/cripto
```

2. **Instale/garanta dependências (Go Modules):**
```bash
go mod tidy
```

3. **Execute o projeto:**
```bash
go run .
```

4. **Verifique os resultados:**

- O console exibirá as etapas de geração e validação da assinatura.

- A pasta arquivos_gerados/ conterá os arquivos gerados: .pem, .pdf, .txt e a assinatura.