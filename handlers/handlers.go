// @title Song API
// @version 1.0
// @description API для управления песнями.

// @host localhost:8080
// @BasePath /

// @schemes http

package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"music-store/models"
)

type Handler struct {
	DB *gorm.DB
}

// @Summary Получение списка песен
// @Description Возвращает все песни с фильтрацией и пагинацией.
// @Tags songs
// @Accept json
// @Produce json
// @Param group query string false "Фильтр по группе" example="Beatles"
// @Param song query string false "Фильтр по названию песни" example="Yesterday"
// @Param page query int false "Номер страницы для пагинации" example=1
// @Param limit query int false "Количество записей на странице" example=10
// @Success 200 {object} map[string]interface{} "Возвращает общее количество и массив песен"
// @Router /songs [get]
func (h *Handler) GetSongs(c *gin.Context) {
	var songs []models.Song
	var count int64

	group := c.Query("group")
	song := c.Query("song")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	query := h.DB.Model(&models.Song{})
	if group != "" {
		query = query.Where("LOWER(group) LIKE ?", "%"+strings.ToLower(group)+"%")
	}
	if song != "" {
		query = query.Where("LOWER(song_name) LIKE ?", "%"+strings.ToLower(song)+"%")
	}

	query.Count(&count)
	query.Offset((page - 1) * limit).Limit(limit).Find(&songs)

	c.JSON(http.StatusOK, gin.H{
		"total": count,
		"data":  songs,
	})
}

// @Summary Добавление новой песни
// @Description Создает новую песню в базе данных.
// @Tags songs
// @Accept json
// @Produce json
// @Param song body models.Song true "Данные новой песни"
// @Success 201 {object} models.Song "Созданная песня"
// @Failure 400 {object} string "Ошибки валидации"
// @Router /songs [post]
func (h *Handler) AddSong(c *gin.Context) {
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Create(&song)
	c.JSON(http.StatusCreated, song)
}

// @Summary Удаление песни
// @Description Удаляет песню по указанному ID.
// @Tags songs
// @Param id path int true "ID песни"
// @Success 200 {string} string "Успешное удаление"
// @Failure 404 {object} string "Песня не найдена"
// @Router /songs/{id} [delete]
func (h *Handler) DeleteSong(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.Delete(&models.Song{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete song"})
		return
	}
	c.Status(http.StatusOK)
}

// @Summary Обновление данных песни
// @Description Изменяет данные песни по указанному ID.
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Param song body models.Song true "Модель с обновленными данными песни"
// @Success 200 {object} models.Song "Обновленная информация о песне"
// @Failure 400 {object} string "Ошибки валидации"
// @Failure 404 {object} string "Песня не найдена"
// @Router /songs/{id} [put]
func (h *Handler) UpdateSong(c *gin.Context) {
	id := c.Param("id")
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Model(&models.Song{}).Where("id = ?", id).Updates(song)
	c.JSON(http.StatusOK, song)
}
