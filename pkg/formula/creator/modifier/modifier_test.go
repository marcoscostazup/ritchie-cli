/*
 * Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package modifier

import (
	"reflect"
	"testing"

	"github.com/ZupIT/ritchie-cli/pkg/formula"
)

func TestNewModifiers(t *testing.T) {
	type args struct {
		create formula.Create
	}

	tests := []struct {
		name string
		args args
		in   string
		want string
	}{
		{
			name: "modify with success",
			args: args{
				create: formula.Create{
					FormulaCmd: "rit testing formula",
				},
			},
			in:   `tags: "#rit-replace{formulaTags}" cmd: #rit-replace{formulaCmd}`,
			want: `tags: "testing", "formula" cmd: rit testing formula`,
		},
		{
			name: "not modify",
			args: args{
				create: formula.Create{
					FormulaCmd: "rit testing formula",
				},
			},
			in:   `some test`,
			want: `some test`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewModifiers(tt.args.create)
			got := Modify([]byte(tt.in), m)
			if !reflect.DeepEqual(got, []byte(tt.want)) {
				t.Errorf("\nModify() =\n%v\nwant:\n%v", string(got), tt.want)
			}
		})
	}
}
