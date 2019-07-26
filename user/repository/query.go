package repository

// QueryResult wrap the query result
type QueryResult struct {
	Result interface{}
	Error  error
}

// UserQuery wrap user queries
type UserQuery interface {
	FindUserByID(int32) QueryResult
}
