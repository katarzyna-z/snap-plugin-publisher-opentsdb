/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"os"

	"github.com/intelsdi-x/pulse-plugin-publisher-opentsdb/opentsdb"
	"github.com/intelsdi-x/pulse/control/plugin"
)

func main() {
	meta := opentsdb.Meta()
	plugin.Start(meta, opentsdb.NewOpentsdbPublisher(), os.Args[1])
}
