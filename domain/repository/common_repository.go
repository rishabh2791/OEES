package repository

type CommonRepository interface {
	GetTables() ([]string, error)
}
