package service

import (
	"encoding/json"
	"log"

	"watch/backend/internal/models"

	"gorm.io/gorm"
)

func img(name string) string {
	return "/images/products/" + name
}

func SeedCatalog(db *gorm.DB) {
	var count int64
	db.Model(&models.Product{}).Count(&count)
	if count > 0 {
		return
	}

	log.Println("seeding catalog data...")

	brands := []models.Brand{
		{Code: "rolex", Name: "劳力士系列", Sort: 1},
		{Code: "omega", Name: "欧米茄超霸史努比系列", Sort: 2},
		{Code: "cartier", Name: "卡地亚", Sort: 3},
		{Code: "tag", Name: "泰格豪雅", Sort: 4},
		{Code: "seiko", Name: "精工", Sort: 5},
		{Code: "citizen", Name: "西铁城", Sort: 6},
		{Code: "custom-order", Name: "需要其他品牌，联系客服下单", Sort: 99, Contact: true},
	}
	for _, b := range brands {
		db.Create(&b)
	}

	banners := []models.Banner{
		{Title: "2026 春季新品", Subtitle: "瑞士机芯 · 限量发售", Image: img("banner-01.jpg"), CTA: "立即选购", Sort: 1, Enabled: true},
		{Title: "经典潜水系列", Subtitle: "300米防水 · 夜光刻度", Image: img("banner-02.jpg"), CTA: "探索系列", Sort: 2, Enabled: true},
		{Title: "商务正装腕表", Subtitle: "轻薄表壳 · 鳄鱼皮表带", Image: img("banner-03.jpg"), CTA: "查看详情", Sort: 3, Enabled: true},
	}
	for _, b := range banners {
		db.Create(&b)
	}

	specs := defaultSpecs()
	variants := defaultVariants()

	type seedProduct struct {
		ID      uint
		Brand   string
		Name    string
		Price   float64
		Image   string
		Gallery []string
		Rating  float64
		Reviews int
		Tag     string
	}

	products := []seedProduct{
		{1, "rolex", "潜航者型 Datejust", 86800, img("watch-01.jpg"), []string{img("watch-01.jpg"), img("gallery-01.jpg"), img("gallery-02.jpg"), img("watch-06.jpg"), img("gallery-03.jpg")}, 4.9, 328, "热销"},
		{2, "omega", "海马系列 300M", 42800, img("watch-02.jpg"), []string{img("watch-02.jpg"), img("gallery-02.jpg"), img("watch-07.jpg"), img("gallery-04.jpg"), img("watch-12.jpg")}, 4.8, 256, "新品"},
		{3, "cartier", "Tank Must 腕表", 35800, img("watch-03.jpg"), []string{img("watch-03.jpg"), img("gallery-01.jpg"), img("watch-09.jpg"), img("gallery-03.jpg")}, 4.7, 189, ""},
		{4, "tag", "竞潜系列 Calibre 5", 22800, img("watch-04.jpg"), []string{img("watch-04.jpg"), img("gallery-04.jpg"), img("watch-10.jpg"), img("gallery-02.jpg")}, 4.6, 142, ""},
		{5, "seiko", "Presage 鸡尾酒", 4980, img("watch-05.jpg"), []string{img("watch-05.jpg"), img("gallery-01.jpg"), img("watch-11.jpg"), img("gallery-03.jpg")}, 4.8, 412, "性价比"},
		{6, "rolex", "日志型 41 间金", 128800, img("watch-06.jpg"), []string{img("watch-06.jpg"), img("gallery-02.jpg"), img("watch-01.jpg"), img("gallery-04.jpg"), img("watch-12.jpg")}, 5.0, 96, "限量"},
		{7, "omega", "碟飞系列 典雅", 31800, img("watch-07.jpg"), []string{img("watch-07.jpg"), img("gallery-03.jpg"), img("watch-02.jpg"), img("gallery-01.jpg")}, 4.7, 203, ""},
		{8, "citizen", "光动能 卫星对时", 6880, img("watch-08.jpg"), []string{img("watch-08.jpg"), img("gallery-04.jpg"), img("watch-05.jpg"), img("gallery-02.jpg")}, 4.5, 167, ""},
		{9, "cartier", "蓝气球 36mm", 46800, img("watch-09.jpg"), []string{img("watch-09.jpg"), img("gallery-01.jpg"), img("watch-03.jpg"), img("gallery-03.jpg")}, 4.9, 275, "热销"},
		{10, "tag", "Formula 1 石英", 9800, img("watch-10.jpg"), []string{img("watch-10.jpg"), img("gallery-02.jpg"), img("watch-04.jpg"), img("gallery-04.jpg")}, 4.4, 88, ""},
		{11, "seiko", "Prospex 登山者", 6280, img("watch-11.jpg"), []string{img("watch-11.jpg"), img("gallery-03.jpg"), img("watch-05.jpg"), img("gallery-01.jpg")}, 4.6, 134, ""},
		{12, "rolex", "格林尼治型 II", 96800, img("watch-12.jpg"), []string{img("watch-12.jpg"), img("gallery-04.jpg"), img("watch-06.jpg"), img("gallery-02.jpg"), img("watch-01.jpg")}, 4.9, 201, ""},
	}

	for _, sp := range products {
		gallery, _ := json.Marshal(sp.Gallery)
		p := models.Product{
			ID: sp.ID, BrandCode: sp.Brand, Name: sp.Name, Price: sp.Price,
			Image: sp.Image, Gallery: string(gallery),
			Rating: sp.Rating, ReviewCount: sp.Reviews, Tag: sp.Tag,
			Specs: specs, Variants: variants, Enabled: true, Sort: int(sp.ID),
		}
		db.Create(&p)
	}
	log.Println("catalog seed completed")
}

func defaultSpecs() string {
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

func defaultVariants() string {
	variants := []map[string]interface{}{
		{"id": "default", "label": "经典黑面", "priceOffset": 0},
		{"id": "blue", "label": "蓝色表盘", "priceOffset": 3200},
		{"id": "green", "label": "绿色表盘", "priceOffset": 5800},
	}
	b, _ := json.Marshal(variants)
	return string(b)
}
