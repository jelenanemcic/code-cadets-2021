package controllers

import "github.com/jelenanemcic/code-cadets-2021/homework_4/event_api/internal/api/controllers/models"

// EventUpdateValidator validates event update requests.
type EventUpdateValidator interface {
	EventUpdateIsValid(eventUpdateRequestDto models.EventUpdateRequestDto) bool
}
