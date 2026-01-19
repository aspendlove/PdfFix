package pdf

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SubmitHandler(c *gin.Context) {
	const MaxSize = 10 * 1024 * 1024
	file, err := c.FormFile("pdf")
	if err != nil {
		c.JSON(500, gin.H{"error": "No file"})
		return
	}

	if file.Size > MaxSize {
		c.JSON(500, gin.H{"error": "File too large (10mb max)"})
		return
	}

	f, _ := file.Open()
	defer f.Close()

	header := make([]byte, 5)
	f.Read(header)
	if string(header) != "%PDF-" {
		c.JSON(500, gin.H{"error": "Not a pdf"})
		return
	}

	dst := "uploads/" + uuid.New().String() + ".pdf"
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(500, gin.H{"error": "Cannot save file"})
		return
	}
	c.JSON(200, gin.H{"success": "File uploaded successfully"})
}

func rasterize(filepath string) (string, error) {
	
	
	return "", nil
}
