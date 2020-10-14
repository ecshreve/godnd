package client_test

import (
	"context"
	"testing"

	"github.com/ecshreve/godnd/pkg/client"
	"github.com/samsarahq/go/snapshotter"
	"github.com/stretchr/testify/assert"
)

func TestGetClasses(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()
	ctx := context.Background()

	c := client.NewClient()
	res, err := c.GetClasses(ctx)
	assert.NoError(t, err)
	snap.Snapshot("GetClasses", res)
}
