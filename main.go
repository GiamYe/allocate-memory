package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/satori/go.uuid"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var (
	//memoryPool              = map[string](*[1024 * 1024]complex128){}
	memoryAllocatedInMB int = 0
	release_interval    int
)

const RELEASE_INTERVAL int = 10

func init() {
	if os.Getenv("RELEASE_INTERVAL") != "" {
		value, err := strconv.Atoi(os.Getenv("RELEASE_INTERVAL"))
		if err != nil {
			panic(err)
		}

		if value <= 0 {
			panic(fmt.Errorf("RELEASE_INTERVAL(%d) should be greater than 0.)", value))
		}

		release_interval = value
	} else {
		release_interval = RELEASE_INTERVAL
	}
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.GET("/_ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.GET("/memory", AllocateQuotaMemory)
	//router.GET("/memory/:size/action/allocate", AllocateCustomMemory)
	router.GET("/cpu", ConsumeCPU)
	router.Run(":8080")
}

func AllocateQuotaMemory(c *gin.Context) {
	//id := uuid.NewV4().String()

	a := [1024 * 1024]complex128{}

	fmt.Println(a)

	time.Sleep(time.Duration(release_interval) * time.Second)

	//delete(memoryPool, id)
	runtime.GC()

	c.String(http.StatusOK, "OK")
}

/*func AllocateCustomMemory(c *gin.Context) {
	size, err := strconv.Atoi(c.Param("size"))
	if err != nil {
		c.String(http.StatusBadRequest, "memory allocate input should be an interger.")
		return
	}

	if size <= 0 {
		c.String(http.StatusBadRequest, "memory allocate input should be larger than 0.")
		return
	}

	for i := 0; i < size; i++ {
		memoryPool = append(memoryPool, &([1024 * 64]complex128{}))
		//runtime.GC()
		memoryAllocatedInMB++
	}

	message := "Allocated about " + strconv.Itoa(memoryAllocatedInMB) + " MB memory."

	c.String(200, message)
}*/

func ConsumeCPU(c *gin.Context) {
	cmd := exec.Command("bash", "-c", "awk 'BEGIN{while (i=1) {}}'")
	if err := cmd.Run(); err != nil {
		c.String(500, err.Error())
	}
}
