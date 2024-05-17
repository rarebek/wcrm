package v1

import (
	pbu "api-gateway/genproto/product"
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type File struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// @Summary 		Image upload
// @Security 		ApiKeyAuth
// @Description 	Api for image upload
// @Tags 			file-upload
// @Accept 			json
// @Produce 		json
// @Param 			file formData file true "Image"
// @Success 		200 {object} string
// @Failure 		400 {object} string
// @Failure 		500 {object} string
// @Router 			/v1/file-upload [post]
func (h *HandlerV1) UploadImage(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.Config.CtxTimeout))
	defer cancel()

	// MinIO serveriga bog'lanish
	endpoint := "18.158.24.26:9000" // MinIO serveringizning manzili
	accessKeyID := "fnatic1111"     // MinIO serveringizning foydalanuvchi nomi
	secretAccessKey := "12345678"   // MinIO serveringizning maxfiy kodi
	bucketName := "picture"         // Saqlash uchun tanlangan kovatcha (bucket) nomi
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error connecting to MiniIO server")
		log.Println("Error connecting to MiniIO server:", err)
		return
	}

	// Faylni olish
	var file File
	if err := c.ShouldBind(&file); err != nil {
		c.JSON(http.StatusBadRequest, "Error uploading file")
		log.Println("Error uploading file:", err)
		return
	}

	// Faylning tipini tekshirish
	ext := filepath.Ext(file.File.Filename)
	if ext != ".png" && ext != ".jpg" && ext != ".svg" && ext != ".jpeg" {
		c.JSON(http.StatusBadRequest, "Bad request of file image")
		log.Println("Bad request of file image")
		return
	}

	// Faylni minIOga yuklash
	id := uuid.New().String()
	objectName := id + ext
	contentType := "image/jpeg"

	fileReader, err := file.File.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error opening file")
		log.Println("Error opening file:", err)
		return
	}
	defer fileReader.Close()

	_, err = minioClient.PutObject(context.Background(), bucketName, objectName, fileReader, -1, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error uploading file to MinIO server")
		log.Println("Error uploading file to MinIO server:", err)
		return
	}

	// MinIO serverining URL manzili
	minioURL := fmt.Sprintf("https://%s/%s/%s", "media.rarebek.uz", bucketName, objectName)

	id = c.Param("id")

	h.Service.ProductService().UpdateProduct(ctx, &pbu.Product{
		Id:      id,
		Picture: minioURL,
	})

	c.JSON(http.StatusOK, gin.H{
		"url": minioURL,
	})
}
