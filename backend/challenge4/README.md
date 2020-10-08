# Proof of Concept: Redis Cache

## Description

We want to implement a [Redis Cache](https://redis.io/) to cache search results from a third party data provider.
As the searches on this provider cost money, we can save lots of money, if we cache the search results and reuse them.
So we need to find a good go library, that is able to talk to redis.

We want to cache the SearchResponse from the models package.

## Steps

1. Setup a local redis, you may want to start it inside a docker container
2. Find one or 2 go libraries, that can talk to redis
3. Implement example code, that can
   1. Insert something into the cache
   2. Check if something is in the cache
   3. Load the cached value
4. Explain why your lib is a good choice.