package services

import (
	m "DWM/model"
	v "DWM/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

//this file is for parallel Download if user have given input
func ParaDownload(url string, channel chan m.ErrorType, ID uuid.UUID) {
	/* my aim is here to get the extension of the url
		so that i could download file in same extension
	splitting the url into several parts by "/"
	*/
	//  urlInfo := strings.Split(url, "/")
	//extracting the url and getting last part of the url after "."
	//  fileName := strings.Split(urlInfo[len(urlInfo)-1], ".")
	//we get the extension
	//  extension := fileName[1]

	extension := v.Extensions(url)

	switch extension[1] {
	case "jpg", "png", "gif":
		response, err := http.Get(url)
		if err != nil {
			channel <- m.ErrorType{
				ResponseMessage:    "Cannot fetch Url",
				ResponseStatusCode: 1000,
			}
		}
		defer response.Body.Close()

		file, err := os.Create(filepath.Join(ID.String(), filepath.Base(extension[0]+"."+extension[1])))
		if err != nil {
			channel <- m.ErrorType{
				ResponseMessage:    "Unable to create file",
				ResponseStatusCode: 403,
			}
		}
		defer file.Close()
		_, err = io.Copy(file, response.Body)
		if err != nil {
			channel <- m.ErrorType{
				ResponseMessage:    "Unable to save file",
				ResponseStatusCode: 403,
			}
		}
		var directory m.Directory
		directory.FileName = extension[0] + "." + extension[1]
		directory.FilePath = "/Users/vikash/Desktop/DownloadManager/app" + "/" + directory.FileName
		Direc = append(Direc, directory)
		channel <- m.ErrorType{
			ResponseMessage:    "File has been successfully downloaded",
			ResponseStatusCode: 200,
		}
	default:
		channel <- m.ErrorType{
			ResponseMessage:    "you tried to download wrong extension type file try .jpg  .png  or .gif",
			ResponseStatusCode: 403,
		}
	}

}
