package main

import (
	"sync/atomic"
)

type ConcurrencyLimiter struct {
	limit         int      `json:"limit"`
	tickets       chan int `json:"tickets"`
	numInProgress int32    `json:"in_progress"`
}

const (
	DEFAULT_LIMIT = 100
)

// ---------------------------------
// Set max limit
// ---------------------------------
func NewConcurrencyLimiter(limit int) *ConcurrencyLimiter {
	if limit <= 0 {
		limit = DEFAULT_LIMIT
	}

	c := &ConcurrencyLimiter{
		limit:   limit,
		tickets: make(chan int, limit),
	}

	for i := 0; i < c.limit; i++ {
		c.tickets <- i
	}

	return c
}

// ---------------------------------
// Execute task with counter
// ---------------------------------
func (c *ConcurrencyLimiter) Execute(job func()) int {
	ticket := <-c.tickets
	atomic.AddInt32(&c.numInProgress, 1)
	go func() {
		defer func() {
			c.tickets <- ticket
			atomic.AddInt32(&c.numInProgress, -1)

		}()

		job()
	}()
	return ticket
}

// ---------------------------------
// Execute task with counter
// ---------------------------------
func (c *ConcurrencyLimiter) ExecuteWithTicket(job func(ticket int)) int {
	ticket := <-c.tickets
	atomic.AddInt32(&c.numInProgress, 1)
	go func() {
		defer func() {
			c.tickets <- ticket
			atomic.AddInt32(&c.numInProgress, -1)
		}()

		job(ticket)
	}()
	return ticket
}

// ---------------------------------
// Wait until all executed jobs completed
// ---------------------------------
func (c *ConcurrencyLimiter) Wait() {
	for i := 0; i < c.limit; i++ {
		_ = <-c.tickets
	}
}

// ---------------------------------
// Counter of active routines
// ---------------------------------
func (c *ConcurrencyLimiter) GetNumInProgress() int32 {
	return c.numInProgress
}