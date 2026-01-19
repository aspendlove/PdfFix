package pdf

import (
	"fmt"
	"io"
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
	if err != nil {
		c.JSON(500, gin.H{"error": "Processing failed"})
		return
	}
	defer os.Remove(dst)
	defer os.Remove(outputFile)

	out, err := os.Open(outputFile)
	if err != nil {
		c.JSON(500, gin.H{"error": "Cannot read output file"})
		return
	}
	defer out.Close()

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", fmt.Sprintf("inline; filename=fixed_%s.pdf", file.Filename))
	c.Header("Content-Type", "application/pdf")
	
	if _, err := io.Copy(c.Writer, out); err != nil {
		fmt.Printf("Error writing file: %v\n", err)
	}
}

func rasterize(filepath string) (string, error) {
	outputFile := uuid.New().String() + ".pdf"
	args := []string{
		"./pdfTmp", filepath, outputFile,
	}
	cmd := exec.Command(
		"./scripts/pdfFix.sh", args...,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Script output: %s\n", string(output))
		return "", fmt.Errorf("Could not rasterize pdf: %w", err)
	}
	return outputFile, nil
}
