package persistance

import (
	"errors"
	"log"
	"oees/domain/entity"
	"oees/domain/repository"
	"reflect"
	"strings"

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

func newJobRepo(db *gorm.DB, warehouseDB *gorm.DB, logger hclog.Logger) *jobRepo {
	return &jobRepo{
		db:          db,
		warehouseDB: warehouseDB,
		logger:      logger,
	}
}

type RemoteJob struct {
	Job       string
	StockCode string
	QtyToMake float32
}

type RemoteMaterial struct {
	StockCode    string
	Description  string
	CaseLot      float32
	LowRunSpeed  int
	HighRunSpeed int
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
	caseLot, caseLotErr := jobRepo.GetCaseLot(stockCode)
	if caseLotErr != nil {
		return nil, caseLotErr
	}
	remoteMaterial.CaseLot = caseLot
	lowRunSpeed, _ := jobRepo.GetLowRunSpeed(stockCode)
	remoteMaterial.LowRunSpeed = int(lowRunSpeed / 60)
	highRunSpeed, _ := jobRepo.GetHighRunSpeed(stockCode)
	remoteMaterial.HighRunSpeed = int(highRunSpeed / 60)
	return &remoteMaterial, nil
}

func (jobRepo *jobRepo) createSKU(remoteMaterial *RemoteMaterial, username string) (*entity.SKU, error) {
	stock := entity.SKU{}
	stock.Code = remoteMaterial.StockCode
	stock.Description = remoteMaterial.Description
	stock.CaseLot = remoteMaterial.CaseLot
	stock.LowRunSpeed = remoteMaterial.LowRunSpeed
	stock.HighRunSpeed = remoteMaterial.HighRunSpeed

	stock.CreatedByUsername = username
	stock.UpdatedByUsername = username

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
		getSKUError := jobRepo.db.Where("code = ?", stockCode).Take(&existingSKU).Error
		if getSKUError != nil {
			//Not Created
			remoteMaterial, remoteErr := jobRepo.getSKUFromRemote(stockCode)
			if remoteErr != nil {
				return nil, remoteErr
			}

			//Create Material
			material, getErr := jobRepo.createSKU(remoteMaterial, job.CreatedByUsername)
			if getErr != nil {
				return nil, getErr
			}
			existingSKU = *material
		}
		job.SKUID = existingSKU.ID
		job.SKU = &existingSKU
		job.Plan = float32(quantity)

		creationErr := jobRepo.db.Create(&job).Error
		return job, creationErr

	}

	return nil, errors.New("Stock Code Not Found.")
}

