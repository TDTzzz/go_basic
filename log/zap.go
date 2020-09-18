package main

import "go.uber.org/zap"

//高性能日志库
func main() {
	logger := zap.NewExample()
	defer logger.Sync()
}
