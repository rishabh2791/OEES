package persistance

import (
	"errors"
	"oees/domain/entity"
	"oees/domain/repository"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type jobRepo struct {
	db          *gorm.DB
	warehouseDB *gorm.DB
	logger      hclog.Logger
}

var _ repository.JobRepository = &jobRepo{}

func newJobRepo(db *gorm.DB, logger hclog.Logger) *jobRepo {
	return &jobRepo{
		db:     db,
		logger: logger,
	}
}

type RemoteMaterial struct {
	StockCode   string
	Description string
}

func (jobRepo *jobRepo) getSKUFromRemote(stockCode string) (*RemoteMaterial, error) {
	remoteMaterial := RemoteMaterial{}
	stockCodeQuery := "SELECT StockCode, Description FROM dbo.InvMaster WHERE StockCode = '" + stockCode + "'"
	rows, getErr := jobRepo.warehouseDB.Raw(stockCodeQuery).Rows()
	defer rows.Close()
	if getErr != nil {
		return nil, getErr
	}
	for rows.Next() {
		scanErr := rows.Scan(&remoteMaterial.StockCode, &remoteMaterial.Description)
		if scanErr != nil {
			return nil, scanErr
		}
	}
	return &remoteMaterial, nil
}

func (jobRepo *jobRepo) createSKU(remoteMaterial *RemoteMaterial, job *entity.Job, ty string) (*entity.SKU, error) {
	stock := entity.SKU{}
	stock.PlantCode = job.PlantCode
	stock.Code = remoteMaterial.StockCode
	stock.Description = remoteMaterial.Description
	stock.CreatedByUsername = job.CreatedByUsername
	stock.UpdatedByUsername = job.UpdatedByUsername

	// Create Material
	stockCreationErr := jobRepo.db.Create(&stock).Error
	if stockCreationErr != nil {
		return nil, stockCreationErr
	}
	return &stock, nil
}

func (jobRepo *jobRepo) Create(job *entity.Job) (*entity.Job, error) {
	var stockCode string
	var quantity float32
	quantity = float32(job.Plan)
	stockCode = job.SKU.Code

	if stockCode != "" || len(stockCode) != 0 {
		existingSKU := entity.SKU{}
		getSKUError := jobRepo.db.Where("code = ? AND plant_id=?", stockCode, job.PlantCode).Take(&existingSKU).Error
		if getSKUError != nil {
			//Not Created
			remoteMaterial, remoteErr := jobRepo.getSKUFromRemote(stockCode)
			if remoteErr != nil {
				return nil, remoteErr
			}

			//Create Material
			material, getErr := jobRepo.createSKU(remoteMaterial, job, "Bulk")
			if getErr != nil {
				return nil, getErr
			}
			existingSKU = *material
		}
		job.SKUID = existingSKU.ID
		job.Plan = int16(quantity)

		creationErr := jobRepo.db.Create(&job).Error
		return job, creationErr

	}

	return nil, errors.New("Stock Code Not Found.")
}

func (jobRepo *jobRepo) Get(id string) (*entity.Job, error) {
	job := entity.Job{}
	getErr := jobRepo.db.
		Preload("SKU.Plant").
		Preload("SKU.Plant.CreatedBy").
		Preload("SKU.Plant.CreatedBy.UserRole").
		Preload("SKU.Plant.UpdatedBy").
		Preload("SKU.Plant.UpdatedBy.UserRole").
		Preload("SKU.CreatedBy").
		Preload("SKU.UpdatedBy").
		Preload("SKU.CreatedBy.UserRole").
		Preload("SKU.UpdatedBy.UserRole").
		Preload("Plant.CreatedBy").
		Preload("Plant.CreatedBy.UserRole").
		Preload("Plant.UpdatedBy").
		Preload("Plant.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).Take(&job).Error
	return &job, getErr
}

func (jobRepo *jobRepo) List(conditions string) ([]entity.Job, error) {
	jobs := []entity.Job{}
	getErr := jobRepo.db.
		Preload("SKU.Plant").
		Preload("SKU.Plant.CreatedBy").
		Preload("SKU.Plant.CreatedBy.UserRole").
		Preload("SKU.Plant.UpdatedBy").
		Preload("SKU.Plant.UpdatedBy.UserRole").
		Preload("SKU.CreatedBy").
		Preload("SKU.UpdatedBy").
		Preload("SKU.CreatedBy.UserRole").
		Preload("SKU.UpdatedBy.UserRole").
		Preload("Plant.CreatedBy").
		Preload("Plant.CreatedBy.UserRole").
		Preload("Plant.UpdatedBy").
		Preload("Plant.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(jobs).Error
	return jobs, getErr
}
