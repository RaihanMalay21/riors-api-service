package service

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"github.com/RaihanMalay21/api-service-riors/domain"
	"github.com/RaihanMalay21/api-service-riors/middlewares"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func DestinationFolder(pathFolder, nameFile string) string {
	// menentukan path dir untuk file yang akan di create
	filePath := filepath.Join(pathFolder, nameFile)

	return filePath
}

func CreateImage(data *domain.Product, image multipart.File, fileHeader *multipart.FileHeader, ext string) error {
	nameFile := filepath.Base(fileHeader.Filename[:len(fileHeader.Filename)-len(ext)]) // mengambil nama filenya saja
	hasher := sha256.Sum256([]byte(nameFile))                                          // mengkonversi nama file menggunakan sha256 menjadi byte dan ubah menjadi string
	hashingNameImageString := hex.EncodeToString(hasher[:])
	data.Image = hashingNameImageString + strconv.Itoa(int(data.Id)) + ext

	createdPathImage := DestinationFolder("C:\\Users\\acer\\Downloads\\riors-service\\api-service-riors\\assets", data.Image)

	outfile, err := os.Create(createdPathImage)
	if err != nil {
		return err
	}
	defer outfile.Close()

	if _, err := io.Copy(outfile, image); err != nil {
		os.Remove(createdPathImage)
		return err
	}

	return nil
}

func UploadToS3(data *domain.Product, file multipart.File, fileHeader *multipart.FileHeader, ext string, imageType string) error {

	sess, err := middlewares.InitAWSSession()
	if err != nil {
		return err
	}

	// Membuat client s3
	s3Client := s3.New(sess)

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(file)
	if err != nil {
		return fmt.Errorf("gagal membaca file ke buffer: %v", err)
	}

	nameFile := filepath.Base(fileHeader.Filename[:len(fileHeader.Filename)-len(ext)]) // mengambil nama filenya saja
	hasher := sha256.Sum256([]byte(nameFile))                                          // mengkonversi nama file menggunakan sha256 menjadi byte dan ubah menjadi string
	hashingNameImageString := hex.EncodeToString(hasher[:])
	data.Image = hashingNameImageString + strconv.Itoa(int(data.Id)) + ext

	s3Key := fmt.Sprintf("%s/%s", os.Getenv("BUCKET_FOLDER"), data.Image)

	if _, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("BUCKET_NAME")),
		Key:         aws.String(s3Key),
		Body:        bytes.NewReader(buffer.Bytes()),
		ContentType: aws.String(imageType),
		ACL:         aws.String("public-read"),
	}); err != nil {
		return fmt.Errorf("gagal mengunggah file ke S3: %v", err)
	}

	return nil
}
