package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Função para conectar ao banco de dados
func ConectaDB() *sql.DB {
	conexao := "user=postgres dbname=TheBoyzz password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Estrutura de informações pessoais
type InfosPessoas struct {
	Nome      string  `json:"nome_real"`
	Sexo      string  `json:"sexo"`
	Peso      float64 `json:"peso"`
	Altura    float64 `json:"altura"`
	DataNasc  string  `json:"data_nasc"`
	LocalNasc string  `json:"local_nasc"`
}

// Estrutura dos Heróis
type Herois struct {
	InfosPessoas
	NomeHeroi    string `json:"nome_heroi"`
	Poder        string `json:"poder"`
	Popularidade int    `json:"popularidade"`
	Status       string `json:"status"`
	Forca        int    `json:"forca"`
}

// Método para exibir as informações dos heróis
func (h Herois) ExibeInfos() {
	db := ConectaDB()
	defer db.Close() // Garantir que o banco de dados seja fechado após o uso

	// Executa a consulta com JOIN
	query := `
		SELECT 
			h.nome, h.sexo, h.peso, h.altura, h.data_nasc, h.local_nasc, 
			h.nome_heroi, h.popularidade, h.status, h.forca, p.poder
		FROM 
			Herois h
		LEFT JOIN 
			Poderes p ON h.id_heroi = p.id_heroi;
	`
	allInfos, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer allInfos.Close() // Garantir que o resultado seja fechado após o uso

	// Cria uma slice para armazenar os heróis
	var informacoes []Herois

	// Itera sobre os resultados da consulta
	for allInfos.Next() {
		var heroi Herois
		err := allInfos.Scan(
			&heroi.Nome,
			&heroi.Sexo,
			&heroi.Peso,
			&heroi.Altura,
			&heroi.DataNasc,
			&heroi.LocalNasc,
			&heroi.NomeHeroi,
			&heroi.Popularidade,
			&heroi.Status,
			&heroi.Forca,
			&heroi.Poder,
		)
		if err != nil {
			log.Fatal(err)
		}
		informacoes = append(informacoes, heroi)
	}

	// Exibe as informações dos heróis
	for _, heroi := range informacoes {
		fmt.Printf("Nome: %s, Nome do Herói: %s, Poder: %s, Popularidade: %d, Força: %d\n",
			heroi.Nome, heroi.NomeHeroi, heroi.Poder, heroi.Popularidade, heroi.Forca)
	}
}
