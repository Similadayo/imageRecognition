# Simi's-Image-Recognizer

An image recognition system written in Go for image identification and details. 

## About
This is a go programming language with a gorilla web toolkit that provides useful, composable packages for writing HTTP-based applications. The libary used in is the [AWS Rekognition service](http://github.com/aws/aws-sdk-go/service/rekognition) and [AWS SDK go](http://github.com/aws/aws-sdk-go). It takes a 64-bit image and gives the full details of the object found in the image.

## Prerequesites
Make sure you have go installed on your pc. you can get the functional set up by following the steps [here](http://https://go.dev/learn/).

## Usage 
Change into project directory containing main.go.

To start the server: 
```
go run main.go
```
Then you proceed to postman for API and put in this URL:
```
localhost:4000/api/image-identifier
```
After with you change you method to a `POST` method and in the body tag you select `raw` and you toggle to `JSON`. Then you input the base64 code like this 
```
{
"image": "base64 code"
}
```
Then you get all the details from your image. 
