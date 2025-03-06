package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

// monitor status
func monitorPrintStats(crackedCount, linesProcessed, totalHashesGenerated *int32, stopChan <-chan struct{}, startTime time.Time, totalHashCount int, wg *sync.WaitGroup, interval int) {
	var ticker *time.Ticker
	if interval > 0 {
		ticker = time.NewTicker(time.Duration(interval) * time.Second)
		defer ticker.Stop()
	}

	for {
		select {
		case <-stopChan:
			// print final stats and exit
			printStats(time.Since(startTime), int(atomic.LoadInt32(crackedCount)), totalHashCount, int(atomic.LoadInt32(linesProcessed)), true, atomic.LoadInt32(totalHashesGenerated))
			wg.Done()
			return
		case <-func() <-chan time.Time {
			if ticker != nil {
				return ticker.C
			}
			return nil
		}():
			if interval > 0 {
				printStats(time.Since(startTime), int(atomic.LoadInt32(crackedCount)), totalHashCount, int(atomic.LoadInt32(linesProcessed)), false, atomic.LoadInt32(totalHashesGenerated))
			}
		}
	}
}

// printStats
func printStats(elapsedTime time.Duration, crackedCount int, totalHashCount, linesProcessed int, exitProgram bool, totalHashesGenerated int32) {
	hours := int(elapsedTime.Hours())
	minutes := int(elapsedTime.Minutes()) % 60
	seconds := int(elapsedTime.Seconds()) % 60
	hashesPerSecond := float64(atomic.LoadInt32(&totalHashesGenerated)) / elapsedTime.Seconds()
	log.Printf("Cracked: %d/%d %.2f h/s %02dh:%02dm:%02ds", crackedCount, totalHashCount, hashesPerSecond, hours, minutes, seconds)

	if exitProgram {
		fmt.Println("")
		time.Sleep(100 * time.Millisecond)
		os.Exit(0) // exit only if exitProgram bool
	}
}
