package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	//assim como no outro programa, vai ler a response e criar um map
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	//criando o map, e colocar as palavras
	words := make(map[string]string)
	//antes usava ReadAll e colocava na slice of bytes
	//agora, usar NewScanner e bufio, input output buffer, tipo impressora, n deixa o pc preso pagina a pagina
	//produz toda area que sera armazenada o que deve ser imprimido e cria esse bucket, buffer, area pra ficar o arquivo
	//e libera o pc pra fazer outras coisas
	sc := bufio.NewScanner(res.Body) //colocando todas as coisas do res.Body
	//dividindo por palavras
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = "" //pegar o token pra cada coisa que scaniou e dividiu em palavras e coloca no map
	}
	//checkar se scanner tem algum erro
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	i := 0
	for k := range words { // range de todas as palavras ate 200 e imprimir a key, o valor n é nada
		//pq o objetivo é colocar todas as palavras no map, entao associa um valor nada: ""
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++
	}

}
