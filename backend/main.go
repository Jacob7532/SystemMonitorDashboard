package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type SystemStats struct {
	CPUUsage    float64 `json:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage"`
	DiskUsage   float64 `json:"disk_usage"`
}

func getSystemStats() (SystemStats, error) {
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return SystemStats{}, err
	}

	memoryStats, err := mem.VirtualMemory()
	if err != nil {
		return SystemStats{}, err
	}

	diskStats, err := disk.Usage("/")
	if err != nil {
		return SystemStats{}, err
	}

	return SystemStats{
		CPUUsage:    cpuPercent[0],
		MemoryUsage: memoryStats.UsedPercent,
		DiskUsage:   diskStats.UsedPercent,
	}, nil
}

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/api/stats", func(c *gin.Context) {
		stats, err := getSystemStats()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get system stats"})
			return
		}
		c.JSON(http.StatusOK, stats)
	})

	log.Println("Server starting on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
