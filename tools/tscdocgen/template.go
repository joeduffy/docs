// Copyright 2016-2018, Pulumi Corporation.
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

package main

import (
	"fmt"

	"github.com/cbroglie/mustache"
	"github.com/gobuffalo/packr"
)

var indexTemplate = parseNodeJSTemplate()

type boxPartialProvider struct {
	box packr.Box
}

func (p boxPartialProvider) Get(name string) (string, error) {
	return p.box.FindString(name + ".tmpl")
}

func parseNodeJSTemplate() *mustache.Template {
	templates := packr.NewBox("./templates")

	tstr, err := templates.FindString("nodejs.tmpl")
	if err != nil {
		panic(fmt.Sprintf("missing templates/nodejs.tmpl template: %v", err))
	}

	partialProvider := boxPartialProvider{box: templates}
	t, err := mustache.ParseStringPartials(tstr, partialProvider)
	if err != nil {
		panic(fmt.Sprintf("error parsing templates/nodejs.tmpl template: %v", err))
	}

	return t
}
