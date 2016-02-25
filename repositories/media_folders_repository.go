package repositories

import (
	"github.com/deluan/gosonic/models"
	"github.com/astaxie/beego"
)

type MediaFolderRepository struct {}

func (*MediaFolderRepository) GetAll() ([]*models.MediaFolder, error) {
	mediaFolder := models.MediaFolder{Id: "0", Name: "iTunes Library", Path: beego.AppConfig.String("musicFolder")}
	result := make([]*models.MediaFolder, 1)
	result[0] = &mediaFolder
	return result, nil
}