package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

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
	NomeHeroi    string   `json:"nome_heroi"`
	Poderes      []string `json:"poder"`
	Popularidade int      `json:"popularidade"`
	Status       string   `json:"status"`
	Forca        int      `json:"forca"`
}

// Método para exibir as informações dos heróis
func (h Herois) ExibeInfosGerais() []Herois {
	db := ConectaDB()
	defer db.Close() // Garantir que o banco de dados seja fechado após o uso

	query := `
		SELECT 
			h.nome, h.sexo, h.peso, h.altura, h.data_nasc, h.local_nasc, 
			h.nome_heroi, h.popularidade, h.status, h.forca, 
			STRING_AGG(p.poder, ', ') AS poderes
		FROM 
			Herois h
		LEFT JOIN 
			Poderes p ON h.id_heroi = p.id_heroi
		GROUP BY 
			h.id_heroi, h.nome, h.sexo, h.peso, h.altura, h.data_nasc, h.local_nasc, 
			h.nome_heroi, h.popularidade, h.status, h.forca;
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
		var poderes string
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
			&poderes,
		)
		if err != nil {
			log.Fatal(err)
		}
		// Divide a string de poderes em uma slice
		heroi.Poderes = splitPoderes(poderes)
		informacoes = append(informacoes, heroi)
	}

	// Exibe as informações dos heróis
	// for _, heroi := range informacoes {
	// 	fmt.Printf("Nome: %s, Nome do Herói: %s, Poderes: %v, Popularidade: %d, Força: %d\n",
	// 		heroi.Nome, heroi.NomeHeroi, heroi.Poderes, heroi.Popularidade, heroi.Forca)
	// }
	return informacoes
}

// Função para dividir poderes em uma slice
func splitPoderes(poderes string) []string {
	if poderes == "" {
		return []string{}
	}
	return strings.Split(poderes, ", ")
}

func BuscaHeroiPorNome(nomeHeroi string) (*Herois, error) {
	db := ConectaDB()
	defer db.Close() // Garantir que o banco de dados seja fechado após o uso

	// Consulta para buscar um herói específico pelo nome do herói
	query := `
		SELECT 
			h.nome, h.sexo, h.peso, h.altura, h.data_nasc, h.local_nasc, 
			h.nome_heroi, h.popularidade, h.status, h.forca, 
			STRING_AGG(p.poder, ', ') AS poderes
		FROM 
			Herois h
		LEFT JOIN 
			Poderes p ON h.id_heroi = p.id_heroi
		WHERE 
			h.nome_heroi = $1
		GROUP BY 
			h.id_heroi, h.nome, h.sexo, h.peso, h.altura, h.data_nasc, h.local_nasc, 
			h.nome_heroi, h.popularidade, h.status, h.forca;
	`

	// Executa a consulta
	var heroi Herois
	var poderes string
	err := db.QueryRow(query, nomeHeroi).Scan(
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
		&poderes,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("herói com nome %s não encontrado", nomeHeroi)
		}
		return nil, err
	}

	// Divide a string de poderes em uma slice
	heroi.Poderes = splitPoderes(poderes)

	return &heroi, nil
}

func BuscaHeroisPorPopularidade(minPopularidade, maxPopularidade int) ([]Herois, error) {
	db := ConectaDB()
	defer db.Close()

	query := `
		SELECT nome, sexo, peso, altura, data_nasc, local_nasc, 
		       nome_heroi, popularidade, status, forca
		FROM Herois
		WHERE popularidade BETWEEN $1 AND $2
		ORDER BY popularidade DESC;
	`

	rows, err := db.Query(query, minPopularidade)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var herois []Herois
	for rows.Next() {
		var heroi Herois
		err := rows.Scan(
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
		)
		if err != nil {
			return nil, err
		}
		herois = append(herois, heroi)
	}

	return herois, nil
}

func BuscaHeroisPorStatus(status string) ([]Herois, error) {
	db := ConectaDB()
	defer db.Close()

	// Consulta SQL para buscar heróis pelo status
	query := `
		SELECT nome, sexo, peso, altura, data_nasc, local_nasc, 
		       nome_heroi, popularidade, status_atividade, forca
		FROM Herois
		WHERE status_atividade = $1;
	`

	rows, err := db.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var herois []Herois
	for rows.Next() {
		var heroi Herois
		err := rows.Scan(
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
		)
		if err != nil {
			return nil, err
		}
		herois = append(herois, heroi)
	}

	return herois, nil
}

func CadastrarHeroiComPoderesNormalizados(heroi Herois, poderes []struct {
	Poder     string
	Descricao string
}) error {
	db := ConectaDB()
	defer db.Close()

	// Inicia uma transação para garantir consistência entre as tabelas
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("erro ao iniciar transação: %w", err)
	}

	// Consulta para inserir o herói
	queryHeroi := `
		INSERT INTO Herois (
			nome_heroi, nome_real, sexo, altura, peso, data_nasc, local_nasc, 
			popularidade, forca, status_atividade
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id_heroi;
	`

	// Executa a consulta e captura o id do herói recém-inserido
	var idHeroi int
	err = tx.QueryRow(queryHeroi,
		heroi.NomeHeroi,
		heroi.Nome, // Nome real
		heroi.Sexo,
		heroi.Altura,
		heroi.Peso,
		heroi.DataNasc,
		heroi.LocalNasc,
		heroi.Popularidade,
		heroi.Forca,
		heroi.Status,
	).Scan(&idHeroi)

	if err != nil {
		tx.Rollback() // Reverte a transação em caso de erro
		return fmt.Errorf("erro ao cadastrar o herói: %w", err)
	}

	// Consulta para inserir os poderes
	queryPoder := `
		INSERT INTO Poderes (
			id_heroi, poder, descricao
		) VALUES ($1, $2, $3);
	`

	// Itera sobre os poderes e os insere na tabela
	for _, poder := range poderes {
		_, err := tx.Exec(queryPoder, idHeroi, poder.Poder, poder.Descricao)
		if err != nil {
			tx.Rollback() // Reverte a transação em caso de erro
			return fmt.Errorf("erro ao cadastrar poder '%s': %w", poder.Poder, err)
		}
	}

	// Confirma a transação
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("erro ao confirmar transação: %w", err)
	}

	fmt.Println("Herói e poderes cadastrados com sucesso!")
	return nil
}

func Remove(id int) error {
	db := ConectaDB()
	defer db.Close()

	query := `DELETE FROM Herois WHERE id_heroi = $1`

	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erro ao remover herói com id %d: %w", id, err)
	}
	return nil
}
