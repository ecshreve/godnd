package client_test

import (
	"context"
	"testing"

	"github.com/ecshreve/godnd/pkg/client"
	"github.com/samsarahq/go/snapshotter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestGetClassByIndex(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()
	ctx := context.Background()

	c := client.NewClient()
	classes, err := c.GetClasses(ctx)
	require.NoError(t, err)

	for _, class := range classes.Resources {
		t.Run(class.Index, func(t *testing.T) {
			res, err := c.GetClassByIndex(ctx, class.Index)
			assert.NoError(t, err)
			snap.Snapshot(class.Index, res)
		})
	}
}
