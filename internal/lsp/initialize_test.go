package lsp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"terragrunt-ls/internal/lsp"
)

func TestNewInitializeResponse_AdvertisesDocumentFormatting(t *testing.T) {
	t.Parallel()

	response := lsp.NewInitializeResponse(1)

	assert.True(t, response.Result.Capabilities.DocumentFormattingProvider)
}
