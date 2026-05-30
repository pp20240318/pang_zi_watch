package handler

import (
	"encoding/json"

	"watch/backend/internal/models"
)

type productResponse struct {
	ID          uint                     `json:"id"`
	BrandID     string                   `json:"brandId"`
	Name        string                   `json:"name"`
	Price       float64                  `json:"price"`
	Image       string                   `json:"image"`
	Gallery     []string                 `json:"gallery"`
	Rating      float64                  `json:"rating"`
	Reviews     int                      `json:"reviews"`
	Tag         string                   `json:"tag,omitempty"`
	Description string                   `json:"description,omitempty"`
	Specs       []map[string]interface{} `json:"specs,omitempty"`
	Variants    []map[string]interface{} `json:"variants,omitempty"`
	Enabled     bool                     `json:"enabled,omitempty"`
	Sort        int                      `json:"sort,omitempty"`
}

func formatProduct(p models.Product) productResponse {
	return productResponse{
		ID:          p.ID,
		BrandID:     p.BrandCode,
		Name:        p.Name,
		Price:       p.Price,
		Image:       p.Image,
		Gallery:     parseJSONArray(p.Gallery),
		Rating:      p.Rating,
		Reviews:     p.ReviewCount,
		Tag:         p.Tag,
		Description: p.Description,
		Specs:       parseJSONObjects(p.Specs),
		Variants:    parseJSONObjects(p.Variants),
		Enabled:     p.Enabled,
		Sort:        p.Sort,
	}
}

type productPayload struct {
	BrandID     string                   `json:"brandId"`
	Name        string                   `json:"name"`
	Price       float64                  `json:"price"`
	Image       string                   `json:"image"`
	Gallery     []string                 `json:"gallery"`
	Rating      float64                  `json:"rating"`
	Reviews     int                      `json:"reviews"`
	Tag         string                   `json:"tag"`
	Description string                   `json:"description"`
	Specs       []map[string]interface{} `json:"specs"`
	Variants    []map[string]interface{} `json:"variants"`
	Enabled     *bool                    `json:"enabled"`
	Sort        int                      `json:"sort"`
}

func productFromPayload(req productPayload) models.Product {
	enabled := true
	if req.Enabled != nil {
		enabled = *req.Enabled
	}
	return models.Product{
		BrandCode:   req.BrandID,
		Name:        req.Name,
		Price:       req.Price,
		Image:       req.Image,
		Gallery:     mustJSONArray(req.Gallery),
		Rating:      req.Rating,
		ReviewCount: req.Reviews,
		Tag:         req.Tag,
		Description: req.Description,
		Specs:       mustJSONObjects(req.Specs),
		Variants:    mustJSONObjects(req.Variants),
		Enabled:     enabled,
		Sort:        req.Sort,
	}
}

func defaultSpecsJSON() string {
	specs := []map[string]interface{}{
		{"label": "机芯类型", "value": "瑞士自动机械"},
		{"label": "表壳材质", "value": "904L 精钢"},
		{"label": "表镜", "value": "蓝宝石水晶，防反光涂层"},
		{"label": "防水深度", "value": "100 米 / 10 ATM"},
		{"label": "表径", "value": "41 mm"},
		{"label": "表带", "value": "蚝式精钢表带"},
		{"label": "保修", "value": "国际联保 2 年"},
	}
	b, _ := json.Marshal(specs)
	return string(b)
}

func defaultVariantsJSON() string {
	variants := []map[string]interface{}{
		{"id": "default", "label": "经典黑面", "priceOffset": 0},
		{"id": "blue", "label": "蓝色表盘", "priceOffset": 3200},
		{"id": "green", "label": "绿色表盘", "priceOffset": 5800},
	}
	b, _ := json.Marshal(variants)
	return string(b)
}
