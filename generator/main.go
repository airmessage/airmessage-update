package main

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v3"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//The directory to read input files from
const dirInput = "input"
//The directory to write output files
const dirOutput = "output"
//The name of the base update data file
const fileUpdate = "update.yaml"
//The prefix of file names used to hold notes
const prefixNotes = "notes-"
//The permission to apply to output files
const outputPerm os.FileMode = 0755

// UpdateOutputTuple holds a writable update along with its channel ID
type UpdateOutputTuple struct {
	Channel string
	Output UpdateOutput
}

func main() {
	//Iterate through all files in the input directory
	files, err := os.ReadDir(dirInput)
	if err != nil {
		log.Fatal(err)
	}

	var updateOutputSlice []UpdateOutputTuple
	for _, file := range files {
		if !file.IsDir() {
			log.Printf("Skipping non-dir file %s\n", file.Name())
			continue
		}

		//Create the update
		channelID := file.Name() //Use the directory name as the channel ID
		update := processDir(file)
		updateOutputSlice = append(updateOutputSlice, UpdateOutputTuple{
			Channel: channelID,
			Output:  update,
		})

		log.Printf("Loaded channel %s with version %s and %d note(s)\n", channelID, update.Data["versionName"], len(update.Notes))
	}

	//Create the output directory
	err = os.Mkdir(dirOutput, outputPerm)
	if err != nil {
		if errors.Is(err, fs.ErrExist) {
			//If the directory already exists, make sure it's empty
			outputFileSlice, err := os.ReadDir(dirOutput)
			if err != nil {
				log.Fatal(err)
			}

			for _, outputFile := range outputFileSlice {
				err = os.RemoveAll(filepath.Join(dirOutput, outputFile.Name()))
				if err != nil {
					log.Fatal(err)
				}
			}
		} else {
			log.Fatal(err)
		}
	}

	//Write output files
	for _, update := range updateOutputSlice {
		//Combine the update data and notes
		updateOutput := update.Output.Data
		updateOutput["notes"] = update.Output.Notes

		updateBytes, err := json.Marshal(updateOutput)
		if err != nil {
			log.Fatal(err)
		}

		fileName := filepath.Join(dirOutput, update.Channel + ".json")
		err = os.WriteFile(fileName, updateBytes, outputPerm)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Wrote update file %s\n", fileName)
	}
}

// processDir reads files from an input directory and generates an UpdateOutput
func processDir(dir fs.DirEntry) UpdateOutput {
	updateDir := filepath.Join(dirInput, dir.Name())

	//Read the update file
	updateFile := filepath.Join(updateDir, "update.yaml")
	updateFileData, err := os.ReadFile(updateFile)
	if err != nil {
		log.Fatal(err)
	}

	//Parse the update file
	var updateData map[string]interface{}
	err = yaml.Unmarshal(updateFileData, &updateData)
	if err != nil {
		log.Fatal(err)
	}

	//Read update notes
	files, err := os.ReadDir(updateDir)
	if err != nil {
		log.Fatal(err)
	}

	var updateNotesSlice []UpdateNotes
	for _, file := range files {
		filePath := filepath.Join(updateDir, file.Name())

		//Ignore directories
		if file.IsDir() {
			log.Printf("Skipping dir in update directory %s\n", filePath)
			continue
		}

		//Ignore the update file
		if file.Name() == fileUpdate {
			continue
		}

		//Ignore files that don't start with the notes prefix
		if !strings.HasPrefix(file.Name(), prefixNotes) {
			log.Printf("Skipping unknown file in update directory %s\n", filePath)
			continue
		}

		//Get the language from the file name
		updateNotesLang := strings.TrimPrefix(file.Name(), prefixNotes)
		updateNotesLang = strings.TrimSuffix(updateNotesLang, filepath.Ext(updateNotesLang))

		//Read the file data
		updateNotesData, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}
		updateNotesString := string(updateNotesData)

		//Append the notes slice to the array
		updateNotesSlice = append(updateNotesSlice, UpdateNotes{
			Lang: updateNotesLang,
			Message: updateNotesString,
		})
	}

	return UpdateOutput{
		Data:  updateData,
		Notes: updateNotesSlice,
	}
}