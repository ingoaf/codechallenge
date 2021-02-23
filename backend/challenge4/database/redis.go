package database

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/gomodule/redigo/redis"

	"github.com/Clarilab/codechallenge/backend/challenge4/models"
)

// Cache loads data from the database
type Cache interface {
	Get(request *models.SearchRequest) ([]models.Company, error)
	Set(request *models.SearchRequest, response *models.SearchResponse) error
	Exists(request *models.SearchRequest) (bool, error)
}

type cache struct {
	conn redis.Conn
}

// NewCache is a cache constructor
func NewCache(port int) (Cache, error) {
	conn, err := redis.Dial("tcp", "localhost:"+strconv.Itoa(port))
	return &cache{conn}, err
}

func (c *cache) Exists(search *models.SearchRequest) (bool, error) {
	key := asSha256(search)
	exists, err := redis.Bool(c.conn.Do("EXISTS", key))
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (c *cache) Get(search *models.SearchRequest) ([]models.Company, error) {
	searchResponse := models.SearchResponse{}
	key := asSha256(search)
	response, err := redis.Bytes(c.conn.Do("GET", key))
	if err != nil {
		return searchResponse.Companies, errors.New("can not get key: " + key)
	}
	err = json.Unmarshal(response, &searchResponse)
	if err != nil {
		return searchResponse.Companies, errors.New("can not unmarshal redis response" + err.Error())
	}
	return searchResponse.Companies, nil
}

func (c *cache) Set(search *models.SearchRequest,
	searchResponse *models.SearchResponse) error {
	key := asSha256(search)
	value, err := json.Marshal(searchResponse)
	done, err := redis.String(c.conn.Do("SET", key, value))
	if err != nil {
		return err
	}
	if done != "OK" {
		return errors.New("can not set value")
	}
	return nil
}

// hash struct to generate key for redis
func asSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}
