package presence

import (
	"context"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"pycrs.cz/what-it-doo/internal/bus"
	"pycrs.cz/what-it-doo/internal/bus/payload"
	"pycrs.cz/what-it-doo/internal/domain/service"
)

type PresenceManager struct {
	mu          sync.RWMutex
	r           *redis.Client
	gatewayID   string
	userService service.UserService
	bus         bus.CommnunicationBus

	local map[uuid.UUID]int
}

func NewPresenceManager(r *redis.Client, gatewayID string, userService service.UserService, bus bus.CommnunicationBus) *PresenceManager {
	return &PresenceManager{
		local:       make(map[uuid.UUID]int),
		r:           r,
		gatewayID:   gatewayID,
		userService: userService,
		bus:         bus,
	}
}

func (pm *PresenceManager) AddConnection(ctx context.Context, userID uuid.UUID) error {
	// if first connection for user, add gateway id to redis
	// otherwise just increment
	pm.mu.Lock()
	defer pm.mu.Unlock()

	pm.local[userID]++

	if pm.local[userID] == 1 {
		// check redis if user online gateways set is empty, if so set online status in db to true and add gateway id
		cnt, err := pm.r.SCard(ctx, presenceKey(userID)).Result()
		if err == nil && cnt == 0 {
			return pm.setPresence(ctx, userID, true)
		}

		// first connection for user, add gateway id to redis
		pm.r.SAdd(ctx, presenceKey(userID), pm.gatewayID)
	}
	return nil
}

func (pm *PresenceManager) RemoveConnection(ctx context.Context, userID uuid.UUID) error {
	// if last connection for user, remove gateway id from redis
	// otherwise just decrement
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if count, ok := pm.local[userID]; ok {
		if count <= 1 {
			// last connection for user, remove gateway id from redis
			delete(pm.local, userID)
			pm.r.SRem(ctx, presenceKey(userID), pm.gatewayID)

			// check redis if user online gateways set is empty, if so set online status in db to false
			cnt, err := pm.r.SCard(ctx, presenceKey(userID)).Result()
			if err == nil && cnt == 0 {
				return pm.setPresence(ctx, userID, false)
			}
		} else {
			pm.local[userID]--
		}
	}
	return nil
}

func presenceKey(userID uuid.UUID) string {
	return "presence:" + userID.String()
}

func (pm *PresenceManager) setPresence(ctx context.Context, userID uuid.UUID, online bool) error {
	log.Printf("Chaninging presence for user %s to %v", userID, online)
	if err := pm.userService.SetPresence(ctx, userID, online); err != nil {
		return err
	}

	return pm.bus.DispatchGlobalEvent(ctx, bus.PresenceChangeEventType, payload.PresenceChangeEventPayload{
		UserID: userID,
		Online: online,
	})
}
