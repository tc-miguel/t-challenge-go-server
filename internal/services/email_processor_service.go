package email_processor

import (
	"bufio"
	zincModels "example/hello/internal/zinc_search_api/models"
	zincService "example/hello/internal/zinc_search_api/services"
	"fmt"
	"log"
	"os"
	"strings"
)

const MainDirectoryPath string = "/home/miguel/Documents/Tchallenge/data"

/*func processDirectory(directoryPath string) ([]models.EmailJsonResult, error) {
	log.Printf("--------------------------Start directory %s process --------------------------\n", directoryPath)
	emailFileList, error := os.ReadDir(directoryPath)
	//Array to save Json result for send Text processing API
	var emailJsonResultList = make([]models.EmailJsonResult, 1)
	if error != nil {
		return emailJsonResultList, error
	}
	for _, file := range emailFileList {
		emailReadRecord := processEmailFile(directoryPath, file)
		emailJsonResultList = append(emailJsonResultList, emailReadRecord)
	}
	log.Printf("--------------------------End directory %s process --------------------------\n", directoryPath)
	return emailJsonResultList, nil
}*/

func ProcessEmailFiles() (bool, error) {
	log.Println("Start proccess email file reading")
	return processDirectoryOrFile(MainDirectoryPath)
}

func processDirectoryOrFile(directoryPath string) (bool, error) {
	//dirList: List of directories for processing in MainDirectoryPath
	//error: Object instance with error if present
	dirList, error := os.ReadDir(directoryPath)
	if error != nil {
		log.Printf("Error trying to read parent directory: %v\n", error.Error())
		return false, error
	}

	emailJsonResultList := make([]zincModels.ZincEmailDocument, 0)

	for _, directoryOrFile := range dirList {
		if directoryOrFile.IsDir() {
			log.Printf("Directorio a procesar es: %s", directoryPath+"/"+directoryOrFile.Name())
			dirPath := directoryPath + "/" + directoryOrFile.Name()
			processDirectoryOrFile(dirPath)
		} else {
			emailJsonItem := processEmailFile(directoryPath, directoryOrFile.Name()) /*    processDirectory(directoryPath)*/
			emailJsonResultList = append(emailJsonResultList, emailJsonItem)
			emailJsonResultList = sendZincSearchItems(emailJsonResultList)
		}
	}
	sendMissingZincSearchItems(emailJsonResultList)
	log.Println("End proccess email file reading")
	return true, nil
}

func sendMissingZincSearchItems(emailJsonResultList []zincModels.ZincEmailDocument) {
	if len(emailJsonResultList) > 0 {
		zincService.AddEmailDocuments(emailJsonResultList)
	}
}

func sendZincSearchItems(emailJsonResultList []zincModels.ZincEmailDocument) []zincModels.ZincEmailDocument {
	if len(emailJsonResultList) >= 50 {
		zincService.AddEmailDocuments(emailJsonResultList)
		return make([]zincModels.ZincEmailDocument, 0)
	}
	return emailJsonResultList

}

func processEmailFile(directoryPath string, emailFileName string) zincModels.ZincEmailDocument {
	//Try Open file Stream
	filePath := directoryPath + "/" + emailFileName
	file, error := os.Open(filePath)
	if error != nil {
		fmt.Printf("Could not open file %s", error.Error())
		return zincModels.ZincEmailDocument{
			Error: error,
		}
	}
	defer file.Close()

	return readEmailFile(file)
}

func readEmailFile(file *os.File) zincModels.ZincEmailDocument {
	startBodyRead := false
	var bodyBuilder strings.Builder
	item := zincModels.ZincEmailDocument{}
	scanner := bufio.NewScanner(file)

	config := GetFieldConfig(&item)
	fieldProcessName := ""
	for scanner.Scan() {
		line := scanner.Text()
		if isBodyRead(startBodyRead, line) {
			startBodyRead = true
			bodyBuilder.WriteString(line)
		} else {
			for key, value := range config {
				if matches := value.Regex.FindStringSubmatch(line); len(matches) > 1 {
					fieldProcessName = key
					break
				}
			}
			changeFieldValue(config, fieldProcessName, line)
		}
	}
	item.Body = bodyBuilder.String()
	//TODO clear this code after finish
	/*jsonResult, error := json.MarshalIndent(item, "", "    ")
	if error != nil {
		log.Fatal("is not possible parse from objet to Json format")
	}
	log.Printf("Item finish process: info: %s", string(jsonResult))*/
	return item
}

func isBodyRead(bodyRead bool, line string) bool {
	if bodyRead {
		return bodyRead
	} else {
		return len(strings.TrimSpace(line)) == 0
	}
}
