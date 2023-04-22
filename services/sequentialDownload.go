package services

import (
	m "DWM/model"
	"DWM/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

var Direc []m.Directory

func SeqDownload(url string, ID uuid.UUID) *m.ErrorType {
	// urlInfo := strings.Split(url, "/")
	// fileName := strings.Split(urlInfo[len(urlInfo)-1], ".")
	// extension := fileName[1]

	extension := utils.Extensions(url)

	switch extension[1] {
	case "jpg", "png", "gif":
		response, err := http.Get(url)
		if err != nil {
			return &m.ErrorType{
				ResponseMessage:    "Cannot fetch url",
				ResponseStatusCode: 403,
			}
		}
		defer response.Body.Close()

		file, err := os.Create(filepath.Join(ID.String(), filepath.Base(extension[0]+"."+extension[1])))
		if err != nil {
			return &m.ErrorType{
				ResponseMessage:    "Unable to create file",
				ResponseStatusCode: 1003,
			}
		}
		defer file.Close()
		_, err = io.Copy(file, response.Body)
		if err != nil {
			return &m.ErrorType{
				ResponseMessage:    "Unable to save file",
				ResponseStatusCode: 1003,
			}
		}
		var directory m.Directory
		directory.FileName = extension[0] + "." + extension[1]
		directory.FilePath = "/Users/vikash/Desktop/DownloadManager/app" + "/" + directory.FileName
		Direc = append(Direc, directory)
		return &m.ErrorType{
			ResponseMessage:    "File has been successfully downloaded",
			ResponseStatusCode: 200,
		}
	default:
		return &m.ErrorType{
			ResponseMessage:    "you tried to download wrong extension type file try .jpg  .png  or .gif",
			ResponseStatusCode: 404,
		}
	}

}
