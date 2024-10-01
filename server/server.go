package main

import (
	"fmt"
	"imageServer/cmd/app"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

const PORT string = ":8080"

func main() {
	// Hacemos una nueva instancia de ServeMux
	mux := http.NewServeMux()

	// Verificar si se pasó un argumento para el directorio
	if len(os.Args) < 2 {
		log.Fatalf("Error %s <directorio de imágenes>", os.Args[0])
	}
	// Obtener el Hostname
	hostName := app.GetHostName()
	fmt.Printf("Nombre del host: %s\n", hostName)

	type ImageData struct {
		Name      string
		Base64Img string
	}

	// Ruta para renderizar el template HTML
	mux.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {

		// Obtener el directorio de imágenes desde los argumentos de la línea de comandos
		folderPath := os.Args[1]

		// Obtener las imágenes de la carpeta
		images, err := app.GetImagesInFolder(folderPath)
		if err != nil {
			http.Error(w, "Error al leer la carpeta de imágenes", http.StatusInternalServerError)
			return
		}

		// Variables para las imágenes seleccionadas
		var selectedImages []ImageData

		// Seleccionar al azar tres imágenes si hay suficientes
		if len(images) > 0 {

			// Generar un número aleatorio entre 3 y 4 para la cantidad de imágenes
			numImagesToSelect := 3 + rand.Intn(2) // rand.Intn(2) genera 0 o 1, por lo tanto 3+0=3 o 3+1=4

			// Si hay menos imágenes que el número seleccionado, ajustarlo
			if len(images) < numImagesToSelect {
				numImagesToSelect = len(images)
			}

			// Seleccionar al azar numImagesToSelect imágenes
			randomImages := app.GetRandomImages(images, numImagesToSelect)

			// Codificar y almacenar las imágenes seleccionadas
			for _, imageName := range randomImages {
				imagePath := filepath.Join(folderPath, imageName)
				fmt.Printf("Imagen seleccionada al azar: %s\n", imageName)

				// Codificar la imagen en Base64
				base64Img, err := app.GetEncodeImageToBase64(imagePath)
				if err != nil {
					log.Fatalf("Error al codificar la imagen: %v", err)
				}

				// Agregar a la lista selectedImages
				selectedImages = append(selectedImages, ImageData{
					Name:      imageName,
					Base64Img: base64Img,
				})
			}
		} else {
			fmt.Println("No se encontraron imágenes en la carpeta.")
		}

		data := map[string]interface{}{
			"HostName": hostName,
			"Tema":     "anime",
			"Images":   selectedImages, // Pasar la lista de ImageData
		}

		tmpl := template.Must(template.ParseFiles("./Views/index.html"))
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Error al renderizar el template", http.StatusInternalServerError)
			return
		}
	})

	//Ruta para servir las imágenes estáticas desde la carpeta local
	imageDir := "../IMG"
	fileServer := http.FileServer(http.Dir(imageDir))

	// Manejar la ruta "/image" para servir tanto el HTML como las imágenes
	mux.Handle("/image/", http.StripPrefix("/image/", fileServer))

	// Ruta para servir los archivos CSS desde la carpeta local
	cssDir := "./Style"
	cssFileServer := http.FileServer(http.Dir(cssDir))
	mux.Handle("/Style/", http.StripPrefix("/Style/", cssFileServer))

	fmt.Printf("Servidor ejecutándose en el puerto %s\n", PORT)

	// Iniciar el servidor en el puerto especificado
	log.Fatal(http.ListenAndServe(PORT, mux))
}
