// this package is used to control the endpoints we would be sending our image to.
package controller

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/joho/godotenv"
)

// A structure that accepts image amd sends to the api
type ImageData struct {
	Image string `json:"image"`
	//Anytime the image is sent to the endpoint, it is sent as json text to the ImageController function.
}

// this structs contains the response we would be getting from AWS rekognition
type Analysis struct {
	LabelModelVersion string `json:"labelmodelversion"`
	// this label contains it's own object
	Labels []struct {
		Confidence float64 `json:"confidence"`
		Name       string  `json:"name"`
		Parents    []struct {
			Name string `json:"name"`
		}
	}
}

// A function that handles AWS image Recognition service. It returns all the response gotten from the API in a JSON format, that can then be passed to the frontend.
func imageIdentifier(image string) Analysis {
	//accepts image in bit64

	// get our access key and secret key from thr env file using the godotenv package.
	godotenv.Load()

	sess := session.New(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("ACCESS_KEY_ID"),
			os.Getenv("SECRET_KEY"),
			"",
		),
	})
	svc := rekognition.New(sess)
	//Decode at 64bit image for image detection.
	decodedImage, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		fmt.Println(err)
	}
	// sending a request to AWS rekognition service
	input := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image{
			//sending the image to the rekognition service
			Bytes: decodedImage,
		},
	}

	// Result gotten from the request sent to the rekognition service.
	result, err := svc.DetectLabels(input)
	if err != nil {
		fmt.Println(err)
	}

	//marshal our result first.
	output, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(output)
	// Create an Analysis struct. It is a structure of the json gotten from the image rekognition service
	var responseData Analysis
	//We unmarshal what we got inside the struct.
	if err := json.Unmarshal(output, &responseData); err != nil {
		panic(err)
	}
	// returning our response.
	return responseData
}

func ImageController(w http.ResponseWriter, r *http.Request) {
	//This function controls what is going to happen when the  image is sent to it from the server.
	var imageData ImageData

	//Decoding any data coming from imageData (which is also our struct with the json)
	err := json.NewDecoder(r.Body).Decode(&imageData)
	if err != nil {
		fmt.Println(err, "error reading payload")
		w.WriteHeader(http.StatusBadRequest) //StatusBadRequest = 400.
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	// called our imageIdentifier function to identify image sent from the client.
	var recievedData = imageIdentifier(imageData.Image)
	json.NewEncoder(w).Encode(recievedData) // returning received data as json format.
}
