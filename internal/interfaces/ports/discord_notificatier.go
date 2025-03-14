
package ports

type DiscordNotifier interface {
    Send(channel, message string) error
}