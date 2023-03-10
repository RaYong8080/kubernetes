/*
Copyright 2017 The Kubernetes Authors.

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

package benchmark

import (
	"flag"
	"testing"

	"k8s.io/klog/v2/ktesting"
	"k8s.io/kubernetes/test/integration/framework"
)

func TestMain(m *testing.M) {
	// Run with -v=2, this is the default log level in production.
	ktesting.DefaultConfig = ktesting.NewConfig(ktesting.Verbosity(2))
	ktesting.DefaultConfig.AddFlags(flag.CommandLine)
	flag.Parse()

	framework.EtcdMain(m.Run)
}
