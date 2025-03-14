package service

import (
    "project-notification/internal/domain"
   "project-notification/internal/interfaces/ports"
)

type EventService struct {
    notifier ports.DiscordNotifier
}

func NewEventService(notifier ports.DiscordNotifier) *EventService {
    return &EventService{
        notifier: notifier,
    }
}

func (s *EventService) ProcessEvent(event domain.GitHubEvent) error {
    var channel string
    var message string

    // Determina el canal y el mensaje según el tipo de evento
    switch event.Type {
    case "pull_request":
        channel = "desarrollo"
        message = "PR " + event.Action + " en " + event.Repo + " por " + event.User + " (Rama: " + event.Branch + "): " + event.PRURL
    case "push":
        channel = "pruebas"
        message = "Push en " + event.Repo + " por " + event.User + " (Rama: " + event.Branch + ", Commit: " + event.CommitID + ")"
    case "tests":
        channel = "pruebas"
        message = "Pruebas exitosas en " + event.Repo + " por " + event.User + " (Rama: " + event.Branch + ")"
    default:
        channel = "general"
        message = "Evento en " + event.Repo + ": " + event.Type
    }

    // Envía la notificación a Discord
    return s.notifier.Send(channel, message)
}