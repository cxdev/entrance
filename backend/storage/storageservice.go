package storage

import (
	"entrance/backend"
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

type StorageService struct {
	db *gorm.DB
}

func NewStorageService(db *gorm.DB) (*StorageService, error) {
	result := db.AutoMigrate(&CommandEntity{})
	if result.Error != nil {
		return nil, result.Error
	}

	result = db.AutoMigrate(&JobEntity{})
	if result.Error != nil {
		return nil, result.Error
	}

	return &StorageService{db}, nil
}

func (service *StorageService) CreateCommand(name string, commandtype entrance.CommandType, commandSegments []entrance.CommandSegment) (*entrance.Command, error) {
	command := entrance.Command{entrance.Base{}, name, commandtype, commandSegments}
	commandEntity, err := NewCommandEntity(&command)
	if err != nil {
		return nil, err
	}
	result := service.db.Create(commandEntity)
	if result.Error != nil {
		return nil, result.Error
	}
	return commandEntity.ToCommand()
}

func (service *StorageService) Command(cID uint) (*entrance.Command, error) {
	var commandEntity CommandEntity
	result := service.db.Find(&commandEntity, cID)
	if result.Error != nil {
		return nil, result.Error
	}
	return commandEntity.ToCommand()
}
func (service *StorageService) Commands(qc *entrance.QueryCondition) (*[]entrance.Command, error) {
	var commandEntities []CommandEntity

	conditions, values := separateConditionsWithValues(qc)
	result := service.db.Where(conditions, values...).Find(&commandEntities)
	if result.Error != nil {
		return nil, result.Error
	}

	return ToCommands(&commandEntities)
}

func (service *StorageService) SaveCommand(command *entrance.Command) error {
	commandEntity, err := NewCommandEntity(command)
	if err != nil {
		return err
	}
	result := service.db.Save(commandEntity)
	return result.Error
}

func (service *StorageService) CreateJob(cID uint, arguments *entrance.Arguments) (*entrance.Job, error) {
	command, err := service.Command(cID)
	if err != nil {
		return nil, err
	}

	var commandSegments = command.CommandSegments
	sysCmd, err := createSysCmd(&commandSegments, arguments)
	if err != nil {
		return nil, err
	}

	job := entrance.Job{entrance.Base{}, entrance.WAITING, cID, arguments, sysCmd}
	jobEntity, err := NewJobEntity(&job)
	if err != nil {
		return nil, err
	}

	result := service.db.Create(&jobEntity)
	if result.Error != nil {
		return nil, result.Error
	}
	return jobEntity.ToJob()
}

func (service *StorageService) Job(jID uint) (*entrance.Job, error) {
	var jobEntity JobEntity
	result := service.db.Find(&jobEntity, jID)
	if result.Error != nil {
		return nil, result.Error
	}
	return jobEntity.ToJob()
}

func (service *StorageService) Jobs(qc *entrance.QueryCondition) (*[]entrance.Job, error) {
	var jobEntities []JobEntity

	conditions, values := separateConditionsWithValues(qc)
	result := service.db.Where(conditions, values...).Find(&jobEntities)
	if result.Error != nil {
		return nil, result.Error
	}

	return ToJobs(&jobEntities)

}

func (service *StorageService) SaveJob(job *entrance.Job) error {
	jobEntity, err := NewJobEntity(job)
	if err != nil {
		return err
	}
	result := service.db.Save(&jobEntity)
	return result.Error
}

func createSysCmd(cs *[]entrance.CommandSegment, arguments *entrance.Arguments) (string, error) {
	var sb strings.Builder
	for _, segment := range *cs {
		if segment.IsRequired {
			sb.WriteString(segment.Base)
			sb.WriteString(" ")
			if segment.IsValuable {
				if argVal, ok := (*arguments)[segment.Key]; ok {
					sb.WriteString(argVal)
					sb.WriteString(" ")
				} else {
					return "", errors.New("Not found error")
				}
			}
		} else {
			if argVal, ok := (*arguments)[segment.Key]; ok {
				sb.WriteString(segment.Base)
				sb.WriteString(" ")
				if argVal != "" {
					sb.WriteString(argVal)
					sb.WriteString(" ")
				}
			}
		}
	}
	return strings.TrimSpace(sb.String()), nil
}

func separateConditionsWithValues(qc *entrance.QueryCondition) (string, []interface{}) {
	var conditionItems []string
	var values []interface{}
	if qc != nil {
		for key, value := range *qc {
			conditionItems = append(conditionItems, fmt.Sprintf("%s = ?", key))
			values = append(values, value)
		}
	}
	return strings.Join(conditionItems, " AND "), values
}
