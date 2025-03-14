
package discord

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

type DiscordAdapter struct {
    webhookURLs map[string]string // Mapa de canales y URLs de webhook
}

func NewDiscordAdapter(webhookURLs map[string]string) *DiscordAdapter {
    return &DiscordAdapter{
        webhookURLs: webhookURLs,
    }
}

func (d *DiscordAdapter) Send(channel, message string) error {
    webhookURL, ok := d.webhookURLs[channel]
    if !ok {
        return fmt.Errorf("canal no encontrado: %s", channel)
    }

    payload := map[string]string{
        "content": message,
    }

    jsonData, err := json.Marshal(payload)
    if err != nil {
        return fmt.Errorf("error al serializar el payload: %v", err)
    }

    resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return fmt.Errorf("error al enviar la notificación a Discord: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusNoContent {
        return fmt.Errorf("respuesta inesperada de Discord: %s", resp.Status)
    }

    return nil
}