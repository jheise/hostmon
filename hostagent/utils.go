package main

type StorageAgent interface {
	Process(data string) error
}