func (jobRepo *jobRepo) Get(id string) (*entity.Job, error) {
	job := entity.Job{}
	getErr := jobRepo.db.
		Preload("SKU.CreatedBy").
		Preload("SKU.UpdatedBy").
		Preload("SKU.CreatedBy.UserRole").
		Preload("SKU.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where("id = ?", id).Take(&job).Error
	return &job, getErr
}

func (jobRepo *jobRepo) List(conditions string) ([]entity.Job, error) {
	jobs := []entity.Job{}
	getErr := jobRepo.db.
		Preload("SKU.CreatedBy").
		Preload("SKU.UpdatedBy").
		Preload("SKU.CreatedBy.UserRole").
		Preload("SKU.UpdatedBy.UserRole").
		Preload("CreatedBy.UserRole").
		Preload("UpdatedBy.UserRole").
		Preload(clause.Associations).Where(conditions).Find(&jobs).Error
	return jobs, getErr
}

func (jobRepo *jobRepo) GetOpenJobs() ([]RemoteJob, error) {
	remoteJobs := []RemoteJob{}
	jobQuery := "SELECT Job, StockCode, QtyToMake FROM dbo.WipMaster WHERE Complete = 'N'"
	rows, getErr := jobRepo.warehouseDB.Raw(jobQuery).Rows()
	defer rows.Close()
	if getErr != nil {
		return nil, getErr
	}
	for rows.Next() {
		remoteJob := RemoteJob{}
		scanErr := rows.Scan(&remoteJob.Job, &remoteJob.StockCode, &remoteJob.QtyToMake)
		if scanErr != nil {
			return nil, scanErr
		}
		remoteJobs = append(remoteJobs, remoteJob)
	}
	return remoteJobs, nil
}

func (jobRepo *jobRepo) GetCaseLot(stockCode string) (float32, error) {
	var unitsPerCase float32
	var thisStockCode string
	queryString := "SELECT StockCode, UnitsPerCase FROM [dbo].[InvMaster+] WHERE StockCode = '" + stockCode + "';"
	rows, getErr := jobRepo.warehouseDB.Raw(queryString).Rows()
	if getErr != nil {
		return 0, getErr
	}
	for rows.Next() {
		scanErr := rows.Scan(&thisStockCode, &unitsPerCase)
		if scanErr != nil {
			return 0, scanErr
		}
	}
	return unitsPerCase, nil
}

func (jobRepo *jobRepo) GetHighRunSpeed(stockCode string) (int, error) {
	var runSpeed int
	var thisStockCode string
	queryString := "SELECT StockCode, RunSpeed2 FROM [dbo].[InvMaster+] WHERE StockCode = '" + stockCode + "';"
	rows, getErr := jobRepo.warehouseDB.Raw(queryString).Rows()
	if getErr != nil {
		return 12, getErr
	}
	for rows.Next() {
		scanErr := rows.Scan(&thisStockCode, &runSpeed)
		if scanErr != nil {
			return 12, scanErr
		}
	}
	return runSpeed, nil
}

func (jobRepo *jobRepo) GetLowRunSpeed(stockCode string) (int, error) {
	var runSpeed int
	var thisStockCode string
	queryString := "SELECT StockCode, RunSpeed FROM [dbo].[InvMaster+] WHERE StockCode = '" + stockCode + "';"
	rows, getErr := jobRepo.warehouseDB.Raw(queryString).Rows()
	if getErr != nil {
		return 10, getErr
	}
	for rows.Next() {
		scanErr := rows.Scan(&thisStockCode, &runSpeed)
		if scanErr != nil {
			return 10, scanErr
		}
	}
	return runSpeed, nil
}

func (jobRepo *jobRepo) PullFromRemote(username string) error {
	remoteJobs, remoteErr := jobRepo.GetOpenJobs()
	error := ""
	if remoteErr != nil {
		return remoteErr
	}
	for _, remoteJob := range remoteJobs {
		log.Println(remoteJob)
		jobCode := remoteJob.Job[9:len(remoteJob.Job)]
		existingSKU := entity.SKU{}
		getSKUError := jobRepo.db.Where("code = ?", remoteJob.StockCode).Take(&existingSKU).Error
		if getSKUError != nil {
			//Not Created
			remoteMaterial, remoteErr := jobRepo.getSKUFromRemote(remoteJob.StockCode)
			if remoteErr != nil {
				error += remoteErr.Error()
			} else {
				material, getErr := jobRepo.createSKU(remoteMaterial, username)
				if getErr != nil {
					error += getErr.Error()
				}
				existingSKU = *material
			}
			//Create Material
		}
		if (!reflect.DeepEqual(existingSKU, entity.SKU{})) {
			job := entity.Job{}
			job.Code = jobCode
			job.SKUID = existingSKU.ID
			job.Plan = remoteJob.QtyToMake * existingSKU.CaseLot
			job.CreatedByUsername = username
			job.UpdatedByUsername = username
			jobCreationError := jobRepo.db.Create(&job).Error
			if jobCreationError != nil {
				if strings.Contains(jobCreationError.Error(), "Duplicate") {
					error += "Job " + job.Code + " already created.\n"
				} else {
					error += jobCreationError.Error() + "\n"
				}
			}
		}

	}
	if len(error) == 0 {
		return nil
	}
	log.Println(error)
	return errors.New(error)
}

// package persistance

// import (
// 	"errors"
// 	"oees/domain/entity"
// 	"oees/domain/repository"
// 	"strings"

// 	"github.com/hashicorp/go-hclog"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/clause"
// )

// type jobRepo struct {
// 	db          *gorm.DB
// 	warehouseDB *gorm.DB
// 	logger      hclog.Logger
// }

// var _ repository.JobRepository = &jobRepo{}

// func newJobRepo(db *gorm.DB, warehouseDB *gorm.DB, logger hclog.Logger) *jobRepo {
// 	return &jobRepo{
// 		db:          db,
// 		warehouseDB: warehouseDB,
// 		logger:      logger,
// 	}
// }

// type RemoteJob struct {
// 	Job       string
// 	StockCode string
// 	QtyToMake float32
// }

// type RemoteMaterial struct {
// 	StockCode   string
// 	Description string
// 	CaseLot     float32
// 	RunSpeed1   float32
// 	RunSpeed2   float32
// }

// func (jobRepo *jobRepo) getSKUFromRemote(stockCode string) (*RemoteMaterial, error) {
// 	remoteMaterial := RemoteMaterial{}
// 	stockCodeQuery := "SELECT StockCode, Description FROM dbo.InvMaster WHERE StockCode = '" + stockCode + "'"
// 	rows, getErr := jobRepo.warehouseDB.Raw(stockCodeQuery).Rows()
// 	defer rows.Close()
// 	if getErr != nil {
// 		return nil, getErr
// 	}
// 	for rows.Next() {
// 		scanErr := rows.Scan(&remoteMaterial.StockCode, &remoteMaterial.Description)
// 		if scanErr != nil {
// 			return nil, scanErr
// 		}
// 	}
// 	caseLot, runSpeed1, runSpeed2, caseLotErr := jobRepo.GetCaseLot(stockCode)
// 	if caseLotErr != nil {
// 		return nil, caseLotErr
// 	}
// 	remoteMaterial.CaseLot = caseLot
// 	remoteMaterial.RunSpeed1 = runSpeed1
// 	remoteMaterial.RunSpeed2 = runSpeed2
// 	return &remoteMaterial, nil
// }

// func (jobRepo *jobRepo) createSKU(remoteMaterial *RemoteMaterial, username string) (*entity.SKU, error) {
// 	stock := entity.SKU{}
// 	stock.Code = remoteMaterial.StockCode
// 	stock.Description = remoteMaterial.Description
// 	stock.CaseLot = remoteMaterial.CaseLot
// 	stock.CreatedByUsername = username
// 	stock.UpdatedByUsername = username

// 	// Create Material
// 	stockCreationErr := jobRepo.db.Create(&stock).Error
// 	if stockCreationErr != nil {
// 		return nil, stockCreationErr
// 	}
// 	return &stock, nil
// }

// func (jobRepo *jobRepo) Create(job *entity.Job) (*entity.Job, error) {
// 	var stockCode string
// 	var quantity float32
// 	quantity = float32(job.Plan)
// 	stockCode = job.SKU.Code

// 	if stockCode != "" || len(stockCode) != 0 {
// 		existingSKU := entity.SKU{}
// 		getSKUError := jobRepo.db.Where("code = ?", stockCode).Take(&existingSKU).Error
// 		if getSKUError != nil {
// 			//Not Created
// 			remoteMaterial, remoteErr := jobRepo.getSKUFromRemote(stockCode)
// 			if remoteErr != nil {
// 				return nil, remoteErr
// 			}

// 			//Create Material
// 			material, getErr := jobRepo.createSKU(remoteMaterial, job.CreatedByUsername)
// 			if getErr != nil {
// 				return nil, getErr
// 			}
// 			existingSKU = *material
// 		}
// 		job.SKUID = existingSKU.ID
// 		job.SKU = &existingSKU
// 		job.Plan = float32(quantity)

// 		creationErr := jobRepo.db.Create(&job).Error
// 		return job, creationErr

// 	}

// 	return nil, errors.New("Stock Code Not Found.")
// }

// func (jobRepo *jobRepo) Get(id string) (*entity.Job, error) {
// 	job := entity.Job{}
// 	getErr := jobRepo.db.
// 		Preload("SKU.CreatedBy").
// 		Preload("SKU.UpdatedBy").
// 		Preload("SKU.CreatedBy.UserRole").
// 		Preload("SKU.UpdatedBy.UserRole").
// 		Preload("CreatedBy.UserRole").
// 		Preload("UpdatedBy.UserRole").
// 		Preload(clause.Associations).Where("id = ?", id).Take(&job).Error
// 	return &job, getErr
// }

// func (jobRepo *jobRepo) List(conditions string) ([]entity.Job, error) {
// 	jobs := []entity.Job{}
// 	getErr := jobRepo.db.
// 		Preload("SKU.CreatedBy").
// 		Preload("SKU.UpdatedBy").
// 		Preload("SKU.CreatedBy.UserRole").
// 		Preload("SKU.UpdatedBy.UserRole").
// 		Preload("CreatedBy.UserRole").
// 		Preload("UpdatedBy.UserRole").
// 		Preload(clause.Associations).Where(conditions).Find(&jobs).Error
// 	return jobs, getErr
// }

// func (jobRepo *jobRepo) GetOpenJobs() ([]RemoteJob, error) {
// 	remoteJobs := []RemoteJob{}
// 	jobQuery := "SELECT Job, StockCode, QtyToMake FROM dbo.WipMaster WHERE Complete = 'N' AND (StockCode LIKE '40%' OR StockCode LIKE '80%')"
// 	rows, getErr := jobRepo.warehouseDB.Raw(jobQuery).Rows()
// 	defer rows.Close()
// 	if getErr != nil {
// 		return nil, getErr
// 	}
// 	for rows.Next() {
// 		remoteJob := RemoteJob{}
// 		scanErr := rows.Scan(&remoteJob.Job, &remoteJob.StockCode, &remoteJob.QtyToMake)
// 		if scanErr != nil {
// 			return nil, scanErr
// 		}
// 		remoteJobs = append(remoteJobs, remoteJob)
// 	}
// 	return remoteJobs, nil
// }

// func (jobRepo *jobRepo) GetCaseLot(stockCode string) (float32, float32, float32, error) {
// 	var unitsPerCase float32
// 	var runSpeed1 float32
// 	var runSpeed2 float32
// 	var thisStockCode string
// 	queryString := "SELECT StockCode, UnitsPerCase, ...... FROM [dbo].[InvMaster+] WHERE StockCode = '" + stockCode + "';"
// 	rows, getErr := jobRepo.warehouseDB.Raw(queryString).Rows()
// 	if getErr != nil {
// 		return 0, 0, 0, getErr
// 	}
// 	for rows.Next() {
// 		scanErr := rows.Scan(&thisStockCode, &unitsPerCase, &runSpeed1, &runSpeed2)
// 		if scanErr != nil {
// 			return 0, 0, 0, scanErr
// 		}
// 	}
// 	return unitsPerCase, runSpeed1, runSpeed2, nil
// }

// func (jobRepo *jobRepo) PullFromRemote(username string) error {
// 	remoteJobs, remoteErr := jobRepo.GetOpenJobs()
// 	error := ""
// 	if remoteErr != nil {
// 		return remoteErr
// 	}
// 	for _, remoteJob := range remoteJobs {
// 		jobCode := remoteJob.Job[9:len(remoteJob.Job)]
// 		existingSKU := entity.SKU{}
// 		getSKUError := jobRepo.db.Where("code = ?", remoteJob.StockCode).Take(&existingSKU).Error
// 		if getSKUError != nil {
// 			//Not Created
// 			remoteMaterial, remoteErr := jobRepo.getSKUFromRemote(remoteJob.StockCode)
// 			if remoteErr != nil {
// 				error += remoteErr.Error()
// 			}

// 			//Create Material
// 			material, getErr := jobRepo.createSKU(remoteMaterial, username)
// 			if getErr != nil {
// 				error += getErr.Error()
// 			}
// 			existingSKU = *material
// 		}
// 		job := entity.Job{}
// 		job.Code = jobCode
// 		job.SKUID = existingSKU.ID
// 		job.Plan = remoteJob.QtyToMake * existingSKU.CaseLot
// 		job.CreatedByUsername = username
// 		job.UpdatedByUsername = username
// 		jobCreationError := jobRepo.db.Create(&job).Error
// 		if jobCreationError != nil {
// 			if strings.Contains(jobCreationError.Error(), "Duplicate") {
// 				error += "Job " + job.Code + " already created.\n"
// 			} else {
// 				error += jobCreationError.Error() + "\n"
// 			}
// 		}
// 	}
// 	if len(error) == 0 {
// 		return nil
// 	}
// 	return errors.New(error)
// }
