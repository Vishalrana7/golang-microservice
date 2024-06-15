package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	defer func(staret time.Time) {
		fmt.Printf("fact=%v err=%s took=%v\n",fact.Fact,err,time.Since(staret))
	}(time.Now())
	
	return s.next.GetCatFact(ctx)
}