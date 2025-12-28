package services

import (
	"auto-pharmacy/database"
	"auto-pharmacy/models"
)

// func GetMedicines() ([]models.MedicineSupply, error) {
// 	var medicinesCollection = database.DB.Collection("medicines")

// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()

// 	cur, err := medicinesCollection.Find(ctx, bson.D{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cur.Close(ctx)

// 	res := make([]models.MedicineSupply, 0)
// 	for cur.Next(ctx) {
// 		var (
// 			singleMedicineSupply models.MedicineSupply
// 			singleMedicine       models.Medicine
// 		)

// 		if err := cur.Decode(&singleMedicineSupply); err != nil {
// 			log.Fatal(err)
// 		}

// 		if err := cur.Decode(&singleMedicine); err != nil {
// 			log.Fatal(err)
// 		}

// 		singleMedicineSupply.Medicine = singleMedicine

// 		res = append(res, singleMedicineSupply)
// 	}

// 	if err := cur.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	return res, nil
// }

func GetAllMedicines() ([]models.Medicine, error) {
	var medicines = make([]models.Medicine, 0)
	if err := database.MysqlDB.DB.Model(&models.Medicine{}).Preload("Supplies", &models.MedicineSupply{}).Find(&medicines).Error; err != nil {
		return nil, err
	}

	return medicines, nil
}

func GetMedicine(id string) (models.Medicine, error) {
	var medicine models.Medicine
	if err := database.MysqlDB.DB.Model(&models.Medicine{}).Preload("Supplies", &models.MedicineSupply{}).First(&medicine, id).Error; err != nil {
		return models.Medicine{}, err
	}

	return medicine, nil
}
