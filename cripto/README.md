# Aprendendo Go: Pequena aplicação de certificação digital

Este projeto faz parte da minha jornada de estudos em Go. O objetivo principal é implementar um fluxo criptográfico básico de ICP, aplicando conceitos de design patterns (principalmente Strategy) e boas práticas de arquitetura de software.

O código demonstra como estruturar uma aplicação Go de forma modular, permitindo trocar algoritmos em tempo de execução sem alterar a lógica principal do sistema. Cada pacote tem responsabilidades bem definidas e conta com documentação em português explicando o contexto de uso.

---

## Funcionalidades

O projeto simula um cenário real de assinatura digital de documentos, executando as etapas abaixo:

1. **Gera chaves** — cria pares de chaves RSA (pública e privada) e grava em disco (`chave_autoridade_*.pem`, `chave_usuario_*.pem`).  
2. **Cria certificados** — gera uma Autoridade Certificadora (CA) raiz e emite um certificado de usuário assinado por ela.  
3. **Prepara um documento** — cria um arquivo `.txt` com uma mensagem e o converte para `.pdf`.  
4. **Assina e verifica**:  
   - Gera um resumo criptográfico (hash SHA-256) do PDF.  
   - Assina digitalmente esse resumo usando a chave privada do usuário.  
   - Verifica a validade da assinatura usando a chave pública do usuário.

Todo o resultado (chaves, certificados, PDF e arquivos de assinatura) é salvo na pasta `arquivos_gerados/`.

---

## Arquitetura e boas práticas

Para manter o código limpo, testável e flexível, o projeto é dividido em pacotes com responsabilidades únicas:

- **`package criptografia`**  
  - Contém apenas a lógica pura de criptografia (opera em `[]byte`).
  - Não possui dependência de I/O ou caminhos de arquivo.  
  - Implementa as interfaces `IEstrategiaChave`, `IEstrategiaCertificado`, `IEstrategiaResumo`, `IEstrategiaAssinatura`.  
  - Exemplos: `GerarCertificadoAutoassinado`, `ChavePrivadaParaPEM`, `NovaEstrategiaAssinaturaPkcs1v15`.

- **`package servicos`**  
  - Responsável por todo o I/O (leitura/escrita de `.pem`, `.pdf`, `.txt`).
  - Orquestra o fluxo chamando o pacote `criptografia` e persistindo os resultados.  
  - Exemplos: `GerarParDeChaves`, `GerarCertificados`, `AssinarDocumentoPDF`, `VerificarAssinaturaDocumentoPDF`.

- **`package main`**  
  - Faz a "ligação" do sistema, injeta as dependências e documenta o uso do Strategy Pattern através da struct `ConfiguracaoCriptografia`.

Há comentários de bloco no início de cada arquivo e em todas as funções para contextualizar o papel de cada peça na demonstração — um diferencial para o portfólio.

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

## Execução

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
- A pasta `arquivos_gerados/` conterá os arquivos gerados (`.pem`, `.pdf`, `.txt` e a assinatura).

---

## Testes

O projeto inclui testes para todos os pacotes, validando geração de chaves/certificados, fluxo completo de PDF e configuração padrão.

```bash
go test ./...
```

Os testes usam chaves RSA de 1024 bits para acelerar a execução, enquanto a aplicação principal gera chaves de 2048 bits. Ajuste o tamanho conforme necessário para seus experimentos.