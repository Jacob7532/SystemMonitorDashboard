package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func getGPUStats() (string, float64, error) {
	cmd := exec.Command("nvidia-smi", "--query-gpu=temperature.gpu,utilization.gpu", "--format=csv,noheader,nounits")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", 0, err
	}

	var gpuTemp string
	var gpuUsage float64
	fmt.Sscanf(out.String(), "%s,%f", &gpuTemp, &gpuUsage)

	return gpuTemp, gpuUsage, nil
}

func main() {
	r := gin.Default()

	r.GET("/api/stats", func(c *gin.Context) {
		cpuUsage, _ := cpu.Percent(0, false)
		memStats, _ := mem.VirtualMemory()
		diskStats, _ := disk.Usage("/")

		gpuTemp, gpuUsage, err := getGPUStats()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get GPU stats"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"cpu_usage":  cpuUsage[0],
			"mem_usage":  memStats.UsedPercent,
			"disk_usage": diskStats.UsedPercent,
			"gpu_temp":   gpuTemp,
			"gpu_usage":  gpuUsage,
		})
	})

	r.Run(":8080") // Start the server on port 8080
}
