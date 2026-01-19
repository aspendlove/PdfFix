package pdf

import (
	"fmt"
	"os"
	"os/exec"

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
	outputFile, err := rasterize(dst)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", fmt.Sprintf("inline; filename=fixed_%s.pdf", file.Filename))
	c.Header("Content-Type", "application/pdf")
	c.File(outputFile)

	defer os.Remove(dst)
	defer os.Remove(outputFile)
}

func rasterize(filepath string) (string, error) {
	outputFile := uuid.New().String() + ".pdf"
	args := []string{
		"./pdfTmp", filepath, outputFile,
	}
	cmd := exec.Command(
		"./scripts/pdfFix.sh", args...,
	)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("Could not rasterize pdf: %w", err)
	}
	return outputFile, nil
}
