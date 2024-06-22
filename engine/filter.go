package engine

import (
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

// gjson.Type :
// json:5, array:5, int:2, string:3

// match verifies that data matches the conditions
func match(filter gjson.Result, data string) (result bool, err error) {
	// TODO should return syntax error if op unknown

	result = true

	filter.ForEach(func(qk, qv gjson.Result) bool {

		dv := gjson.Get(data, qk.String())

		if qv.Type == 5 { // 5:json, int:2, string:3
			qv.ForEach(func(sqk, sqv gjson.Result) bool {

				if sqv.Type == 3 { // 3:string,
					//fmt.Println("here with: ", sqk.String())

					switch sqk.String() {
					case "$gt":
						if !(dv.String() > sqv.String()) {
							result = false
						}
						return result

					case "$lt":
						if !(dv.String() < sqv.String()) {
							result = false
						}
						return result

					case "$gte":
						if !(dv.String() >= sqv.String()) {
							result = false
						}
						return result

					case "$lte":
						if !(dv.String() <= sqv.String()) {
							result = false
						}
						return result

					case "$eq":
						if dv.String() != sqv.String() {
							result = false
						}
						return result
					case "$ne":
						if dv.String() == sqv.String() {
							result = false
						}
						return result

					case "$or": // not in

						result = false
						return result

					case "$st": // is start with
						if !strings.HasPrefix(dv.String(), sqv.String()) {
							result = false
						}
						return result

					case "$en": // is end with
						if !strings.HasSuffix(dv.String(), sqv.String()) {
							result = false
						}
						return result

					case "$c": // is contains
						if !strings.Contains(dv.String(), sqv.String()) {
							result = false
						}
						return result

					default:
						err = fmt.Errorf("unknown %s operation", sqk.Value())
						//fmt.Println("..wher here", sqk.Value(), sqk.Type)
						result = false
						return result
					}

				}

				switch sqk.String() {
				case "$gt":
					if !(dv.Int() > sqv.Int()) {
						result = false
					}
					return result

				case "$lt":
					if !(dv.Int() < sqv.Int()) {
						result = false
					}
					return result

				case "$gte":
					if !(dv.Int() >= sqv.Int()) {
						result = false
					}
					return result

				case "$lte":
					if !(dv.Int() <= sqv.Int()) {
						result = false
					}
					return result

				case "$eq":
					if dv.Int() != sqv.Int() {
						result = false
					}
					return result

				case "$ne":
					if dv.Int() == sqv.Int() {
						result = false
					}
					return result
				case "$in": // in array
					for _, v := range sqv.Array() {
						if dv.String() == v.String() {
							return result
						}
					}
					result = false
					return result

				case "$nin": // not in
					for _, v := range sqv.Array() {
						if dv.String() == v.String() {
							result = false
							return result
						}
					}
					return result

				default:
					// {$and:[{name:{$eq:"adam"}},{name:{$eq:"jawad"}}]}
					// {$or: [{name:{$eq:"adam"}}, {name:{$eq:"jawad"}}]}

					if qk.Str == "$and" {

						for _, v := range qv.Array() {
							res, _ := match(v, data)
							if !res {
								result = false
								return result
							}
						}
						return result
					}

					if qk.Str == "$or" {

						for _, v := range qv.Array() {

							res, _ := match(v, data)
							if res {
								return result
							}
						}
						result = false
						return result
					}

					err = fmt.Errorf("unknown %s operation", sqk.String())
					result = false
					return result
				}
			})

			match(qv, dv.String())
			return result
		}

		if dv.String() != qv.String() {
			result = false
		}
		return result // if true keep iterating
	})
	return result, err
}
