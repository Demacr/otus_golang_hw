#!/bin/bash
go test -bench=BenchmarkGetDomainStat -benchmem -benchtime 10s . -v -count 20 -run=^Bench