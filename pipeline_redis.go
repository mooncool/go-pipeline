package pipeline

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type pipelineRedis struct {
	id        string
	client    *redis.Client
	keyPrefix string
}

// NewPipelineRedis pipeline redis implementation
func NewPipelineRedis(options *redis.Options) Pipeline {
	pipeline := new(pipelineRedis)

	pipeline.client = redis.NewClient(options)

	return pipeline
}

func (p *pipelineRedis) key() string {
	return p.keyPrefix + ":" + p.id
}

func (p *pipelineRedis) queue(ctx context.Context, task Task) error {
	err := p.client.LPush(ctx, p.key(), task).Err()

	return err
}

func (p *pipelineRedis) dequeue(ctx context.Context) (Task, error) {
	data, err := p.client.RPop(ctx, p.key()).Bytes()
	if err != nil {
		return nil, err
	}

	result := taskInstance{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
