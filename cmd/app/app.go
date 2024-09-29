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

// func main() {
// 	// Verificar si se pasó un argumento para el directorio
// 	if len(os.Args) < 2 {
// 		log.Fatalf("Error %s <directorio de imágenes>", os.Args[0])
// 	}

// 	// Obtener el directorio de imágenes desde los argumentos de la línea de comandos
// 	folderPath := os.Args[1]

// 	// Imprimir el nombre del host
// 	printHostName()

// 	// Obtener las imágenes de la carpeta
// 	images, err := getImagesInFolder(folderPath)
// 	if err != nil {
// 		log.Fatalf("Error al leer la carpeta: %v", err)
// 	}

// 	// Imprimir la cantidad de imágenes encontradas
// 	fmt.Printf("Cantidad de imágenes encontradas: %d\n", len(images))

// 	// Seleccionar al azar tres imágenes si hay suficientes
// 	if len(images) > 0 {
// 		rand.Seed(time.Now().UnixNano())

// 		// Si hay menos de tres imágenes, seleccionarlas todas
// 		numImagesToSelect := 3
// 		if len(images) < 3 {
// 			numImagesToSelect = len(images)
// 		}

// 		// Seleccionar al azar `numImagesToSelect` imágenes
// 		selectedImages := selectRandomImages(images, numImagesToSelect)

// 		// Codificar y mostrar las imágenes seleccionadas
// 		for _, imageName := range selectedImages {
// 			imagePath := filepath.Join(folderPath, imageName)
// 			fmt.Printf("Imagen seleccionada al azar: %s\n", imageName)

// 			// Codificar la imagen en Base64 y almacenarla en una variable
// 			encodeImageToBase64(imagePath)
// 			if err != nil {
// 				log.Fatalf("Error al codificar la imagen: %v", err)
// 			}

// 			// Mostrar la imagen codificada
// 			//fmt.Printf("Imagen codificada en Base64: %s\n", encoded)
// 		}
// 	} else {
// 		fmt.Println("No se encontraron imágenes en la carpeta.")
// 	}
// }

// Función para obtener y mostrar el nombre del host
// func printHostName() {
// 	hostName, err := os.Hostname()
// 	if err != nil {
// 		log.Fatalf("Error al obtener el nombre del host: %v", err)
// 	}
// 	fmt.Printf("Nombre del host: %s\n", hostName)
// }

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
