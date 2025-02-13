package storage

import (
	"context"

	"github.com/mohammadne/takhir/internal/entities"
	"github.com/mohammadne/takhir/pkg/databases/postgres"
	"go.uber.org/zap"
)

type Credential interface {
}

func NewCredential(lg *zap.Logger, postgres *postgres.Postgres) Credential {
	return &credential{logger: lg, postgres: postgres}
}

type credential struct {
	logger   *zap.Logger
	postgres *postgres.Postgres
}

func (c *credential) FindCredentialByMethodAndIdentifier(ctx context.Context,
	method entities.CredentialMethod, identifier entities.CredentialIdentifier) {

	// query := `
	// SELECT ID, DESCRIPTION, RESOURCE_ID, DP_LINK, PWA_LINK, HEIGHT, WIDTH
	// FROM "BANNER"
	// WHERE ID = :BANNER_ID`
}
