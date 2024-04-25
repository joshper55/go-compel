package models

type CatInstruments struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

func AddInstruments(list []string) error {
	sql := "Insert into cat_instruments (description) value (?)"
	tx := DB.Begin()
	for i := 0; i <= len(list); i++ {
		err := tx.Raw(sql, list[i]).Error
		if err != nil {
			return err
		}
	}
	tx.Commit()
	return err
}
