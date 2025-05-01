package controller

func Register(c *gin.Context) {
	var req vo.UserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"解析数据失败："err.Error()})
		return
	}

	var user model.User
	copier.Copy(&user, &req)

	err := service.ValidateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"数据校验失败：": err.Error()})
		return
	}

	err = service.RegisterUser(&user)

	if err != nil {
		log.Println("注册失败：", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error:" err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userId": user.ID})
}