package application

import (
	"oees/domain/entity"
	"oees/domain/repository"
)

type jobApp struct {
	jobReporitory repository.JobRepository
}

var _ jobAppInterface = &jobApp{}

func newJobApp(jobrepository repository.JobRepository) *jobApp {
	return &jobApp{
		jobReporitory: jobrepository,
	}
}

type jobAppInterface interface {
	Create(job *entity.Job) (*entity.Job, error)
	Get(id string) (*entity.Job, error)
	List(conditions string) ([]entity.Job, error)
	PullFromRemote(username string) error
}

func (jobApp *jobApp) Create(job *entity.Job) (*entity.Job, error) {
	return jobApp.jobReporitory.Create(job)
}

func (jobApp *jobApp) Get(id string) (*entity.Job, error) {
	return jobApp.jobReporitory.Get(id)
}

func (jobApp *jobApp) List(conditions string) ([]entity.Job, error) {
	return jobApp.jobReporitory.List(conditions)
}

func (jobApp *jobApp) PullFromRemote(username string) error {
	return jobApp.jobReporitory.PullFromRemote(username)
}
