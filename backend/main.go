package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-contrib/cors"
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
		return "", 0, fmt.Errorf("nvidia-smi command failed: %w", err)
	}

	var gpuTemp string
	var gpuUsage float64
	_, err = fmt.Sscanf(out.String(), "%s,%f", &gpuTemp, &gpuUsage)
	if err != nil {
		return "", 0, fmt.Errorf("failed to parse GPU stats: %w", err)
	}

	return gpuTemp, gpuUsage, nil
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/api/stats", func(c *gin.Context) {
		cpuUsage, err := cpu.Percent(0, false)
		if err != nil {
			log.Printf("Error fetching CPU usage: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get CPU stats"})
			return
		}

		memStats, err := mem.VirtualMemory()
		if err != nil {
			log.Printf("Error fetching memory usage: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get memory stats"})
			return
		}

		diskStats, err := disk.Usage("/")
		if err != nil {
			log.Printf("Error fetching disk usage: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get disk stats"})
			return
		}

		gpuTemp, gpuUsage, err := getGPUStats()
		if err != nil {
			log.Printf("Error fetching GPU stats: %v", err)
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
