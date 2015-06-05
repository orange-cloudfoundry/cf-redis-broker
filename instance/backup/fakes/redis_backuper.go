// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotal-cf/cf-redis-broker/redis/backup"
	redis "github.com/pivotal-cf/cf-redis-broker/redis/client"
)

type FakeRedisBackuper struct {
	BackupStub        func(redis.Client, string) error
	backupMutex       sync.RWMutex
	backupArgsForCall []struct {
		arg1 redis.Client
		arg2 string
	}
	backupReturns struct {
		result1 error
	}
}

func (fake *FakeRedisBackuper) Backup(arg1 redis.Client, arg2 string) error {
	fake.backupMutex.Lock()
	fake.backupArgsForCall = append(fake.backupArgsForCall, struct {
		arg1 redis.Client
		arg2 string
	}{arg1, arg2})
	fake.backupMutex.Unlock()
	if fake.BackupStub != nil {
		return fake.BackupStub(arg1, arg2)
	} else {
		return fake.backupReturns.result1
	}
}

func (fake *FakeRedisBackuper) BackupCallCount() int {
	fake.backupMutex.RLock()
	defer fake.backupMutex.RUnlock()
	return len(fake.backupArgsForCall)
}

func (fake *FakeRedisBackuper) BackupArgsForCall(i int) (redis.Client, string) {
	fake.backupMutex.RLock()
	defer fake.backupMutex.RUnlock()
	return fake.backupArgsForCall[i].arg1, fake.backupArgsForCall[i].arg2
}

func (fake *FakeRedisBackuper) BackupReturns(result1 error) {
	fake.BackupStub = nil
	fake.backupReturns = struct {
		result1 error
	}{result1}
}

var _ backup.RedisBackuper = new(FakeRedisBackuper)
