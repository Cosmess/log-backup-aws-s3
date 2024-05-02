package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar arquivo .env:", err)
		return
	}

	// Obter os nomes dos buckets do ambiente
	bucketNames := strings.Split(os.Getenv("BUCKET_NAMES"), ",")

	// Configuração do AWS SDK
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	})
	if err != nil {
		fmt.Println("Erro ao criar sessão AWS:", err)
		return
	}
	svc := s3.New(sess)

	// Função para obter informações do último objeto em um bucket
	getLastObjectInfo := func(bucketName string) {
		input := &s3.ListObjectsV2Input{
			Bucket:  aws.String(bucketName),
			MaxKeys: aws.Int64(1),
		}
		result, err := svc.ListObjectsV2(input)
		if err != nil {
			fmt.Println("Erro ao listar objetos:", err)
			return
		}

		if len(result.Contents) > 0 {
			lastObject := result.Contents[len(result.Contents)-1]
			// Convertendo o tamanho para megabytes
			sizeInMB := float64(*lastObject.Size) / (1024 * 1024)
			// Convertendo a data de modificação para o fuso horário de Brasília
			lastModified := lastObject.LastModified.In(time.FixedZone("BRT", -3*60*60)) // UTC-3 (Brasília)
			// Imprimir informações do objeto
			fmt.Println()
			fmt.Printf(color.YellowString("Último objeto em %s:\n"), bucketName)
			fmt.Printf(color.CyanString("Nome: %s\n"), *lastObject.Key)
			fmt.Printf(color.GreenString("Data de Modificação: %s\n"), lastModified.Format("2006-01-02 15:04:05"))
			fmt.Printf(color.BlueString("Tamanho: %.2f MB\n"), sizeInMB)
		} else {
			fmt.Printf(color.RedString("Não há objetos em %s\n"), bucketName)
		}
	}

	// Iterar sobre os nomes dos buckets e obter informações do último objeto
	for _, bucketName := range bucketNames {
		getLastObjectInfo(bucketName)
	}

	// Aguardar até que o usuário pressione Enter para fechar o programa
	fmt.Println("Pressione Enter para sair...")
	fmt.Scanln()
}
