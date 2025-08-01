// Copyright 2025 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build illumos || solaris

package server

import (
	"golang.org/x/sys/unix"
	"os"
)

func diskAvailable(storeDir string) int64 {
	var ba int64
	if _, err := os.Stat(storeDir); os.IsNotExist(err) {
		os.MkdirAll(storeDir, defaultDirPerms)
	}
	var fs unix.Statvfs_t
	if err := unix.Statvfs(storeDir, &fs); err == nil {
		// Estimate 75% of available storage.
		ba = int64(uint64(fs.Bavail) * uint64(fs.Frsize) / 4 * 3)
	} else {
		// Used 1TB default as a guess if all else fails.
		ba = JetStreamMaxStoreDefault
	}
	return ba
}
