package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMongoDBConnection(t *testing.T) {

	cfg := DBConfig{}
	_, err := NewMongoHandler(cfg)
	require.NoError(t, err, "connection must be stablished with default params")
}
