// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Licensed under the Simple Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://opensource.org/licenses/Simple-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hugolib

import (
	"errors"
	"reflect"
	"sort"
)

type PageGroup struct {
	Key  interface{}
	Data Pages
}

type mapKeyValues []reflect.Value

func (v mapKeyValues) Len() int      { return len(v) }
func (v mapKeyValues) Swap(i, j int) { v[i], v[j] = v[j], v[i] }

type mapKeyByInt struct{ mapKeyValues }

func (s mapKeyByInt) Less(i, j int) bool { return s.mapKeyValues[i].Int() < s.mapKeyValues[j].Int() }

type mapKeyByStr struct{ mapKeyValues }

func (s mapKeyByStr) Less(i, j int) bool {
	return s.mapKeyValues[i].String() < s.mapKeyValues[j].String()
}

func sortKeys(v []reflect.Value, order string) []reflect.Value {
	if len(v) <= 1 {
		return v
	}

	switch v[0].Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if order == "desc" {
			sort.Sort(sort.Reverse(mapKeyByInt{v}))
		} else {
			sort.Sort(mapKeyByInt{v})
		}
	case reflect.String:
		if order == "desc" {
			sort.Sort(sort.Reverse(mapKeyByStr{v}))
		} else {
			sort.Sort(mapKeyByStr{v})
		}
	}
	return v
}

func (p Pages) GroupBy(key, order string) ([]PageGroup, error) {
	if len(p) < 1 {
		return nil, nil
	}

	if order != "asc" && order != "desc" {
		return nil, errors.New("order argument must be 'asc' or 'desc'")
	}

	ppt := reflect.TypeOf(&Page{})
	ft, ok := ppt.Elem().FieldByName(key)
	if !ok {
		return nil, errors.New("No such field in Page struct")
	}
	tmp := reflect.MakeMap(reflect.MapOf(ft.Type, reflect.SliceOf(ppt)))

	for _, e := range p {
		ppv := reflect.ValueOf(e)
		fv := ppv.Elem().FieldByName(key)
		if !tmp.MapIndex(fv).IsValid() {
			tmp.SetMapIndex(fv, reflect.MakeSlice(reflect.SliceOf(ppt), 0, 0))
		}
		tmp.SetMapIndex(fv, reflect.Append(tmp.MapIndex(fv), ppv))
	}

	var r []PageGroup
	for _, k := range sortKeys(tmp.MapKeys(), order) {
		r = append(r, PageGroup{Key: k.Interface(), Data: tmp.MapIndex(k).Interface().([]*Page)})
	}

	return r, nil
}

func (p Pages) GroupByDate(format, order string) ([]PageGroup, error) {
	if len(p) < 1 {
		return nil, nil
	}

	if order != "asc" && order != "desc" {
		return nil, errors.New("order argument must be 'asc' or 'desc'")
	}

	sp := p.ByDate()
	if order == "desc" {
		sp = sp.Reverse()
	}

	date := sp[0].Date.Format(format)
	var r []PageGroup
	r = append(r, PageGroup{Key: date, Data: make(Pages, 0)})
	r[0].Data = append(r[0].Data, sp[0])

	i := 0
	for _, e := range sp[1:] {
		date = e.Date.Format(format)
		if r[i].Key.(string) != date {
			r = append(r, PageGroup{Key: date})
			i++
		}
		r[i].Data = append(r[i].Data, e)
	}
	return r, nil
}
