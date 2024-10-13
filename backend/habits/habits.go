package habits

import (
	"context"

	"github.com/go-chi/chi/v5"
)

type (
	Dispatcher interface {
		Dispatch(r chi.Router)
	}

	Service struct {
		publicControllers    []Dispatcher
		protectedControllers []Dispatcher
	}
)

func NewService() *Service {
	return &Service{}
}

func (svc *Service) Mount(ctx context.Context, r chi.Router, serverID string) error {
	err := svc.MountStatic(ctx, r, serverID)
	if err != nil {
		return err
	}
	return svc.MountWithoutStatic(ctx, r, serverID)
}

func (svc *Service) MountStatic(ctx context.Context, r chi.Router, servcerID string) error {
	return nil
}

func (svc *Service) MountWithoutStatic(ctx context.Context, r chi.Router, serverID string) error {

	r.Group(func(r chi.Router) {
		// Add middleware here
		for _, v := range svc.protectedControllers {
			v.Dispatch(r)
		}
	})

	for _, v := range svc.publicControllers {
		v.Dispatch(r)
	}

	return nil
}

func (svc *Service) RegisterPublicController(d ...Dispatcher) {
	svc.publicControllers = append(svc.publicControllers, d...)
}

func (svc *Service) RegisterProtectedController(d ...Dispatcher) {
	svc.protectedControllers = append(svc.protectedControllers, d...)
}
