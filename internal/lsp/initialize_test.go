package lsp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"terragrunt-ls/internal/lsp"
)

func TestNewInitializeResponse_AdvertisesDocumentFormatting(t *testing.T) {
	t.Parallel()

	response := lsp.NewInitializeResponse(1)

	provider, ok := response.Result.Capabilities.DocumentFormattingProvider.(bool)
	assert.True(t, ok)
	assert.True(t, provider)
}
