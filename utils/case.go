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

// OWNER BOT
const ownerJID = "6281327393959" // Ganti sesuai device JID kamu

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
• *.jpmall* – Promosi All List
• *.jpmvpn* – Promosi List VPN
• *.jpmvps* – Promosi List VPS
• *.jpmdor* – Promosi List PAKET

Silakan ketik salah satu perintah di atas.`

		// Kirim pesan menu ke WhatsApp
		_, err := client.SendMessage(context.Background(), jid, &waProto.Message{
			Conversation: proto.String(menuText),
		})
		if err != nil {
			logger.Error("Failed to send menu: " + err.Error())
		}
	// --------CASE JPM-------//
	case "jpmall":
		jid := evt.Info.Chat
		sender := evt.Info.Sender.User

		if sender != ownerJID {
			_, _ = client.SendMessage(context.Background(), jid, &waProto.Message{
				Conversation: proto.String("❌ Perintah ini hanya untuk owner."),
			})
			return
		}

		allGroups, err := client.GetJoinedGroups()
		if err != nil {
			logger.Error("Gagal mengambil grup: " + err.Error())
			return
		}

		totalSent := 0
		messageText := `🔰👑𝗣𝗧 𝗥𝗔𝗝𝗔 𝗦𝗘𝗩𝗘𝗥 𝗣𝗥𝗘𝗠𝗜𝗨𝗠👑🔰
