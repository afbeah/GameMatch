package service

import (
	model "GamesInsertion/src/Model"
	"errors"
	"gorm.io/gorm"
)

type MatchService interface {
	AddGame(match *model.Match) error
	GetGame(id int) (*model.Match, error)
	UpdateGame(match *model.Match) error
	DeleteGame(id int) error
	GetAllGames() ([]*model.Match, error)
}

type matchService struct{
	db *gorm.DB
}

func NewMatchService(db *gorm.DB) MatchService {
	return &matchService{db: db}
}

func (s *matchService) AddGame(match *model.Match) error {
	if err := s.db.Create(match).Error; err != nil{
		return errors.New("failed to add game")
	}
	return nil
}


func (s *matchService) GetGame(id int) (*model.Match, error) {
	var match model.Match
	if err := s.db.First(&match, id).Error; err != nil{
		return nil, errors.New("game not found")
	}
	return &match, nil
}

func (s *matchService) UpdateGame(match *model.Match) error {
if err := s.db.Save(match).Error; err != nil{
	return errors.New("failed to update game")
}
	return nil
}

func (s *matchService) DeleteGame(id int) error {
if err := s.db.Delete(&model.Match{}, id).Error; err != nil{
	return errors.New("game not found")
}
	return nil
}

func (s *matchService) GetAllGames() ([]*model.Match, error) {
    var matches []*model.Match
    if err := s.db.Find(&matches).Error; err != nil {
        return nil, err
    }
    return matches, nil
}
