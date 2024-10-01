package app

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

// Función para obtener el nombre del host
func GetHostName() string {
	hostName, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error al obtener el nombre del host: %v", err)
	}
	return hostName
}

// Función que obtiene una lista de imágenes en la carpeta
func GetImagesInFolder(folderPath string) ([]string, error) {
	// Leer el contenido del directorio
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	// Filtros de extensiones de imágenes permitidas
	extensions := []string{".png", ".jpg", ".jpeg"}
	var images []string

	// Recorrer los archivos y agregar los que tienen las extensiones válidas
	for _, file := range files {
		if !file.IsDir() && isImage(file.Name(), extensions) {
			images = append(images, file.Name())
		}
	}
	return images, nil
}

// Función para verificar si un archivo tiene una extensión de imagen permitida
func isImage(fileName string, extensions []string) bool {
	ext := strings.ToLower(filepath.Ext(fileName))
	for _, validExt := range extensions {
		if ext == validExt {
			return true
		}
	}
	return false
}

// Función para seleccionar imágenes al azar
func GetRandomImages(images []string, numImages int) []string {
	perm := rand.Perm(len(images)) // Genera una permutación aleatoria de índices
	selected := make([]string, numImages)
	for i := 0; i < numImages; i++ {
		selected[i] = images[perm[i]]
	}
	return selected
}

// Función para codificar una imagen en Base64 y devolverla
func GetEncodeImageToBase64(imagePath string) (string, error) {
	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return "", err
	}

	// Codificar la imagen en Base64 y devolverla
	encoded := base64.StdEncoding.EncodeToString(imageData)
	return encoded, nil
}
