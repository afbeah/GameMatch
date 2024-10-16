package service

import (
	model "GamesInsertion/src/Model"
	"errors"
)

// Criação do Service
type MatchService interface {
	AddGame(match *model.Match) error
	GetGame(id int) (*model.Match, error)
	UpdateGame(match *model.Match) error
	DeleteGame(id int) error
}

// Utilizando o map para armazenamento em memória (futuramente usar um BD?)
type matchService struct {
	matchs map[int]*model.Match
}

func NewMatchService() MatchService {
	return &matchService{
		matchs: make(map[int]*model.Match),
	}
}

// informa se o livro já existe no sistema
func (s *matchService) AddGame(match *model.Match) error {
	if _, exists := s.matchs[match.ID]; exists {
		return errors.New("game already exists")
	}
	s.matchs[match.ID] = match
	return nil
}

// Informa se o livro não foi encontrado no sistema
func (s *matchService) GetGame(id int) (*model.Match, error) {
	match, exists := s.matchs[id]
	if !exists {
		return nil, errors.New("game not found")
	}
	return match, nil
}

func (s *matchService) UpdateGame(match *model.Match) error {
	if _, exists := s.matchs[match.ID]; !exists {
		return errors.New("game not found")
	}
	s.matchs[match.ID] = match
	return nil
}

func (s *matchService) DeleteGame(id int) error {
	if _, exists := s.matchs[id]; !exists {
		return errors.New("game not found")
	}
	delete(s.matchs, id)
	return nil
}
