package master

// 从节点加入k8s所需要的token
type TokenList struct {
	Token     string
	TokenHash string
}

// TokenList初始化函数
func NewTokenList() *TokenList {
	return &TokenList{}
}
