package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	//get retorna uma *response e um erro, pertence ao pacote net/http
	//
	//cuidou do erro
	if err != nil {
		log.Fatalln(err)
	}
	//Response é uma struct, um dos campos é Body, e é do tipo ioReadCloser
	// string é uma slice of bytes, varias runes, converte para string
	bs, _ := ioutil.ReadAll(res.Body) //ReadAll , vai ler tudo que pegou do GET
	str := string(bs)                 //converte a bs (slice of bytes) em string
	//pra depois printar
	fmt.Println(str) // printa

}
