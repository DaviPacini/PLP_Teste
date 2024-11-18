package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Crimes struct {
	Herois
	NomeCrime       string `json:"nome_crime"`
	Severidade      string `json:"severidade"`
	DataCrime       string `json:"data_crime"`
	DescricaoEvento string `json:"descricao_evento"`
}

// Método para consultar crimes por herói e por severidade
func ConsultaCrimesPorHeroiESeveridade(nomeHeroi string, severidadeMinima int, severidadeMaxima int) ([]Crimes, error) {
	db := ConectaDB()
	defer db.Close() // Garantir que o banco de dados seja fechado após o uso

	// Consulta para buscar crimes com base no nome do herói e na severidade
	query := `
		SELECT 
			c.nome_crime, c.severidade, hc.data_crime, hc.descricao_evento
		FROM 
			Crimes c
		JOIN 
			Herois_Crimes hc ON c.id_crime = hc.id_crime
		JOIN 
			Herois h ON hc.id_heroi = h.id_heroi
		WHERE 
			h.nome_heroi = $1
		AND 
			c.severidade BETWEEN $2 AND $3;
	`

	// Executa a consulta
	rows, err := db.Query(query, nomeHeroi, severidadeMinima, severidadeMaxima)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close() // Garantir que o resultado seja fechado após o uso

	// Cria uma slice para armazenar os crimes
	var crimes []Crimes

	// Itera sobre os resultados da consulta
	for rows.Next() {
		var crime Crimes
		err := rows.Scan(
			&crime.NomeCrime,
			&crime.Severidade,
			&crime.DataCrime,
			&crime.DescricaoEvento,
		)
		if err != nil {
			log.Fatal(err)
		}
		crimes = append(crimes, crime)
	}

	// Verifica se não encontrou nenhum crime
	if len(crimes) == 0 {
		return nil, fmt.Errorf("nenhum crime encontrado para o herói %s com severidade entre %d e %d", nomeHeroi, severidadeMinima, severidadeMaxima)
	}

	return crimes, nil
}

func ConsultaCrimesPorHeroi(nomeHeroi string) ([]Crimes, error) {
	db := ConectaDB()
	defer db.Close() // Garantir que o banco de dados seja fechado após o uso

	// Consulta para buscar crimes com base no nome do herói
	query := `
		SELECT 
			c.nome_crime, c.severidade, hc.data_crime, hc.descricao_evento
		FROM 
			Crimes c
		JOIN 
			Herois_Crimes hc ON c.id_crime = hc.id_crime
		JOIN 
			Herois h ON hc.id_heroi = h.id_heroi
		WHERE 
			h.nome_heroi = $1;
	`

	// Executa a consulta
	rows, err := db.Query(query, nomeHeroi)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close() // Garantir que o resultado seja fechado após o uso

	// Cria uma slice para armazenar os crimes
	var crimes []Crimes

	// Itera sobre os resultados da consulta
	for rows.Next() {
		var crime Crimes
		err := rows.Scan(
			&crime.NomeCrime,
			&crime.Severidade,
			&crime.DataCrime,
			&crime.DescricaoEvento,
		)
		if err != nil {
			log.Fatal(err)
		}
		crimes = append(crimes, crime)
	}

	// Verifica se não encontrou nenhum crime
	if len(crimes) == 0 {
		return nil, fmt.Errorf("nenhum crime encontrado para o herói %s", nomeHeroi)
	}

	return crimes, nil
}

func ConsultaCrimesPorSeveridade(severidadeMinima int, severidadeMaxima int) ([]Crimes, error) {
	db := ConectaDB()
	defer db.Close() // Garantir que o banco de dados seja fechado após o uso

	// Consulta para buscar crimes com base na severidade
	query := `
		SELECT 
			c.nome_crime, c.severidade, hc.data_crime, hc.descricao_evento
		FROM 
			Crimes c
		JOIN 
			Herois_Crimes hc ON c.id_crime = hc.id_crime
		WHERE 
			c.severidade BETWEEN $1 AND $2;
	`

	// Executa a consulta
	rows, err := db.Query(query, severidadeMinima, severidadeMaxima)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close() // Garantir que o resultado seja fechado após o uso

	// Cria uma slice para armazenar os crimes
	var crimes []Crimes

	// Itera sobre os resultados da consulta
	for rows.Next() {
		var crime Crimes
		err := rows.Scan(
			&crime.NomeCrime,
			&crime.Severidade,
			&crime.DataCrime,
			&crime.DescricaoEvento,
		)
		if err != nil {
			log.Fatal(err)
		}
		crimes = append(crimes, crime)
	}

	// Verifica se não encontrou nenhum crime
	if len(crimes) == 0 {
		return nil, fmt.Errorf("nenhum crime encontrado com severidade entre %d e %d", severidadeMinima, severidadeMaxima)
	}

	return crimes, nil
}
