package handle 


func HalloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hallo",
	})
}