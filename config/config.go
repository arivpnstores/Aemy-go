// Package config stores global configuration settings for the WhatsApp bot.
// This includes settings like command prefixes and bot owner identifiers,
// allowing for easy customization without modifying core application logic.
package config

// Prefixes defines a list of characters that are recognized as command prefixes.
// Messages starting with any of these characters will be treated as commands.
var Prefixes = []rune{'!', '.'}

// Owners contains a list of WhatsApp user IDs (phone numbers) that have
// administrative privileges. These users can access special commands
// and perform restricted actions.
var Owners = []string{
	"6281327393959",
}
