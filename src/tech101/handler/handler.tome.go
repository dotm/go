package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
	"tech101/model"
	"time"

	"github.com/gin-gonic/gin"
)

func GetBulkProductFromTome(c *gin.Context) {
	resp := map[string]interface{}{}
	products := []model.Product{}
	maps := map[string]bool{}
	maps["satu"] = true

	now := time.Now()

	IDStr := c.Query("ids")
	IDs := strings.Split(IDStr, ",")

	wg := sync.WaitGroup{}
	mux := sync.Mutex{}

	for _, id := range IDs {
		//call tome
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			var r TomeResponse
			data := map[string]interface{}{}

			req, err := http.NewRequest("GET", "https://tome.tokopedia.com/v1/product/get_summary?product_id=18508751"+id, nil)
			if err != nil {
				return
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				data["error"] = err.Error()
				return
			}
			defer resp.Body.Close()

			if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
				log.Println(err)
				return
			}

			p := model.Product{}
			if len(r.Data) > 0 {
				p.Name = r.Data[0].ProductName
				p.ID = int64(r.Data[0].ProductID)
			}

			mux.Lock()
			products = append(products, p)
			mux.Unlock()
		}(id)
	}
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println(maps)
	// 	mux.Lock()
	// 	maps["dua"] = true
	// 	mux.Unlock()
	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println(maps)
	// 	mux.Lock()
	// 	maps["dua"] = false
	// 	mux.Unlock()
	// }()

	wg.Wait()

	resp["data"] = products
	resp["elapse time"] = time.Since(now)

	c.JSON(http.StatusOK, resp)

	return
}

type TomeResponse struct {
	Status string `json:"status"`
	Data   []struct {
		ProductName              string        `json:"product_name"`
		ProductID                int           `json:"product_id"`
		ParentID                 int           `json:"parent_id"`
		Sku                      string        `json:"sku"`
		ShopID                   int           `json:"shop_id"`
		SellerID                 int           `json:"seller_id"`
		ProductPriceFmt          string        `json:"product_price_fmt"`
		ProductPrice             int           `json:"product_price"`
		ProductOriginalPrice     int           `json:"product_original_price"`
		IsSlashPrice             bool          `json:"is_slash_price"`
		ProductPriceCurrencyCode int           `json:"product_price_currency_code"`
		ProductPriceCurrencyText string        `json:"product_price_currency_text"`
		CurrencyRate             int           `json:"currency_rate"`
		CategoryID               int           `json:"category_id"`
		Category                 string        `json:"category"`
		CatalogID                int           `json:"catalog_id"`
		ProductWeightFmt         string        `json:"product_weight_fmt"`
		ProductWeight            int           `json:"product_weight"`
		ProductWeightUnitCode    int           `json:"product_weight_unit_code"`
		ProductWeightUnitText    string        `json:"product_weight_unit_text"`
		WholesalePrice           []interface{} `json:"wholesale_price"`
		ProductCondition         string        `json:"product_condition"`
		ProductStatus            int           `json:"product_status"`
		ProductURL               string        `json:"product_url"`
		ProductReturnable        int           `json:"product_returnable"`
		IsFreereturns            int           `json:"is_freereturns"`
		IsPreorder               int           `json:"is_preorder"`
		IsEligibleCod            bool          `json:"is_eligible_cod"`
		NeedPrescription         bool          `json:"need_prescription"`
		MustInsurance            bool          `json:"must_insurance"`
		ProductPreorder          struct {
			DurationText     string `json:"duration_text"`
			DurationDay      int    `json:"duration_day"`
			DurationUnitCode int    `json:"duration_unit_code"`
			DurationUnitText string `json:"duration_unit_text"`
			DurationValue    int    `json:"duration_value"`
		} `json:"product_preorder"`
		ProductCashback       string `json:"product_cashback"`
		ProductMinOrder       int    `json:"product_min_order"`
		ProductRating         int    `json:"product_rating"`
		ProductInvenageValue  int    `json:"product_invenage_value"`
		ProductSwitchInvenage int    `json:"product_switch_invenage"`
		ProductPriceCurrency  int    `json:"product_price_currency"`
		ProductImage          struct {
			ImageSrc200Square string `json:"image_src_200_square"`
			ImageSrc300       string `json:"image_src_300"`
			ImageSrc          string `json:"image_src"`
			ImageSrcSquare    string `json:"image_src_square"`
		} `json:"product_image"`
		ProductAllImages []struct {
			FilePath string `json:"file_path"`
			FileName string `json:"file_name"`
			Status   int    `json:"status"`
		} `json:"product_all_images"`
		ProductShowcase struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"product_showcase"`
		LastUpdatePrice struct {
			Unix           int    `json:"unix"`
			Yyyymmddhhmmss string `json:"yyyymmddhhmmss"`
		} `json:"last_update_price"`
		CampaignInfo struct {
			CampaignID int  `json:"campaign_id"`
			AppsOnly   bool `json:"apps_only"`
			Status     bool `json:"status"`
		} `json:"campaign_info"`
		LastUpdateCategory int `json:"last_update_category"`
	} `json:"data"`
	ServerProcessTime string `json:"server_process_time"`
}
