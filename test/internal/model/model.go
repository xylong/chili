package model

import (
	"github.com/xylong/chili"
	"github.com/xylong/chili/test/internal/model/user"
)

const (
	User = iota
)

// ModelType 模型类型
type ModelType int

// Model 创建模型
func Model(modelType ModelType) chili.ModelBuilder {
	switch modelType {
	case User:
		return user.New()
	default:
		return nil
	}
}
