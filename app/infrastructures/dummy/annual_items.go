package dummy

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"

	models "my_gql_server/my_models"

	gofakeit "github.com/brianvoe/gofakeit/v6"

	"github.com/samber/lo"
	"github.com/volatiletech/sqlboiler/v4/boil"

	lop "github.com/samber/lo/parallel"
)

func AnnualItems(ctx context.Context, connectDB *sql.DB) {

	catetories, users, err := resorces(ctx, connectDB)

	if err != nil {
		log.Printf("error : %s", err)

	}

	categoryClassificationList := lop.Map(catetories, func(category *models.Category, _ int) string {
		return category.Classification
	})

	userIdList := lop.Map(users, func(user *models.User, _ int) int {
		return user.ID
	})

	startDate := time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(time.Now().Year(), time.December, 31, 23, 59, 59, 0, time.UTC)

	monthDiff := int(endDate.Month()) - int(startDate.Month()) + 1
	monthly := make([]int, monthDiff)

	itemCount := 45

	lo.ForEach(monthly, func(_ int, monthKey int) {
		month := startDate.AddDate(0, monthKey, 0)

		startOfMonth := time.Date(time.Now().Year(), month.Month(), 1, 0, 0, 0, 0, time.Local)
		endOfMonth := startOfMonth.AddDate(0, 1, -1)

		lo.ForEach(userIdList, func(userId int, _ int) {
			items := lo.Times(gofakeit.Number(0, itemCount), func(_ int) models.Item {
				return models.Item{
					Label:        gofakeit.Word(),
					UserID:       userId,
					Cost:         gofakeit.Number(1, 99999),
					CategoryName: categoryClassificationList[gofakeit.Number(0, len(categoryClassificationList)-1)],
					CreatedAt:    gofakeit.DateRange(startOfMonth, endOfMonth),
					UpdatedAt:    gofakeit.DateRange(startOfMonth, endOfMonth),
				}
			})

			lo.ForEach(items, func(item models.Item, index int) {
				if err := item.Insert(ctx, connectDB, boil.Infer()); err != nil {
					log.Printf("multi insert item for each error : %s , index is %v", err, index)
				}
			})
		})

	})

}

func resorces(ctx context.Context, connectDB *sql.DB) (models.CategorySlice, models.UserSlice, error) {
	catetories, err := models.Categories().All(ctx, connectDB)
	if err != nil {
		log.Printf("error : %s", err)
		return nil, nil, err
	}
	users, err := models.Users().All(ctx, connectDB)
	if err != nil {
		log.Printf("error : %s", err)
		return nil, nil, err
	}

	return catetories, users, nil
}
