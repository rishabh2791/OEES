package application

import "oees/domain/repository"

type commonApp struct {
	commonRepository repository.CommonRepository
}

var _ CommonAppInterface = &commonApp{}

func newCommonApp(commonRepository repository.CommonRepository) *commonApp {
	return &commonApp{
		commonRepository: commonRepository,
	}
}

func (commonApp *commonApp) GetTables() ([]string, error) {
	return commonApp.commonRepository.GetTables()
}

type CommonAppInterface interface {
	GetTables() ([]string, error)
}
