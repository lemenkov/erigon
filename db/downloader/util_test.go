// Copyright 2026 The Erigon Authors
// This file is part of Erigon.
//
// Erigon is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Erigon is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Erigon. If not, see <http://www.gnu.org/licenses/>.

package downloader

import (
	"testing"
	"time"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/stretchr/testify/require"
)

func TestCreateMetaInfoCreationDate(t *testing.T) {
	info := &metainfo.Info{PieceLength: 1024, Name: "test", Length: 0}

	t.Run("fixed value is reproducible", func(t *testing.T) {
		mi, err := CreateMetaInfo(info, nil, ErigonGenesisTimestamp)
		require.NoError(t, err)
		require.Equal(t, ErigonGenesisTimestamp, mi.CreationDate)
	})

	t.Run("override value", func(t *testing.T) {
		const custom int64 = 1700000000
		mi, err := CreateMetaInfo(info, nil, custom)
		require.NoError(t, err)
		require.Equal(t, custom, mi.CreationDate)
	})

	t.Run("non-positive falls back to wall clock", func(t *testing.T) {
		before := time.Now().Unix()
		mi, err := CreateMetaInfo(info, nil, 0)
		require.NoError(t, err)
		require.GreaterOrEqual(t, mi.CreationDate, before)
	})
}
