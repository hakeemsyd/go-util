package Utils

import (
	"fmt"
	"os/exec"
)

// Init input values
type Init struct {
 imageSource, imageDest string
}

// CopyImageFromDockerRepository return script path response
func (e *Init) CopyImageFromDockerRepository() (string) {
    fmt.Println("Started Downloading Image")
    cmd, err := exec.Command("/bin/sh", "copy_image.sh", e.imageSource, e.imageDest).Output()
    if err != nil {
    fmt.Printf("error %s", err)
    }
    output := string(cmd)
    fmt.Println(output)
    return output
}