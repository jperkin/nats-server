// Copyright 2021 The NATS Authors
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
//
//go:build illumos || solaris
// +build illumos solaris

package sysmem

/*
#include <unistd.h>
*/
import "C"

func Memory() int64 {
	pagesize := int64(C.sysconf(C._SC_PAGE_SIZE))
	pagenum := int64(C.sysconf(C._SC_PHYS_PAGES))
	return pagesize * pagenum
}
