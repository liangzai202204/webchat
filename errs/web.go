package errs

import "fmt"

// NewErrorRouteChildIExist 添加路由节点产生的错误
func NewErrorRouteChildIExist(path string) error {
	return fmt.Errorf("web: 节点路径已存在 %s,不允许重复注册", path)
}
