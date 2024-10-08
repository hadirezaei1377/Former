package workers

import (
	"fmt"
	"log"

	"former/internal/services"

	"github.com/benmanns/goworker"
)

func InitWorker() {
	options := goworker.WorkerSettings{
		URI:            "redis://localhost:6379/",
		Connections:    100,
		Queues:         []string{"compress_queue"},
		UseNumber:      true,
		ExitOnComplete: false,
	}
	goworker.SetSettings(options)

	goworker.Register("Compress", Compress)
}

func Compress(queue string, args ...interface{}) error {
	filePath := args[0].(string)
	fileName := args[1].(string)

	compressedFilePath, err := services.CompressImage(filePath, fileName, 80)
	if err != nil {
		return err
	}

	fmt.Printf("Compressed file saved to: %s\n", compressedFilePath)
	return nil
}

func StartWorkers() {
	if err := goworker.Work(); err != nil {
		log.Fatal(err)
	}
}
