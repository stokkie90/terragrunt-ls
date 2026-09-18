package tg

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormatDocumentUsesTerragruntCLI(t *testing.T) {
	original := runTerragruntHclFmt
	t.Cleanup(func() {
		runTerragruntHclFmt = original
	})

	runTerragruntHclFmt = func(ctx context.Context, filename, document string) ([]byte, error) {
		assert.NotNil(t, ctx)
		assert.Equal(t, "/tmp/terragrunt.hcl", filename)
		assert.Equal(t, "locals{}", document)
		return []byte("formatted-by-terragrunt"), nil
	}

	formatted, err := formatDocument(context.Background(), "/tmp/terragrunt.hcl", "locals{}")
	require.NoError(t, err)
	assert.Equal(t, "formatted-by-terragrunt", string(formatted))
}

func TestFormatDocumentFallsBackWhenTerragruntUnavailable(t *testing.T) {
	original := runTerragruntHclFmt
	t.Cleanup(func() {
		runTerragruntHclFmt = original
	})

	runTerragruntHclFmt = func(context.Context, string, string) ([]byte, error) {
		return nil, errors.New("terragrunt not found")
	}

	formatted, err := formatDocument(context.Background(), "/tmp/terragrunt.hcl", "locals{\nfoo=\"bar\"\n}")
	require.Error(t, err)
	assert.Equal(t, "terragrunt not found", err.Error())
	assert.Equal(t, "locals {\n  foo = \"bar\"\n}", string(formatted))
}
