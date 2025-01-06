package helper

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
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
	"github.com/joho/godotenv"
)

type HelperService struct{}

func NewHelperService() *HelperService {
	return &HelperService{}
}

func (hs *HelperService) DestinationFolder(pathFolder, nameFile string) string {
	// menentukan path dir untuk file yang akan di create
	filePath := filepath.Join(pathFolder, nameFile)

	return filePath
}

func (hs *HelperService) CreateImage(data *domain.Product, image multipart.File, fileHeader *multipart.FileHeader, ext string) error {
	nameFile := filepath.Base(fileHeader.Filename[:len(fileHeader.Filename)-len(ext)]) // mengambil nama filenya saja
	hasher := sha256.Sum256([]byte(nameFile))                                          // mengkonversi nama file menggunakan sha256 menjadi byte dan ubah menjadi string
	hashingNameImageString := hex.EncodeToString(hasher[:])
	data.Image = hashingNameImageString + strconv.Itoa(int(data.Id)) + ext

	createdPathImage := hs.DestinationFolder("C:\\Users\\acer\\Downloads\\riors-service\\api-service-riors\\assets", data.Image)

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

func (hs *HelperService) UploadToS3(data *domain.Product, file multipart.File, fileHeader *multipart.FileHeader, ext string, imageType string) error {

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

func (hs *HelperService) UploadToS3Admin(data *domain.Employee, file multipart.File, fileHeader *multipart.FileHeader, ext string, imageType string) error {

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

	s3Key := fmt.Sprintf("%s/%s", os.Getenv("BUCKET_FOLDER_ADMIN"), data.Image)

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

func (hs *HelperService) GenerateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(90000000) + 10000000

	return randomNumber
}

func (hs *HelperService) SendEmailForgotPassword(email, username, link string) error {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	Auth := smtp.PlainAuth(
		"",
		smtpUser,
		smtpPass,
		"smtp.gmail.com",
	)

	subject := "Change Password"

	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Reset Your Password</title>
			<style>
				body { font-family: Arial, sans-serif; line-height: 1.6; }
				.container { max-width: 600px; margin: 20px auto; padding: 20px; border: 1px solid #ddd; border-radius: 8px; }
				.btn { background-color: #007BFF; color: white; text-decoration: none; padding: 12px 20px; border-radius: 5px; display: inline-block; }
				.footer { font-size: 0.9em; color: #555; margin-top: 20px; }
			</style>
		</head>
		<body>
			<div class="container">
				<h2>Reset Your Password</h2>
				<p>Hi %s,</p>
				<p>We received a request to reset your password for your account. Please click the button below to reset your password:</p>
				<p><a href="%s" class="btn">Reset Password</a></p>
				<p>If you did not request this, you can safely ignore this email. This link will expire in %d minute for security reasons.</p>
				<p>Thank you,<br>The Support Team</p>
				<div class="footer">
					<p>Need help? Contact us at <a href="mailto:riors@gmail.com">riors@gmail.com</a></p>
				</div>
			</div>
		</body>
		</html>
	`, username, link, 5)

	header := "MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n"

	msg := []byte("To: " + email + "\r\n" +
		"Subject: " + subject + "\r\n" + header + "\r\n" + body)

	// kirim mesage ke email user
	if err := smtp.SendMail(
		"smtp.gmail.com:587",
		Auth,
		smtpUser,
		[]string{email},
		msg,
	); err != nil {
		return err
	}

	return nil
}

func (hs *HelperService) SendEmailVerificationCode(email *string, verificationCode *int) error {
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	Auth := smtp.PlainAuth(
		"",
		smtpUser,
		smtpPass,
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
		smtpUser,
		[]string{*email},
		msg,
	); err != nil {
		return err
	}

	return nil
}

func (hs *HelperService) ConvertDateStringToTime(date string, response map[string]interface{}) time.Time {
	layout := "2006-01-02"
	dateParse, err := time.Parse(layout, date)
	if err != nil {
		response["error"] = err.Error()
		return time.Time{}
	}
	return dateParse
}
