// Package utils provides helper functions and utilities for the bot.
package utils

import (
	"botwa/types"
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"

	"go.mau.fi/libsignal/logger"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

// Global start time untuk runtime bot
var appStartTime = time.Now()

//OWNER BOT
const ownerJID = "6281327393959:52@s.whatsapp.net" // Ganti sesuai device JID kamu

// HandleCommand routes and processes user commands.
func HandleCommand(client *whatsmeow.Client, m types.Messages, evt *events.Message) {
	if m.Prefix == "" {
		return
	}

	switch m.Command {
	//--------CASE MENU-------//
	case "menu":
		jid := evt.Info.Chat

		// Format pesan menu
		menuText := `*📋 DAFTAR MENU BOT:*

• *.ping* – Cek status server dan bot
• *.jpm* – Promosi All Group

Silakan ketik salah satu perintah di atas.`

		// Kirim pesan menu ke WhatsApp
		_, err := client.SendMessage(context.Background(), jid, &waProto.Message{
			Conversation: proto.String(menuText),
		})
		if err != nil {
			logger.Error("Failed to send menu: " + err.Error())
		}
//--------CASE JPM-------//
case "jpm":
	jid := evt.Info.Chat
	sender := evt.Info.Sender.String()
//fmt.Println("Sender JID:", sender) // debug


	// Cek apakah pengirim bukan owner
	if sender != ownerJID {
		_, _ = client.SendMessage(context.Background(), jid, &waProto.Message{
			Conversation: proto.String("❌ Perintah ini hanya untuk owner."),
		})
		return
	}

	// Cek apakah ada teks yang dikirim
	if m.Text == "" {
		_, _ = client.SendMessage(context.Background(), jid, &waProto.Message{
			Conversation: proto.String("Contoh: .jpm Halo semua!"),
		})
		return
	}

	// Ambil semua grup yang sedang diikuti
	allGroups, err := client.GetJoinedGroups()
	if err != nil {
		logger.Error("Gagal mengambil grup: " + err.Error())
		return
	}

	totalSent := 0
	messageText := m.Text

	// Info awal
	_, _ = client.SendMessage(context.Background(), jid, &waProto.Message{
		Conversation: proto.String(fmt.Sprintf("Memproses *jpm* ke %d grup...", len(allGroups))),
	})

	for _, group := range allGroups {
		// Kirim pesan ke grup
		_, err := client.SendMessage(context.Background(), group.JID, &waProto.Message{
			Conversation: proto.String(messageText),
		})
		if err == nil {
			totalSent++
		}

		// Delay antar pengiriman (misal 5 detik)
		time.Sleep(3 * time.Second)
	}

	// Info akhir
	_, _ = client.SendMessage(context.Background(), jid, &waProto.Message{
		Conversation: proto.String(fmt.Sprintf("*JPM Selesai ✅*\nTotal grup yang berhasil dikirimi pesan: %d", totalSent)),
	})

//--------CASE PING-------//	
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

*🔵 INFORMATION GOLANG BOT*

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
