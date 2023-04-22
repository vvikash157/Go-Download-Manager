package utils

import "strings"

func Extensions(url string) []string {

	urlInfo := strings.Split(url, "/")
	//extracting the url and getting last part of the url after "."
	fileName := strings.Split(urlInfo[len(urlInfo)-1], ".")
	//we get the extension
	// extension_url := fileName[1]
	return fileName

}
