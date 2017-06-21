package main

import (
	"fmt"
	"io/ioutil"
	"bufio"
	"strconv"
	"bytes"
	"os"
)

const arquivoIn = "in.txt"
const arquivoOut = "out.txt"

func main() {
	fmt.Printf("Quebrando arquivo: %v\n", arquivoIn)
	arquivo := NewArquivo()
	arquivo.grave()
	fmt.Printf("Arquivo %v criado com sucesso!\n", arquivoOut)
}

type Arquivo struct {
	data []byte
}

func NewArquivo() *Arquivo {
	arquivo := Arquivo{}
	arquivo.leia()
	return &arquivo
}

func (arq *Arquivo) leia() {
	dat, err := ioutil.ReadFile(arquivoIn)
	check(err)
	arq.data = dat
}

func (arq *Arquivo) maxCaracteresPorLinha() (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(arq.data))
	scanner.Scan()
	linha := scanner.Text()
	return strconv.Atoi(linha)
}

func (arq *Arquivo) texto() string {
	maxPorLinha, err := arq.maxCaracteresPorLinha()
	check(err)

	proxLinhaCount := 0
	var texto, linha string
	enter := "\n\n"

	scanner := bufio.NewScanner(bytes.NewReader(arq.data))
	scanner.Split(bufio.ScanWords) // divide em palavras
	scanner.Scan() // pula primeira linha

	// loop enquanto tiver outra palavra
	for scanner.Scan() {
		palavra := scanner.Text()
		palavra = fmt.Sprintf("%s ", palavra)

		proxLinhaCount = len(fmt.Sprintf("%s%s", linha, palavra))

		if proxLinhaCount <= maxPorLinha {
			linha = fmt.Sprintf("%s%s", linha, palavra)
		} else {
			texto = fmt.Sprintf("%s%s%s", texto, linha, enter)
			linha = palavra
		}
	}
	// não tem mais palavras, acrescente a última linha
	texto = fmt.Sprintf("%s%s%s", texto, linha, enter)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	return texto
}

func (arq *Arquivo) grave() {
	saida := []byte(arq.texto())
	err := ioutil.WriteFile(arquivoOut, saida, 0644)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
