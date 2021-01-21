package Utils

import (
	"fmt"
	"os/exec"
)

// Init input values
type Init struct {
 IMAGE_SOURCE, IMAGE_DEST string
}

// CopyImageFromDockerRepository return script path response
func CopyImageFromDockerRepository(x string, y string) (string) {
    fmt.Println("Started Downloading Image.")
    fmt.Println(e.IMAGE_SOURCE)
    fmt.Println(e.IMAGE_DEST)
    cmd, err := exec.Command("/bin/sh", "copy_image.sh", e.IMAGE_SOURCE , e.IMAGE_DEST).Output()
    if err != nil {
    fmt.Printf("error %s", err)
    }
    output := string(cmd)
    fmt.Println(output)
    return output
}