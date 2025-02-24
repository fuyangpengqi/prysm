package storage

import (
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/v5/beacon-chain/db/filesystem"
	"github.com/prysmaticlabs/prysm/v5/beacon-chain/node"
	"github.com/prysmaticlabs/prysm/v5/cmd"
	"github.com/prysmaticlabs/prysm/v5/config/params"
	"github.com/prysmaticlabs/prysm/v5/consensus-types/primitives"
	"github.com/urfave/cli/v2"
)

var (
	// BlobStoragePathFlag defines a flag to start the beacon chain from a give genesis state file.
	BlobStoragePathFlag = &cli.PathFlag{
		Name:  "blob-path",
		Usage: "Location for blob storage. Default location will be a 'blobs' directory next to the beacon db.",
	}
	BlobRetentionEpochFlag = &cli.Uint64Flag{
		Name:    "blob-retention-epochs",
		Usage:   "Override the default blob retention period (measured in epochs). The node will exit with an error at startup if the value is less than the default of 4096 epochs.",
		Value:   uint64(params.BeaconConfig().MinEpochsForBlobsSidecarsRequest),
		Aliases: []string{"extend-blob-retention-epoch"},
	}
	BlobStorageLayout = &cli.StringFlag{
		Name:  "blob-storage-layout",
		Usage: layoutFlagUsage(),
		Value: filesystem.LayoutNameFlat,
	}
)

func layoutOptions() string {
	return "available options are: " + strings.Join(filesystem.LayoutNames, ", ") + "."
}

func layoutFlagUsage() string {
	return "Dictates how to organize the blob directory structure on disk, " + layoutOptions()
}

func validateLayoutFlag(_ *cli.Context, v string) error {
	for _, l := range filesystem.LayoutNames {
		if v == l {
			return nil
		}
	}
	return errors.Errorf("invalid value '%s' for flag --%s, %s", v, BlobStorageLayout.Name, layoutOptions())
}

// BeaconNodeOptions sets configuration values on the node.BeaconNode value at node startup.
// Note: we can't get the right context from cli.Context, because the beacon node setup code uses this context to
// create a cancellable context. If we switch to using App.RunContext, we can set up this cancellation in the cmd
// package instead, and allow the functional options to tap into context cancellation.
func BeaconNodeOptions(c *cli.Context) ([]node.Option, error) {
	e, err := blobRetentionEpoch(c)
	if err != nil {
		return nil, err
	}
	opts := []node.Option{node.WithBlobStorageOptions(
		filesystem.WithBlobRetentionEpochs(e),
		filesystem.WithBasePath(blobStoragePath(c)),
		filesystem.WithLayout(c.String(BlobStorageLayout.Name)), // This is validated in the Action func for BlobStorageLayout.
	)}
	return opts, nil
}

func blobStoragePath(c *cli.Context) string {
	blobsPath := c.Path(BlobStoragePathFlag.Name)
	if blobsPath == "" {
		// append a "blobs" subdir to the end of the data dir path
		blobsPath = path.Join(c.String(cmd.DataDirFlag.Name), "blobs")
	}
	return blobsPath
}

var errInvalidBlobRetentionEpochs = errors.New("value is smaller than spec minimum")

// blobRetentionEpoch returns the spec default MIN_EPOCHS_FOR_BLOB_SIDECARS_REQUEST
// or a user-specified flag overriding this value. If a user-specified override is
// smaller than the spec default, an error will be returned.
func blobRetentionEpoch(cliCtx *cli.Context) (primitives.Epoch, error) {
	spec := params.BeaconConfig().MinEpochsForBlobsSidecarsRequest
	if !cliCtx.IsSet(BlobRetentionEpochFlag.Name) {
		return spec, nil
	}

	re := primitives.Epoch(cliCtx.Uint64(BlobRetentionEpochFlag.Name))
	// Validate the epoch value against the spec default.
	if re < params.BeaconConfig().MinEpochsForBlobsSidecarsRequest {
		return spec, errors.Wrapf(errInvalidBlobRetentionEpochs, "%s=%d, spec=%d", BlobRetentionEpochFlag.Name, re, spec)
	}

	return re, nil
}

func init() {
	BlobStorageLayout.Action = validateLayoutFlag
}
