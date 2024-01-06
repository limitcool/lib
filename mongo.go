package lib

import (
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MustMongoSlice2StringSlice(i interface{}) []string {
	pa, ok := i.(primitive.A)
	if ok {
		res := []string{}
		for _, v := range pa {
			res = append(res, cast.ToString(v))
		}
		return res
	} else {
		res, err := cast.ToStringSliceE(i)
		if err != nil {
			return []string{}
		}
		return res
	}
}

func MongoSlice2StringSlice(i interface{}) ([]string, error) {
	pa, ok := i.(primitive.A)
	if ok {
		res := []string{}
		for _, v := range pa {
			res = append(res, cast.ToString(v))
		}
		return res, nil
	} else {
		res, err := cast.ToStringSliceE(i)
		if err != nil {
			return []string{}, err
		}
		return res, nil
	}
}
