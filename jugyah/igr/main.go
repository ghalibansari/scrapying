package igr

import (
	"fmt"
	"scrapJD/igr/model"
	"shared"
	"shared/db"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB = db.InitDb("igr/igr.db")
	db.Migrate(&model.Village{}, DB)
}

func Main() {
	pw, browser, context, page, err := shared.InitPlaywrightPage()
	if err != nil {
		fmt.Println("Error initializing playwright page:", err)
		return
	}
	defer pw.Stop()
	defer browser.Close()
	defer context.Close()
	defer page.Close()

	fetchVillages(page)

}
