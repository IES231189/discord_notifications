
package main

import (
    "github.com/gin-gonic/gin"
    "project-notification/internal/Application/services"
    "project-notification/internal/infraestructure/discord"
    "project-notification/internal/infraestructure/http"
    
)

func main() {
    
    webhookURLs := map[string]string{
        "desarrollo": "https://discordapp.com/api/webhooks/1349289090096304181/YeIZrmkXOAQqe1V6-OGi83Of5Q3D1-gDAyR_umnln45XE2VKrajgkpICuyk68hO1Y-kI",
        "pruebas":    "https://discordapp.com/api/webhooks/1349288681114177536/EYCsW6KFvrgGHqL0cSwNnJkdy6Li9jgADRWPMlP8g8-My874NCI2iiEvK3mXIvqUt907",
        "general":    "https://discordapp.com/api/webhooks/1349288135154208798/Wrre3EKCx6hZb04iG2BkV-2aBC5s_XImVNWLLMWhmvtbWDs0JMLBzOMMqN4GxhoBqYAB",
    }

    
    discordAdapter := discord.NewDiscordAdapter(webhookURLs)

    
    eventService := service.NewEventService(discordAdapter)

    
    ginHandler := http.NewGinHandler(eventService)

    
    r := gin.Default()
    r.POST("/webhook", ginHandler.HandleGitHubEvent)

   
    if err := r.Run(":4040"); err != nil {
        panic("No se pudo iniciar el servidor: " + err.Error())
    }
}