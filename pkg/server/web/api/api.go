// Copyright 2019 Squeeze Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"net/http"
	"github.com/agile6v/squeeze/pkg/util"
)

type AppAPI struct {

}

func (api *AppAPI) Init() {
	http.HandleFunc("/", api.Index)
	http.HandleFunc("/api/create", api.create)
	http.HandleFunc("/api/delete", api.delete)
	http.HandleFunc("/api/search", api.search)
	http.HandleFunc("/api/list", api.list)
	http.HandleFunc("/api/start", api.start)
	http.HandleFunc("/api/stop", api.stop)
}

func (api *AppAPI) Index(w http.ResponseWriter, request *http.Request) {
	util.RespondWithJSON(w, http.StatusOK, "")
}

func (api *AppAPI) create(writer http.ResponseWriter, request *http.Request) {

}

func (api *AppAPI) delete(writer http.ResponseWriter, request *http.Request) {

}

func (api *AppAPI) search(writer http.ResponseWriter, request *http.Request) {

}

func (api *AppAPI) list(writer http.ResponseWriter, request *http.Request) {

}

func (api *AppAPI) start(writer http.ResponseWriter, request *http.Request) {

}

func (api *AppAPI) stop(writer http.ResponseWriter, request *http.Request) {

}