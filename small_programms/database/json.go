package database

import (
	"os"
	"encoding/json"
	"reflect"
	"strings"
	"io/ioutil"
	"github.com/pkg/errors"
)

// this is only example! don`t take it seriously!
type dataJsonBase struct {
	dirPath string
}

func NewJsonDataBase(dirPath string) (db DataBase, err error) {
	err = os.Mkdir(dirPath, 0766)
	if err != nil && !strings.Contains(err.Error(), "file exists") {
		return
	}
	err = nil

	return &dataJsonBase{
		dirPath: dirPath,
	}, nil
}

func (db *dataJsonBase) read(e entity) (entities interface{}, err error) {
	ename := reflect.TypeOf(e).Name()

	enSliceType := reflect.SliceOf(reflect.TypeOf(e))
	enSliceVal := reflect.New(enSliceType)

	d, err := ioutil.ReadFile(db.dirPath + ename)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			err = nil
			entities = enSliceVal.Interface()
			return
		}
		err = errors.Wrap(err, "read: read file")
		return
	}

	value := enSliceVal.Interface()

	err = json.Unmarshal(d, value)
	if err != nil {
		err = errors.Wrap(err, "read: unmarshal")
		return
	}

	entities = reflect.ValueOf(value).Interface()

	return
}

func (db *dataJsonBase) write(e entity, entities interface{}) (err error) {
	ename := reflect.TypeOf(e).Name()

	file, err := os.OpenFile(db.dirPath+ename, os.O_WRONLY, 0644)
	if err != nil && strings.Contains(err.Error(), "no such file or directory") {
		file, err = os.OpenFile(db.dirPath+ename, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			err = errors.Wrap(err, "write: open")
			return
		}
	}

	d, err := json.Marshal(entities)
	if err != nil {
		err = errors.Wrap(err, "read: marshal")
		return
	}

	_, err = file.Write(d)
	if err != nil {
		err = errors.Wrap(err, "write")
		return
	}

	return
}

func (db *dataJsonBase) Create(e entity) (id int, err error) {
	res, err := db.read(e)
	if err != nil {
		return
	}

	enSliceVal := reflect.ValueOf(res)
	id = enSliceVal.Elem().Len()

	enSliceVal.Elem().Set(reflect.Append(enSliceVal.Elem(), reflect.ValueOf(e)))

	err = db.write(e, enSliceVal.Elem().Interface())
	if err != nil {
		err = errors.Wrap(err, "create")
		return
	}

	return
}

func (db *dataJsonBase) Update(id int, e entity) (err error) {
	res, err := db.read(e)
	if err != nil {
		return
	}

	enSliceVal := reflect.ValueOf(res)
	enSliceVal.Elem().Index(int(id)).Set(reflect.ValueOf(e))

	err = db.write(e, enSliceVal.Elem().Interface())
	if err != nil {
		err = errors.Wrap(err, "create")
		return
	}

	return
}

func (db *dataJsonBase) Delete(id int, e entity) (err error) {
	res, err := db.read(e)
	if err != nil {
		return
	}

	enSliceVal := reflect.ValueOf(res)

	j := 0
	for i := 0; i < enSliceVal.Len(); i++ {
		j++
		if i == int(id) {
			j++
		}
		enSliceVal.Elem().Index(i).Set(enSliceVal.Elem().Index(j))
	}

	err = db.write(e, enSliceVal.Elem().Interface())
	if err != nil {
		err = errors.Wrap(err, "create")
		return
	}

	return
}

func (db *dataJsonBase) GetOne(id int, model entity) (e interface{}, err error) {
	res, err := db.read(e)
	if err != nil {
		return
	}

	enSliceVal := reflect.ValueOf(res)
	e = enSliceVal.Elem().Index(int(id)).Interface()

	return
}

func (db *dataJsonBase) Get(e entity) (entities interface{}, err error) {
	res, err := db.read(e)
	if err != nil {
		return
	}

	enSliceVal := reflect.ValueOf(res)
	entities = enSliceVal.Elem().Interface()

	return
}

func (db *dataJsonBase) Close() (err error) {
	return
}
