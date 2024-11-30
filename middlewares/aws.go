package middlewares

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Fungsi untuk inisialisasi sesi AWS
func InitAWSSession() (*session.Session, error) {
	// Mengambil variabel lingkungan
	awsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	bucketRegion := os.Getenv("BUCKET_REGION")

	if awsAccessKeyID == "" || awsSecretAccessKey == "" || bucketRegion == "" {
		return nil, fmt.Errorf("AWS credentials atau region belum diatur di environment variables")
	}

	// Membuat sesi AWS dengan konfigurasi
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(bucketRegion),
	})
	if err != nil {
		return nil, err
	}

	return sess, nil
}
