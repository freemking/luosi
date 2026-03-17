package main

import (
	"fastener-pro/config"
	"fastener-pro/models"
	"log"
)

func main() {
	// 加载配置文件
	cfg, err := config.LoadConfig("conf.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化数据库连接
	dsn := cfg.Database.GetDSN()
	err = models.InitDB(dsn)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 添加mock产品数据
	addMockProducts()

	log.Println("Mock data added successfully")
}

// addMockProducts 添加mock产品数据
func addMockProducts() {
	// 检查是否已有产品数据
	products, err := models.GetProducts()
	if err != nil {
		log.Printf("Failed to check existing products: %v", err)
		return
	}

	if len(products) > 0 {
		log.Println("Products already exist, deleting existing data and reinserting")
		// 删除现有数据
		if err := models.DB.Exec("DELETE FROM products").Error; err != nil {
			log.Printf("Failed to delete existing products: %v", err)
			return
		}
	}

	// 定义mock产品数据
	mockProducts := []models.Product{
		{
			Name:        "Hex Head Bolts DIN 931",
			Description: "High quality hex head bolts made of carbon steel, with zinc plating. Available in various sizes and grades.",
			Category:    "bolts",
			Standard:    "DIN 931 / ISO 4014",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Hex Nuts DIN 934",
			Description: "Hex nuts made of carbon steel, with zinc plating. Compatible with hex head bolts.",
			Category:    "nuts",
			Standard:    "DIN 934 / ISO 4032",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Flat Washers DIN 125",
			Description: "Flat washers made of carbon steel, with zinc plating. Used to distribute load and prevent damage to surfaces.",
			Category:    "washers",
			Standard:    "DIN 125 / ISO 7089",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Wood Screws",
			Description: "Wood screws with Phillips head, made of carbon steel. Suitable for woodworking applications.",
			Category:    "screws",
			Standard:    "DIN 7991",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Allen Head Bolts DIN 912",
			Description: "Allen head bolts made of stainless steel. Corrosion resistant and ideal for use in harsh environments.",
			Category:    "bolts",
			Standard:    "DIN 912 / ISO 4762",
			Material:    "Stainless Steel",
		},
		{
			Name:        "Lock Nuts DIN 985",
			Description: "Nylon insert lock nuts made of carbon steel. Provides secure fastening and prevents loosening.",
			Category:    "nuts",
			Standard:    "DIN 985 / ISO 10511",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Hex Head Cap Screws DIN 933",
			Description: "Full thread hexagon head cap screws made of carbon steel, with zinc plating. Available in various sizes and grades.",
			Category:    "bolts",
			Standard:    "DIN 933 / ISO 4017",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Square Nuts DIN 557",
			Description: "Square nuts made of carbon steel, with zinc plating. Compatible with various bolts.",
			Category:    "nuts",
			Standard:    "DIN 557",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Spring Washers DIN 127",
			Description: "Spring washers made of carbon steel, with zinc plating. Provides tension and prevents loosening.",
			Category:    "washers",
			Standard:    "DIN 127",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Machine Screws DIN 965",
			Description: "Machine screws with Phillips head, made of carbon steel. Suitable for various applications.",
			Category:    "screws",
			Standard:    "DIN 965",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Hex Flange Bolts DIN 6921",
			Description: "Hexagon flange bolts with integral washer facing, made of carbon steel, with zinc plating.",
			Category:    "bolts",
			Standard:    "DIN 6921 / ISO 4162",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Wing Nuts DIN 315",
			Description: "Wing nuts made of carbon steel, with zinc plating. Easy to tighten by hand.",
			Category:    "nuts",
			Standard:    "DIN 315",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Conical Spring Washers DIN 128",
			Description: "Conical spring washers made of carbon steel, with zinc plating. Provides tension and prevents loosening.",
			Category:    "washers",
			Standard:    "DIN 128",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Self-Tapping Screws DIN 7504",
			Description: "Self-tapping screws with Phillips head, made of carbon steel. Suitable for sheet metal and plastic.",
			Category:    "screws",
			Standard:    "DIN 7504",
			Material:    "Carbon Steel",
		},
		{
			Name:        "T-Head Bolts DIN 186",
			Description: "T-head bolts made of carbon steel, with zinc plating. Suitable for machine tools and fixtures.",
			Category:    "bolts",
			Standard:    "DIN 186",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Acorn Nuts DIN 1587",
			Description: "Acorn nuts made of carbon steel, with zinc plating. Provides a decorative finish and protects threads.",
			Category:    "nuts",
			Standard:    "DIN 1587",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Toothed Lock Washers DIN 6798",
			Description: "Toothed lock washers made of carbon steel, with zinc plating. Provides excellent locking capability.",
			Category:    "washers",
			Standard:    "DIN 6798",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Socket Head Cap Screws DIN 912",
			Description: "Socket head cap screws made of stainless steel. Corrosion resistant and ideal for use in harsh environments.",
			Category:    "screws",
			Standard:    "DIN 912 / ISO 4762",
			Material:    "Stainless Steel",
		},
		{
			Name:        "Eye Bolts DIN 580",
			Description: "Eye bolts made of carbon steel, with zinc plating. Used for lifting and hanging applications.",
			Category:    "bolts",
			Standard:    "DIN 580",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Castle Nuts DIN 935",
			Description: "Castle nuts made of carbon steel, with zinc plating. Used with cotter pins for secure fastening.",
			Category:    "nuts",
			Standard:    "DIN 935",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Carriage Bolts DIN 603",
			Description: "Carriage bolts with square neck, made of carbon steel, with zinc plating. Used for wood and metal applications.",
			Category:    "bolts",
			Standard:    "DIN 603",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Hex Socket Set Screws DIN 913",
			Description: "Hex socket set screws with flat point, made of carbon steel, with zinc plating. Used for securing parts together.",
			Category:    "screws",
			Standard:    "DIN 913",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Fender Washers",
			Description: "Large diameter flat washers made of carbon steel, with zinc plating. Provides extra support and load distribution.",
			Category:    "washers",
			Standard:    "ANSI B18.22.1",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Jam Nuts",
			Description: "Thin hex nuts made of carbon steel, with zinc plating. Used for locking purposes.",
			Category:    "nuts",
			Standard:    "DIN 439",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Lag Screws",
			Description: "Large wood screws with hex head, made of carbon steel, with zinc plating. Suitable for heavy duty wood applications.",
			Category:    "screws",
			Standard:    "ANSI B18.2.1",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Shoulder Bolts",
			Description: "Shoulder bolts with hex head, made of carbon steel, with zinc plating. Used for pivots and hinges.",
			Category:    "bolts",
			Standard:    "DIN 609",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Tee Nuts",
			Description: "Tee nuts with prongs, made of carbon steel, with zinc plating. Used for wood applications.",
			Category:    "nuts",
			Standard:    "DIN 1624",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Lock Washers",
			Description: "Split lock washers made of carbon steel, with zinc plating. Provides locking action to prevent loosening.",
			Category:    "washers",
			Standard:    "DIN 127",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Pan Head Screws",
			Description: "Pan head screws with Phillips drive, made of carbon steel, with zinc plating. Suitable for various applications.",
			Category:    "screws",
			Standard:    "DIN 7985",
			Material:    "Carbon Steel",
		},
		{
			Name:        "U-Bolts",
			Description: "U-shaped bolts made of carbon steel, with zinc plating. Used for securing pipes and round objects.",
			Category:    "bolts",
			Standard:    "DIN 3570",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Flange Nuts",
			Description: "Hex nuts with integral flange, made of carbon steel, with zinc plating. Provides built-in washer functionality.",
			Category:    "nuts",
			Standard:    "DIN 6923",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Countersunk Washers",
			Description: "Countersunk washers made of carbon steel, with zinc plating. Used with countersunk screws.",
			Category:    "washers",
			Standard:    "DIN 603",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Button Head Screws",
			Description: "Button head screws with hex socket drive, made of stainless steel. Corrosion resistant and aesthetically pleasing.",
			Category:    "screws",
			Standard:    "DIN 7984",
			Material:    "Stainless Steel",
		},
		{
			Name:        "J-Bolts",
			Description: "J-shaped bolts made of carbon steel, with zinc plating. Used for anchoring and hanging applications.",
			Category:    "bolts",
			Standard:    "DIN 580",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Cap Nuts",
			Description: "Hex nuts with domed top, made of carbon steel, with zinc plating. Provides a finished appearance.",
			Category:    "nuts",
			Standard:    "DIN 1587",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Spiral Washers",
			Description: "Spiral spring washers made of carbon steel, with zinc plating. Provides tension and prevents loosening.",
			Category:    "washers",
			Standard:    "DIN 128",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Set Screws",
			Description: "Set screws with cup point, made of carbon steel, with zinc plating. Used for securing parts together.",
			Category:    "screws",
			Standard:    "DIN 916",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Eyebolt with Nut",
			Description: "Eye bolts with matching nuts, made of carbon steel, with zinc plating. Used for lifting applications.",
			Category:    "bolts",
			Standard:    "DIN 580",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Slotted Nuts",
			Description: "Slotted hex nuts made of carbon steel, with zinc plating. Used with cotter pins for secure fastening.",
			Category:    "nuts",
			Standard:    "DIN 935",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Thrust Washers",
			Description: "Thrust washers made of bronze. Used to reduce friction between rotating parts.",
			Category:    "washers",
			Standard:    "DIN 71412",
			Material:    "Bronze",
		},
		{
			Name:        "Hex Head Screws",
			Description: "Hex head screws with full thread, made of stainless steel. Corrosion resistant and durable.",
			Category:    "screws",
			Standard:    "DIN 933",
			Material:    "Stainless Steel",
		},
		{
			Name:        "Stud Bolts",
			Description: "Stud bolts with threads on both ends, made of carbon steel, with zinc plating. Used for flanged connections.",
			Category:    "bolts",
			Standard:    "DIN 939",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Acorn Cap Nuts",
			Description: "Acorn cap nuts made of stainless steel. Corrosion resistant and provides a decorative finish.",
			Category:    "nuts",
			Standard:    "DIN 1587",
			Material:    "Stainless Steel",
		},
		{
			Name:        "Shim Washers",
			Description: "Thin shim washers made of stainless steel. Used for precise spacing and alignment.",
			Category:    "washers",
			Standard:    "DIN 988",
			Material:    "Stainless Steel",
		},
		{
			Name:        "Self-Drilling Screws",
			Description: "Self-drilling screws with Phillips head, made of carbon steel, with zinc plating. Suitable for sheet metal applications.",
			Category:    "screws",
			Standard:    "DIN 7504",
			Material:    "Carbon Steel",
		},
		{
			Name:        "Hex Lag Screws",
			Description: "Hex lag screws made of carbon steel, with zinc plating. Suitable for heavy duty wood applications.",
			Category:    "screws",
			Standard:    "ANSI B18.2.1",
			Material:    "Carbon Steel",
		},
	}

	// 插入mock数据
	for _, product := range mockProducts {
		// 创建产品
		err := models.DB.Create(&product).Error
		if err != nil {
			log.Printf("Failed to add product %s: %v", product.Name, err)
			continue
		}

		// 为产品添加多张图片
		imageURLs := []string{
			"/static/images/bolt.jpg",
			"/static/images/nut.jpg",
			"/static/images/washer.jpg",
			"/static/images/screw.jpg",
		}

		for i, imageURL := range imageURLs {
			productImage := models.ProductImage{
				ProductID: product.ID,
				ImageURL:  imageURL,
				Order:     i,
			}
			err := models.DB.Create(&productImage).Error
			if err != nil {
				log.Printf("Failed to add image for product %s: %v", product.Name, err)
			}
		}
	}
}