┌──────────────────┐
🥇 𝙱𝚈 𝙰𝚁𝙸 𝚂𝚃𝙾𝚁𝙴 🥇
└──────────────────┘
┌──────────────────┐
🔥 Semua transaksi otomatis via Bot Telegram
✅ Gak perlu ngantri
✅ 24 Jam non-stop
✅ Langsung prose
💻 TELE BOT VPN : t.me/ARI_VPN_STORE_bot
💻 TELE BOT SC TUNNEL : t.me/ARISCTUNNEL_bot
└──────────────────┘
✅ WA : wa.me/6281327393959
✅ TELE : t.me/ARI_VPN_STORE
⏰ OPEN MANUAL : 08.00 00.00
☑ REDY AKUN ARTUN VPN PREMIUM
☑ REDY AKUN YT PREMIUM
☑ REDY AKUN CANVA PREMIUM
☑ REDY VPN PREMIUM
☑ REDY VPS DIGITALOCEAN
☑ REDY SCRIPT TUNNELING
☑ REDY PANEL BOT
☑ REDY PAKET-INJECT
☑ REDY AKUN PAYPAL VERIF
☑ REDY VCC VISA
✅ ORDER : [wa.me/6281327393959]
✅ PRODUK SAYA : [wa.me/c/6281327393959]
✅ CH WA : [whatsapp.com/channel/0029VaOJyEM1XquV2zL3eM0I]
✅ All GC WA : [https://chat.whatsapp.com/EkVrfFu4W8XEmgdF4oqZST?mode=ac_t]
✅ Testimoni & All GC TELE : [t.me/RAJA_VPN_STORE]
✅ 100% Amanah 🙏
✅ SUDAH TERPERCAYA 😉
┌──────────────────┐
    # LIST-VPN
PILIH MODE INJECT NYA
1️⃣ SSH WEBSOCKET Account
2️⃣ SSH SelowDNS Account
3️⃣ OPEN SSH Account 
4️⃣ SSH UDP Account
5️⃣ Xray/Vmess Account
6️⃣ Xray/Vless Account
7️⃣ Trojan Account
  ─────────────────── 
🇸🇬 SERVER SINGAPURA BUKA ✓
📝 LIST HARGA & SERVER :
ISP : DigitalOcean, LLC
CPU : DO-Premium-AMD (4) @ 1.996GHz
RAM : 8 GB / vCPUs : 4 / SSD : 320 GB Disk
Bandwidth : 300GB/BLN
🔹 🇸🇬SG VIP [MAX 2 DEVICE] :
1. Rp 1.750 (3 HARI)
2. Rp 2.000 (7 HARI)
3. Rp 3.500 (15 HARI)
4. Rp 5.000 (22 HARI)
5. Rp 7.000 (30 HARI)
6. Rp 15.000 (60 HARI)
   ───────────────────
ISP : DigitalOcean, LLC 🔥
CPU : DO-Premium-AMD (4) @ 1.996GHz
RAM : 8 GB / vCPUs : 4 / SSD : 320 GB Disk
Bandwidth : 500GB/BLN
[OPEN 3 SERVER]
🔹 🇸🇬SG VVIP [SUPPORT STB] :
1. Rp 2.000 (3 HARI)
2. Rp 3.500 (7 HARI)
3. Rp 5.000 (15 HARI)
4. Rp 7.000 (22 HARI)
5. Rp 10.000 (30 HARI)
6. Rp 20.000 (60 HARI)
   ───────────────────
ISP : DigitalOcean, LLC 
CPU : DO-Premium-AMD (4) @ 1.996GHz
RAM : 8 GB / vCPUs : 4 / SSD : 320 GB Disk
Bandwidth : 200GB/BLN
[OPEN 2 SERVER]
🔹 🇸🇬SG RESSEL [MAX 1 DEVICE] :
1. Rp 5.000 (30 HARI)
2. Rp 10.000 (60 HARI)
   ───────────────────
🇮🇩 SERVER INDONESIA BUKA ✓
📝 LIST HARGA & SERVER :
ISP : Media Antar Nusa PT.
CPU : Intel (Broadwell, IBRS) (1) @ 2.299GHz 
RAM : 2 GB / CPU : 1 / 40 GB SSD Disk
Bandwidth : Unlimited
[OPEN 2 SERVER]
🔹 🇮🇩ID NUSA [MAX 2 DEVICE] :
1. Rp 3.500 (3 HARI)
2. Rp 5.000 (7 HARI)
3. Rp 7.000 (15 HARI)
4. Rp 10.000 (22 HARI)
5. Rp 12.000 (30 HARI)
6. Rp 24.000 (60 HARI)
   ───────────────────
ISP : Media Antar Nusa PT. 
CPU : KVM RHEL 7.6.0 PC (1) @ 2.699GHz   
RAM : 2 GB / 1 CPU / 20 GB SSD Disk
Bandwidth : Unlimited
🔹 🇮🇩ID NUSA [SUPPORT STB] :
1. Rp 5.000 (3 HARI)
2. Rp 7.000 (7 HARI)
3. Rp 10.000 (15 HARI)
4. Rp 12.000 (22 HARI)
5. Rp 14.000 (30 HARI)
6. Rp 25.000 (60 HARI)
   ───────────────────
• MINTA CONFIG MENTAHAN Rp 1.000
• CONFIG PREMIUM DARI SAYA GRATIS
✅ BISA TRIAL DULU
   ───────────────────
KELEBIHAN BELI AKUN PREMIUM
🔹 SUPPORT STB OPENWRT ✓
🔹 Nonton YouTube Lancar 4K ✓
🔹 Server Tidak Gampang Error ✓
🔹 Buat Download Fast Speed ✓
🔹 ANTI RECONNECT ✓
🔹 Support Video Call (GANTI UDP 7100-7900) ✓
🔹 Support GAME (TERGANTUNG TKP) ✓
🔹 Koneksi Stabil ✓
🔹 Fast Connect ✓
🔹 Garansi 100% ✓
🔹 Bonus Config ✓
🔹 DLL
   ───────────────────
⚙ PENGATURAN PEMBELI ⚙
🚫 NO DDOS !!!
🚫 NO HACKING !!!
🚫 NO CARDING !!!
🚫 NO TORRENT !!!
🚫 NO SPAMMING !!!
🚫 NO PLAYING PLAYSTATION !!!
🚫 MAX 2 Device !!!
 ❗ Melanggar AUTO BAN ❗
└──────────────────┘
┌──────────────────┐
   ### LIST-VPS
  ───────────────────
★ LIST VPS DIGITALOCEAN ★

READY VPS | DISKON 20%

2️⃣ 1CPU / 1GB / 25GB SSD / 1TB BW 
• Tanpa SC : 20K 
• SC Arisctunnel v4 : 25K 
• SC Potato : 35K

2️⃣ 1CPU / 2GB / 50GB SSD / 2TB BW 
• Tanpa SC : 35K 
• SC Arisctunnel v4 : 40K 
• SC Potato : 50K

3️⃣ 2CPU / 2GB / 50GB SSD / 3TB BW 
• Tanpa SC : 45K 
• SC Arisctunnel v4 : 50K 
• SC Potato : 60K

4️⃣ 2CPU / 4GB / 80GB SSD / 4TB BW 
• Tanpa SC : 65K 
• SC Arisctunnel v4 : 70K 
• SC Potato : 80K

5️⃣ 4CPU / 8GB / 160GB SSD / 5TB BW 
• Tanpa SC : 85K 
• SC Arisctunnel v4 : 90K 
• SC Potato : 100K 
• SC RDP : 100K 
• SC Pterodactyl : 100K

INFORMASI GARANSI 
- RAM 1 NO GARANSI
- RAM 2-8 GARANSI 1x Replace

 TOS / SYARAT KETENTUAN : 
- Dilarang CPU 100%
- Dilarang DDoS
- Dilarang Torrent
- Dilarang Hacking
- Dilarang Mining
- Dilarang Power Off
- Dilarang Konten Ilegal

INFO:
- Free Domain Premium 
- Free Rebulid (Instal ulang)
- Full SC, Tinggal Pakai 
- Bisa Jualan Panel Bot 
- Bisa Jualan RDP & XRDP 
- Bisa Jualan RDP Forex
- Bisa Buat Website Hosting
- Jualan SSH, XRAY, Trojan 
- Bisa Ganti Script lain 

MINAT?
• Profil: [bit.ly/m/ADMIN-ARISTORE] 
• Testi: [t.me/RAJA_VPN_STORE] 
• Amanah & Terpercaya
└──────────────────┘
┌──────────────────┐
   ### LIST-PAKET-INJECT
└──────────────────┘
┌───────────┐
PROMO PAKET INJECT
🔥XC 1+1GB + XUTSuper ( Rp. 23.000 )
🔥XC 1+1GB + XUTSuper + CONFIG ( Rp. 30.000 )
・ Unlimited Whatsapp
・ Line
・ Gojek
・ Instagram
・ Facebook
Total FUP: 230GB-400GB

🔥XCS + XUTBasic ( Rp. 34.000 )
🔥XCS + XUTBasic + CONFIG ( Rp. 41.000 )
・ Unlimited Whatsapp
・ Line
・ Gojek
Total FUP: 230GB-400GB

🔥XCS + XUTStandard ( Rp. 36.000 )
🔥XCS + XUTStandard + CONFIG ( Rp. 43.000 )
・ Unlimited Whatsapp
・ Line
・ Gojek
・ Facebook
Total FUP: 230GB-400GB

🔥XCS + XUTPremiun ( Rp. 38.000 )
🔥XCS + XUTPremiun + CONFIG ( Rp. 45.000 )
・ Unlimited Whatsapp
・ Line
・ Gojek
・ Instagram
・ Facebook
・ Youtube
Total FUP: 230GB-400GB

🔥 XCS + Multi Add On ( Rp 40.000 )
🔥 XCS + Multi Add On + CONFIG ( Rp 47.000 )
・ XUTBasic
・ XUTStandard
・ XUTSuper
・ XUTPremium
・ XUTYouTube
・ XUTJoox
・ XUTViu
・ XUTNetflix
・ XUTTikTok
Total FUP: 1TB++

Notes :
・Kartu tidak boleh ada
paket xcs/xcp atau
paket xc lainnya
Cek kode dial
808713# pastikan
tidak ada Xtra Combo
varian apapun
・FUP Belum Bisa Di Temtukan Yang Pasti
・Tanpa menyediakan pulsa
・Estimasi masuk 30 menit - 1 jam
・Membutuhkan OTP

⚠ NO GARANSI
❗PASTIKAN DOR JANGAN DI JAM RAWAN MT
❗MASALAH XUT MASUK TIDAKNYA URUSAN XL YA, YANG PASTI XTRA COMBO SPESIAL & XTRA COMBO PASTI MASUKNYA

🔥 XL VIDIO🔥 ( Rp 33.000 )
🔥 XL VIDIO + CONFIG 🔥 ( Rp 40.000 )
・Pengganti Paket Vidio Yang Hilang
・tanpa gandeng
Total FUP: 100GB-150GB

🔥 XL IFLIX🔥 ( Rp 33.000 )
🔥 XL IFLIX + CONFIG 🔥 ( Rp 40.000 )
・Pengganti Paket iFlix Yang Hilang
・tanpa gandeng
Total FUP: 100GB-150GB
└──────────┘
ORDER ⬇
WA : wa.me/6281327393959
TELE : t.me/ARI_VPN_STORE
└──────────────────┘ 
WEB OTP PAKET ARISTORE
link : https://bit.ly/ARISTORE-OTP
Tutorial :
1. MASUKAN NOMER Xl
2. Klik Minta Kode OTP
3. ISI KODE OTP nya
4. Klik Login
5. Secrenshoot kirim ke admin
Note : jangan lupa unreng paket terlebih dahulu
┌──────────────────┐
 JALUR PEMBELIAN 
TF => UNREG => LOGIN OTP => PROSES => DONE
└──────────────────┘ 
┌──────────────────┐
### AKUN PREMIUM
┌─────────────────┐
        EMAIL BUYYER
YT PREM INVITE
   Harga Rp3.000

YT PREM PRIVATE
   Harga Rp5.000

YT PREM FAMILY
   Harga Rp7.000
  ───────────────────

  ───────────────────
     EMAIL DARI SAYA
YT PREM PRIVATE
   Harga Rp7.000

 YT PREM FAMILY
   Harga Rp10.000
  ───────────────────

Fitur YouTube Premium:
- Bebas iklan di YouTube, YouTube Music, dan YouTube Kids
- Pemutaran di background
- Download video dan musik untuk offline
- Akses ke YouTube Music Premium
- Tersedia di aplikasi YouTube, YouTube Music, dan YouTube Kids

NOTE:
EMAIL WAJIB FRESS (BARU)
YT PRIVATE 1 akun
YT FAMILY  Max invite 5 akun
Masa aktif 1 bulan
FULL GARANSI
└─────────────────┘
┌─────────────────┐
🌟 CANVA PRO & TEAM 🌟

📦 Paket Pilihan:
- Canva Invite 1 Bulan: Rp3.000
- Canva Pro 1 Bulan: Rp5.000
- Canva Team Premium 1 Bulan: Rp10.000

✨ Keuntungan:
✅ Full akses admin (kecuali untuk akun invite).
✅ Canva Team dapat mengundang teman atau tim.
✅ Cocok untuk reseller, jual kembali dan raih keuntungan besar!

🚀 Jangan sampai ketinggalan! Stok terbatas!

🎯 Solusi mudah untuk pelajar, desainer, pebisnis, dan siapa saja yang ingin berkreasi tanpa batas!

📩 Pesan sekarang juga!
└─────────────────┘

┌──────────────────┐
### PANEL-DIGITALOCEAN
  ───────────────────
Redy Panel DigitalOcean
  ───────────────────
💻 Droplet   : 10
💳 Billing   : PayPal
💰 Saldo     : $205
⏳ Expired   : 2 Bulan
💵 Harga     : Rp180.000
⭐ Quality   : VERY GOOD
🔑 Login Via : Web & APIKey
🥇 Garansi   : Saat Pertama kali Login
  ───────────────────
Keuntungan & Kegunaannya  :
- Bisa buat VPS Maksimal 10 Vps
- Bisa buat VPS dari ram 1-16 GB
- Bisa Claim Promocode $200 + $5 Dari Paypal
- Panel Insyaallah awet tahan banting
- Dapet cara agar Panel selalu awet
└──────────────────┘
┌──────────────────┐
   ### LIST-SCRIPT
  ───────────────────
1️⃣ SCRIPT BY ARI TUNNEL V4
   ───────────────────
💲 HARGA SCRIPT 
- 1 bulan / 1 IP : Rp 10.000
- 2 bulan / 2 IP : Rp 20.000
   ───────────────────
RESSEL SCRIPT (NEGO PM)
- Lifetime / Unlimited IP : Rp 100.000
   ───────────────────
🖥 OS SYSTEM SUPORT :
- DEBIAN 10 
- UBUNTU 20.04
  ───────────────────
📡 MENU PROTOCOL :
- SSH WS & SSL
- SSH UDP
- SSH SELOWDNS
- SSH OPENVPN
- VMESS
- VLESS
- TROJAN
- SADOWSOCK
- NOOBVPNS
  ───────────────────
📌 FITUR SCRIPT :
- Add AutoReboot 00.00 (bisa diubah) 
- Add AutoUpdate 01.15
- Add AutoBackup 11.15 (bisa diubah) 
- Cek Running Service
- Restart Service
- Auto Reboot
- Monitor VPS
- Speedtest
- Delete All Exp
- Change Domain
- Change Banner
- Fix Error Domain
- Fix Error Proxy
- Menu Cleaner
- Limit Speed
- Switch On & Off Limit 
- Monitor Account
- Create, Delete, Renew, Trial Account
- Lock & Unlock Account
- List Account
- Limit IP & Quota
- AutoKill Account
- Detail Account
- Cek Login UDP
- Recovery Account
- Edit Limit IP & Quota
- Backup & Restore
- AutoBackup
- Bot Tele Create Account
- Bot Tele Notif
- Multi Path (Support OPOK ISAT)
- Total Fitur: 113±
- DLL
   ───────────────────
🌟 BENEFIT :
- SC Ringan
- SC HAProxy
- SC MultiPort
┌──────────────────┐
2️⃣ SCRIPT POTATO TUNNELING
   ───────────────────
💲 HARGA :
- 1 IP/Bulan : Rp 15.000
- 1 IP/Tahun : Rp 75.000
- SC + BOT V1
- 1 IP/Bulan : Rp 19.000
- 1 IP/Tahun : Rp 105.000
- SC + BOT V2
- 1 IP/Bulan : Rp 20.000
- 1 IP/Tahun : Rp 130.000
   ───────────────────
🖥 OS SYSTEM SUPORT :
- Debian 10 (Recommended)
- Ubuntu 20.04
  ───────────────────
📌 FITUR :
1. Notif User Login Bot Telegram
2. Bisa Atur Limit IP per User 
3. Set Time Banned & Auto Kill
4. Support Semua New Metode
5. Multi Port tidak hanya 443/80
6. Bisa Atur Bandwidth per User
7. Support Custom Multi Path
8. Support Tsel OPOK yang Viral & Orbit OPOK
9. Speedtest
10. RAM Usage
11. Cek Bandwidth
12. Change Timezone
13. Change Core
⚠ Note ⚠
Usahakan VPS Support untuk Tunneling. 
Di luar itu, risiko ditanggung sendiri.
No Refunds No Debat
└──────────────────┘

┌──────────────────┐
   ### PANEL-BOT
PTERODACTYL  By ARI STORE
Per Bulan
RAM 1GB Harga: Rp 1.000
RAM 2GB Harga: Rp 2.000
RAM 3GB Harga: Rp 3.000
RAM 4GB Harga: Rp 4.000
RAM 5GB Harga: Rp 5.000
RAM 6GB Harga: Rp 6.000
RAM 7GB Harga: Rp 7.000
RAM 8GB Harga: Rp 8.000
RAM 9GB Harga: Rp 9.000
RAM 10GB Harga: Rp 10.000
RAM UNLIMITED Harga: Rp 11.000

✨ᴘᴇʀᴘᴀɴᴊᴀɴɢ ɴᴏ ʀɪʙᴇᴛ 
✨ᴅɪ ᴊᴀᴍɪɴ 𝟷𝟶𝟶% ᴍᴜʀᴀʜ 
✨ɴᴏ ᴛɪᴘᴜ + ᴛᴇsᴛɪ ʙᴀɴʏᴀᴋ
✨ʙᴏᴛ ғᴀsᴛ ʀᴇsᴘᴏɴ
✨ᴋᴜᴏᴛᴀ + ᴍᴇᴍᴏʀɪ ᴛɪᴅᴀᴋ ᴄᴇᴘᴀᴛ ʙᴏʀos

Kegunaan Panel BOT:
- Bot Online 24 jam
- Hemat Kuota
- Bot Anti Delay
- Hemat Penyimpanan
- Gak Ribet Buat Run Bot
- Bot Fast Respon
- Masih banyak lagi
└──────────────────┘
ORDER ⬇
WA : wa.me/6281327393959
TELE : t.me/ARI_VPN_STORE
`

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

	case "jpmvpn":
		jid := evt.Info.Chat
		sender := evt.Info.Sender.User

		if sender != ownerJID {
			_, _ = client.SendMessage(context.Background(), jid, &waProto.Message{
				Conversation: proto.String("❌ Perintah ini hanya untuk owner."),
			})
			return
		}

		allGroups, err := client.GetJoinedGroups()
		if err != nil {
			logger.Error("Gagal mengambil grup: " + err.Error())
			return
		}

		totalSent := 0
		messageText := `🔰👑𝗣𝗧 𝗥𝗔𝗝𝗔 𝗦𝗘𝗩𝗘𝗥 𝗣𝗥𝗘𝗠𝗜𝗨𝗠👑🔰
┌──────────────────┐
🥇 𝙱𝚈 𝙰𝚁𝙸 𝚂𝚃𝙾𝚁𝙴 🥇
└──────────────────┘
┌──────────────────┐
🔥 Semua transaksi otomatis via Bot Telegram
✅ Gak perlu ngantri
✅ 24 Jam non-stop
✅ Langsung prose
💻 TELE BOT VPN : t.me/ARI_VPN_STORE_bot
💻 TELE BOT SC TUNNEL : t.me/ARISCTUNNEL_bot
└──────────────────┘
┌──────────────────┐
    # LIST-VPN
PILIH MODE INJECT NYA
1️⃣ SSH WEBSOCKET Account
2️⃣ SSH SelowDNS Account
3️⃣ OPEN SSH Account 
4️⃣ SSH UDP Account
5️⃣ Xray/Vmess Account
6️⃣ Xray/Vless Account
7️⃣ Trojan Account
  ─────────────────── 
🇸🇬 SERVER SINGAPURA BUKA ✓
📝 LIST HARGA & SERVER :
ISP : DigitalOcean, LLC
CPU : DO-Premium-AMD (4) @ 1.996GHz
RAM : 8 GB / vCPUs : 4 / SSD : 320 GB Disk
Bandwidth : 300GB/BLN
🔹 🇸🇬SG VIP [MAX 2 DEVICE] :
1. Rp 1.750 (3 HARI)
2. Rp 2.000 (7 HARI)
3. Rp 3.500 (15 HARI)
4. Rp 5.000 (22 HARI)
5. Rp 7.000 (30 HARI)
6. Rp 15.000 (60 HARI)
   ───────────────────
ISP : DigitalOcean, LLC 🔥
CPU : DO-Premium-AMD (4) @ 1.996GHz
RAM : 8 GB / vCPUs : 4 / SSD : 320 GB Disk
Bandwidth : 500GB/BLN
[OPEN 3 SERVER]
🔹 🇸🇬SG VVIP [SUPPORT STB] :
1. Rp 2.000 (3 HARI)
2. Rp 3.500 (7 HARI)
3. Rp 5.000 (15 HARI)
4. Rp 7.000 (22 HARI)
5. Rp 10.000 (30 HARI)
6. Rp 20.000 (60 HARI)
   ───────────────────
ISP : DigitalOcean, LLC 
CPU : DO-Premium-AMD (4) @ 1.996GHz
RAM : 8 GB / vCPUs : 4 / SSD : 320 GB Disk
Bandwidth : 200GB/BLN
[OPEN 2 SERVER]
🔹 🇸🇬SG RESSEL [MAX 1 DEVICE] :
1. Rp 5.000 (30 HARI)
2. Rp 10.000 (60 HARI)
   ───────────────────
🇮🇩 SERVER INDONESIA BUKA ✓
📝 LIST HARGA & SERVER :
ISP : Media Antar Nusa PT.
CPU : Intel (Broadwell, IBRS) (1) @ 2.299GHz 
RAM : 2 GB / CPU : 1 / 40 GB SSD Disk
Bandwidth : Unlimited
[OPEN 2 SERVER]
🔹 🇮🇩ID NUSA [MAX 2 DEVICE] :
1. Rp 3.500 (3 HARI)
2. Rp 5.000 (7 HARI)
3. Rp 7.000 (15 HARI)
4. Rp 10.000 (22 HARI)
5. Rp 12.000 (30 HARI)
6. Rp 24.000 (60 HARI)
   ───────────────────
ISP : Media Antar Nusa PT. 
CPU : KVM RHEL 7.6.0 PC (1) @ 2.699GHz   
RAM : 2 GB / 1 CPU / 20 GB SSD Disk
Bandwidth : Unlimited
🔹 🇮🇩ID NUSA [SUPPORT STB] :
1. Rp 5.000 (3 HARI)
2. Rp 7.000 (7 HARI)
3. Rp 10.000 (15 HARI)
4. Rp 12.000 (22 HARI)
5. Rp 14.000 (30 HARI)
6. Rp 25.000 (60 HARI)
   ───────────────────
• MINTA CONFIG MENTAHAN Rp 1.000
• CONFIG PREMIUM DARI SAYA GRATIS
✅ BISA TRIAL DULU
   ───────────────────
KELEBIHAN BELI AKUN PREMIUM
🔹 SUPPORT STB OPENWRT ✓
🔹 Nonton YouTube Lancar 4K ✓
🔹 Server Tidak Gampang Error ✓
🔹 Buat Download Fast Speed ✓
🔹 ANTI RECONNECT ✓
🔹 Support Video Call (GANTI UDP 7100-7900) ✓
🔹 Support GAME (TERGANTUNG TKP) ✓
🔹 Koneksi Stabil ✓
🔹 Fast Connect ✓
🔹 Garansi 100% ✓
🔹 Bonus Config ✓
🔹 DLL
   ───────────────────
⚙ PENGATURAN PEMBELI ⚙
🚫 NO DDOS !!!
🚫 NO HACKING !!!
🚫 NO CARDING !!!
🚫 NO TORRENT !!!
🚫 NO SPAMMING !!!
🚫 NO PLAYING PLAYSTATION !!!
🚫 MAX 2 Device !!!
 ❗ Melanggar AUTO BAN ❗
└──────────────────┘
ORDER ⬇
WA : wa.me/6281327393959
TELE : t.me/ARI_VPN_STORE
`

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

	case "jpmvps":
		jid := evt.Info.Chat
		sender := evt.Info.Sender.User

		if sender != ownerJID {
			_, _ = client.SendMessage(context.Background(), jid, &waProto.Message{
				Conversation: proto.String("❌ Perintah ini hanya untuk owner."),
			})
			return
		}

		allGroups, err := client.GetJoinedGroups()
		if err != nil {
			logger.Error("Gagal mengambil grup: " + err.Error())
			return
		}

		totalSent := 0
		messageText := `🔰👑𝗣𝗧 𝗥𝗔𝗝𝗔 𝗦𝗘𝗩𝗘𝗥 𝗣𝗥𝗘𝗠𝗜𝗨𝗠👑🔰
┌──────────────────┐
🥇 𝙱𝚈 𝙰𝚁𝙸 𝚂𝚃𝙾𝚁𝙴 🥇
└──────────────────┘
┌──────────────────┐
   ### LIST-VPS
  ───────────────────
★ LIST VPS DIGITALOCEAN ★

READY VPS | DISKON 20%

2️⃣ 1CPU / 1GB / 25GB SSD / 1TB BW 
• Tanpa SC : 20K 
• SC Arisctunnel v4 : 25K 
• SC Potato : 35K

2️⃣ 1CPU / 2GB / 50GB SSD / 2TB BW 
• Tanpa SC : 35K 
• SC Arisctunnel v4 : 40K 
• SC Potato : 50K

3️⃣ 2CPU / 2GB / 50GB SSD / 3TB BW 
• Tanpa SC : 45K 
• SC Arisctunnel v4 : 50K 
• SC Potato : 60K

4️⃣ 2CPU / 4GB / 80GB SSD / 4TB BW 
• Tanpa SC : 65K 
• SC Arisctunnel v4 : 70K 
• SC Potato : 80K

5️⃣ 4CPU / 8GB / 160GB SSD / 5TB BW 
• Tanpa SC : 85K 
• SC Arisctunnel v4 : 90K 
• SC Potato : 100K 
• SC RDP : 100K 
• SC Pterodactyl : 100K

INFORMASI GARANSI 
- RAM 1 NO GARANSI
- RAM 2-8 GARANSI 1x Replace

 TOS / SYARAT KETENTUAN : 
- Dilarang CPU 100%
- Dilarang DDoS
- Dilarang Torrent
- Dilarang Hacking
- Dilarang Mining
- Dilarang Power Off
- Dilarang Konten Ilegal

INFO:
- Free Domain Premium 
- Free Rebulid (Instal ulang)
- Full SC, Tinggal Pakai 
- Bisa Jualan Panel Bot 
- Bisa Jualan RDP & XRDP 
- Bisa Jualan RDP Forex
- Bisa Buat Website Hosting
- Jualan SSH, XRAY, Trojan 
- Bisa Ganti Script lain 

MINAT?
• Profil: [bit.ly/m/ADMIN-ARISTORE] 
• Testi: [t.me/RAJA_VPN_STORE] 
• Amanah & Terpercaya
└──────────────────┘
`

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

	case "jpmdor":
		jid := evt.Info.Chat
		sender := evt.Info.Sender.User

		if sender != ownerJID {
			_, _ = client.SendMessage(context.Background(), jid, &waProto.Message{
				Conversation: proto.String("❌ Perintah ini hanya untuk owner."),
			})
			return
		}

		allGroups, err := client.GetJoinedGroups()
		if err != nil {
			logger.Error("Gagal mengambil grup: " + err.Error())
			return
		}

		totalSent := 0
		messageText := `🔰👑𝗣𝗧 𝗥𝗔𝗝𝗔 𝗦𝗘𝗩𝗘𝗥 𝗣𝗥𝗘𝗠𝗜𝗨𝗠👑🔰
┌──────────────────┐
🥇 𝙱𝚈 𝙰𝚁𝙸 𝚂𝚃𝙾𝚁𝙴 🥇
└──────────────────┘
PROMO PAKET INJECT
🔥 XL VIDIO🔥 ( Rp 33.000 )
🔥 XL VIDIO + CONFIG 🔥 ( Rp 40.000 )
・Pengganti Paket Vidio Yang Hilang
・tanpa gandeng
Total FUP: 100GB-150GB

🔥 XL IFLIX🔥 ( Rp 33.000 )
🔥 XL IFLIX + CONFIG 🔥 ( Rp 40.000 )
・Pengganti Paket iFlix Yang Hilang
・tanpa gandeng
Total FUP: 100GB-150GB

🔥XC 1+1GB + XUTSuper ( Rp. 23.000 )
🔥XC 1+1GB + XUTSuper + CONFIG ( Rp. 30.000 )
・ Unlimited Whatsapp
・ Line
・ Gojek
・ Instagram
・ Facebook
Total FUP: 230GB-400GB

🔥XCS + XUTBasic ( Rp. 34.000 )
🔥XCS + XUTBasic + CONFIG ( Rp. 41.000 )
・ Unlimited Whatsapp
・ Line
・ Gojek
Total FUP: 230GB-400GB

🔥XCS + XUTStandard ( Rp. 36.000 )
🔥XCS + XUTStandard + CONFIG ( Rp. 43.000 )
・ Unlimited Whatsapp
・ Line
・ Gojek
・ Facebook
Total FUP: 230GB-400GB

🔥XCS + XUTPremiun ( Rp. 38.000 )
🔥XCS + XUTPremiun + CONFIG ( Rp. 45.000 )
・ Unlimited Whatsapp
・ Line
・ Gojek
・ Instagram
・ Facebook
・ Youtube
Total FUP: 230GB-400GB

🔥 XCS + Multi Add On ( Rp 40.000 )
🔥 XCS + Multi Add On + CONFIG ( Rp 47.000 )
・ XUTBasic
・ XUTStandard
・ XUTSuper
・ XUTPremium
・ XUTYouTube
・ XUTJoox
・ XUTViu
・ XUTNetflix
・ XUTTikTok
Total FUP: 1TB++

Notes :
・Kartu tidak boleh ada
paket xcs/xcp atau
paket xc lainnya
Cek kode dial
*808*7*1*3# pastikan
tidak ada Xtra Combo
varian apapun
・FUP Belum Bisa Di Temtukan Yang Pasti
・Tanpa menyediakan pulsa
・Estimasi masuk 30 menit - 1 jam
・Membutuhkan OTP

⚠ NO GARANSI

ORDER ⬇
WA : wa.me/6281327393959
TELE : t.me/ARI_VPN_STORE
`

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
