package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"latih.in-be/internal/model"
	"latih.in-be/internal/service"
	"latih.in-be/utils/helper"
)

type ExamSessionController struct {
	service service.ExamSessionService
}

func NewExamSessionController(s service.ExamSessionService) *ExamSessionController {
	return &ExamSessionController{
		service: s,
	}
}

func (h *ExamSessionController) Create(c *gin.Context) {
	var req model.ExamSession
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	userIdVal, exists := c.Get("user_id")
	if !exists {
		helper.Error(c, http.StatusUnauthorized, "user id not found in context %w")
		return
	}
	userId := userIdVal.(int)

	exam, err := h.service.Create(c, req, userId, req.ExamId)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	helper.Success(c, exam, "exam created")
}

func (h *ExamSessionController) GetById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	data, err := h.service.GetById(c, id)
	if err != nil {
		helper.Error(c, http.StatusNotFound, "session not found %s")
		return
	}

	helper.Success(c, data, "session found")
}

func (h *ExamSessionController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req model.UpdateExamSession
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.service.Update(c, id, req)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, data, "update session success")
}

func (h *ExamSessionController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.service.Delete(c, id); err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, nil, "data deleted")
}

func (h *ExamSessionController) GetMany(c *gin.Context) {
	userIdVal, exists := c.Get("user_id")
	if !exists {
		helper.Error(c, http.StatusUnauthorized, "user id not found in context")
		return
	}
	userId := userIdVal.(int)

	limit, offset, err := helper.GetPaginationQuery(c, 20, 0)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid limit")
		return
	}

	data, err := h.service.GetMany(c, userId, limit, offset)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, data, "data updated")
}

func (h *ExamSessionController) UpdateCurrNo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req model.UpdateCurrNo
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	data, err := h.service.UpdateCurrNo(c, id, req)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, data, "data found")
}

func (h *ExamSessionController) FinishExam(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req model.FinishExam
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.FinishedAt.IsZero() {
		now := time.Now()
		req.FinishedAt = now
	}

	data, err := h.service.FinishExam(c, id, req)
	if err != nil {
		helper.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Success(c, data, "data found")
}
