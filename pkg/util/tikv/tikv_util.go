// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tikv

import (
	"github.com/tikv/client-go/v2/config"
	"github.com/tikv/client-go/v2/txnkv"

	"github.com/milvus-io/milvus/pkg/v2/util/paramtable"
)

func GetTiKVClient(cfg *paramtable.TiKVConfig) (*txnkv.Client, error) {
	if cfg.TiKVUseSSL.GetAsBool() {
		f := func(conf *config.Config) {
			conf.Security = config.NewSecurity(cfg.TiKVTLSCACert.GetValue(), cfg.TiKVTLSCert.GetValue(), cfg.TiKVTLSKey.GetValue(), []string{})
		}
		config.UpdateGlobal(f)
		return txnkv.NewClient([]string{cfg.Endpoints.GetValue()})
	}
	return txnkv.NewClient([]string{cfg.Endpoints.GetValue()})
}
