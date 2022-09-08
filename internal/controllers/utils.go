package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sharif-starter-backend/internal/models"
	"strconv"
)

func UploadProjectImage(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)
	id := r.Form.Get("projectId")
	project := &models.Project{}
	db.Where("id = ?", id).First(project)
	address, err := uploadFile(w, r, project.ID)
	fmt.Println(address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	project.ProjectImage = address
	db.Save(project)
}

func uploadFile(w http.ResponseWriter, r *http.Request, id uint) (string, error) {
	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return "", err
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create folder
	path := "files/image/" + strconv.Itoa(int(id))
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Println(err)
		return "", err
	}
	// Create file
	dst, err := os.Create(path + "/" + handler.Filename)
	defer dst.Close()
	if err != nil {
		return "", err
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		return "", err

	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
	return path + "/" + handler.Filename, nil
}
