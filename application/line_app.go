package application

import (
	"oees/domain/entity"
	"oees/domain/repository"
)

type lineApp struct {
	lineRepository repository.LineRepository
}

var _ lineAppInterface = &lineApp{}

func newlineApp(lineRepository repository.LineRepository) *lineApp {
	return &lineApp{
		lineRepository: lineRepository,
	}
}

func (lineApp *lineApp) Create(line *entity.Line) (*entity.Line, error) {
	return lineApp.lineRepository.Create(line)
}

func (lineApp *lineApp) Get(linename string) (*entity.Line, error) {
	return lineApp.lineRepository.Get(linename)
}

func (lineApp *lineApp) List(conditions string) ([]entity.Line, error) {
	return lineApp.lineRepository.List(conditions)
}

func (lineApp *lineApp) Update(linename string, update *entity.Line) (*entity.Line, error) {
	return lineApp.lineRepository.Update(linename, update)
}

type lineAppInterface interface {
	Create(line *entity.Line) (*entity.Line, error)
	Get(linename string) (*entity.Line, error)
	List(conditions string) ([]entity.Line, error)
	Update(linename string, update *entity.Line) (*entity.Line, error)
}
