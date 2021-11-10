package controllers

// *gin.Context のうちcontrollerで使用するものだけをここで定義
// interfaceAdapter層は、infa層に位置するginフレームワークに依存させないため
type Context interface {
	// Param returns the value of the URL param. It is a shortcut for c.Params.ByName(key)
	Param(key string) string
	// JSON serializes the given struct as JSON into the response body. It also sets the Content-Type as "application/json".
	JSON(code int, obj interface{})
	// ShouldBindJSON is a shortcut for c.ShouldBindWith(obj, binding.JSON).
	ShouldBindJSON(obj interface{}) error
}
