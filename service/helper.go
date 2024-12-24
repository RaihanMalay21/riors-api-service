package service

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/smtp"
	"os"
	"path/filepath"
	"strconv"
	"time"

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

func GenerateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(90000000) + 10000000

	return randomNumber
}

func EmailVerificationCode(email *string, verificationCode *int) error {
	Auth := smtp.PlainAuth(
		"",
		"cabangbanyak@gmail.com",
		"lnbq rahl xyyg fwcy",
		"smtp.gmail.com",
	)

	// Subjek email
	subject := "Verifikasi Email Anda"

	// Isi body email
	body := fmt.Sprintf(`
		<html>
			<body style="font-family: Arial, sans-serif; line-height: 1.6; color: white; font-size: 12px">
				<p>Halo,</p>
				<p>Terima kasih telah mendaftar di layanan kami. Berikut adalah kode verifikasi Anda:</p>
				<p style="font-size: 18px; font-weight: bold;">Kode Verifikasi: %d</p>
				<p>Masukkan kode ini untuk menyelesaikan proses registrasi Anda.</p>
				<p>Jika Anda tidak merasa mendaftar di layanan kami, abaikan email ini.</p>
				<p>Salam,</p>
				<p>Tim Riors</p>
			</body>
		</html>
	`, *verificationCode)

	// Header MIME untuk mendukung encoding multipart jika ada lampiran
	header := "MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n"

	msg := []byte("To: " + *email + "\r\n" +
		"Subject: " + subject + "\r\n" + header + "\r\n" + body)

	// kirim mesage ke email user
	if err := smtp.SendMail(
		"smtp.gmail.com:587",
		Auth,
		"cabangbanyak@gmail.com",
		[]string{*email},
		msg,
	); err != nil {
		return err
	}

	return nil
}
