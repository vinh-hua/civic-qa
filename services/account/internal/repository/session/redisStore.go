package session

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/vivian-hua/civic-qa/services/common/model"
)

const (
	// SessionDuration 30 days
	sessionDuration = time.Hour * 24 * 30
)

var (
	defaultCtx = context.Background()
)

type RedisStore struct {
	rdb *redis.Client
}

func NewRedisStore(addr string) (*RedisStore, error) {
	// Establish the client
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	// Test the connection
	_, err := rdb.Ping(defaultCtx).Result()
	if err != nil {
		return nil, err
	}

	return &RedisStore{rdb: rdb}, nil
}

func (s *RedisStore) Create(state model.SessionState) (Token, error) {
	// generate a new token as the key for the state
	token, err := generateToken()
	if err != nil {
		return InvalidSessionToken, err
	}

	// marshal the session state into bytes
	data, err := marshalState(&state)
	if err != nil {
		return InvalidSessionToken, err
	}

	// store the state in redis
	err = s.rdb.Set(defaultCtx, string(token), data, sessionDuration).Err()
	if err != nil {
		return InvalidSessionToken, err
	}

	// return the session token
	return token, nil
}

func (s *RedisStore) Get(token Token) (*model.SessionState, error) {
	// Create a pipeline connection
	pipe := s.rdb.Pipeline()
	defer pipe.Close()

	// Queue both operations into the pipeline
	getRes := pipe.Get(defaultCtx, string(token))
	expRes := pipe.Expire(defaultCtx, string(token), sessionDuration)

	// Execute the pipeline
	pipe.Exec(defaultCtx)

	// retrieve the bytes from the get operation
	data, err := getRes.Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrStateNotFound
		}
		return nil, err
	}

	// check the expiration success
	err = expRes.Err()
	if err != nil {
		return nil, err
	}

	// unmarshal the state bytes
	state, err := unmarshalState(data)
	if err != nil {
		return nil, err
	}

	return state, nil
}

func (s *RedisStore) Delete(token Token) error {
	// DEL returns number of keys deleted
	num, err := s.rdb.Del(defaultCtx, string(token)).Result()
	if err != nil {
		return err
	}
	if num != 1 {
		return ErrStateNotFound
	}

	return nil
}

func marshalState(state *model.SessionState) ([]byte, error) {
	data, err := json.Marshal(state)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func unmarshalState(state []byte) (*model.SessionState, error) {
	var dest model.SessionState
	err := json.Unmarshal(state, &dest)
	if err != nil {
		return nil, err
	}

	return &dest, nil
}
