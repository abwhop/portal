package services

import (
	"encoding/json"
	"github.com/abwhop/portal_models/models"
)

func ConvertFile(fileAPI *models.FileAPI) (*models.FileDB, error) {
	if fileAPI == nil {
		return nil, nil
	}
	return &models.FileDB{
		Id:           fileAPI.Id,
		Link:         fileAPI.Link,
		FileName:     fileAPI.FileName,
		OriginalName: fileAPI.OriginalName,
		ContentType:  fileAPI.ContentType,
		Size:         fileAPI.Size,
		Height:       fileAPI.Height,
		Width:        fileAPI.Width,
	}, nil
}

func ConvertFiles(filesAPI []*models.FileAPI) (*models.ListOfFileDB, error) {
	var fileDB []*models.FileDB
	var list *models.ListOfFileDB
	for _, fileAPI := range filesAPI {
		userDB, err := ConvertFile(fileAPI)
		if err != nil {
			continue
		}
		fileDB = append(fileDB, userDB)
	}
	marshal, err := json.Marshal(fileDB)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(marshal, &list); err != nil {
		return nil, err
	}
	return list, nil
}
