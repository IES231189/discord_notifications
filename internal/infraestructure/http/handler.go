package http

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "project-notification/internal/Application/services"
    "project-notification/internal/domain"
)

type GinHandler struct {
    eventService *service.EventService
}

func NewGinHandler(eventService *service.EventService) *GinHandler {
    return &GinHandler{
        eventService: eventService,
    }
}

func (h *GinHandler) HandleGitHubEvent(c *gin.Context) {
    var event domain.GitHubEvent
    if err := c.ShouldBindJSON(&event); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.eventService.ProcessEvent(event); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "event processed"})
}