// Package utils provides helper functions and utilities for the bot.
package utils

import (
	"botwa/types"
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"

	"go.mau.fi/libsignal/logger"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

// Global start time untuk runtime bot
var appStartTime = time.Now()

// HandleCommand routes and processes user commands.
func HandleCommand(client *whatsmeow.Client, m types.Messages, evt *events.Message) {
	if m.Prefix == "" {
		return
	}

	switch m.Command {
	case "ping", "uptime":
		jid := evt.Info.Chat
		start := time.Now()

		// Ambil info sistem
		platform := runtime.GOOS
		totalRam := getTotalMemory()
		totalDisk := getTotalDiskSpace()
		cpuCount := runtime.NumCPU()
		uptimeVps := getUptime()
		botUptime := formatDuration(time.Since(appStartTime))
		latency := time.Since(start).Seconds()

		// Format pesan
		msg := fmt.Sprintf(`*🔴 INFORMATION SERVER*

• Platform : %s
• Total Ram : %s
• Total Disk : %s
• Total Cpu : %d Core
• Runtime VPS : %s

*🔵 INFORMATION BOT GOLANG*

• Respon Speed : %.4f detik
• Runtime Bot : %s`,
			platform,
			totalRam,
			totalDisk,
			cpuCount,
			uptimeVps,
			latency,
			botUptime,
		)

		// Kirim pesan ke WhatsApp
		_, err := client.SendMessage(context.Background(), jid, &waProto.Message{
			Conversation: proto.String(msg),
		})
		if err != nil {
			logger.Error("Failed to send uptime reply: " + err.Error())
		}
	}
}

// Fungsi bantu untuk format waktu
func formatDuration(d time.Duration) string {
	seconds := int(d.Seconds())

	days := seconds / 86400
	seconds %= 86400
	hours := seconds / 3600
	seconds %= 3600
	minutes := seconds / 60
	seconds %= 60

	result := ""
	if days > 0 {
		result += fmt.Sprintf("%d hari ", days)
	}
	if hours > 0 {
		result += fmt.Sprintf("%d jam ", hours)
	}
	if minutes > 0 {
		result += fmt.Sprintf("%d menit ", minutes)
	}
	if seconds > 0 {
		result += fmt.Sprintf("%d detik", seconds)
	}
	return result
}

// RAM total
func getTotalMemory() string {
	v, err := mem.VirtualMemory()
	if err != nil {
		return "Unknown"
	}
	return fmt.Sprintf("%.2f GB", float64(v.Total)/1e9)
}

// Disk total
func getTotalDiskSpace() string {
	d, err := disk.Usage("/")
	if err != nil {
		return "Unknown"
	}
	return fmt.Sprintf("%.2f GB", float64(d.Total)/1e9)
}

// Uptime VPS
func getUptime() string {
	uptimeSec, err := host.Uptime()
	if err != nil {
		return "Unknown"
	}
	return formatDuration(time.Duration(uptimeSec) * time.Second)
}
