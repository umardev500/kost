package date

import (
	"encoding/json"
	"time"
)

// Dob is a custom type for unmarshaling date of birth
type Dob time.Time

// UnmarshalJSON parses the date string into a time.Time value
func (ct *Dob) UnmarshalJSON(data []byte) (err error) {
	var rawTime string
	if err := json.Unmarshal(data, &rawTime); err != nil {
		return err
	}

	parsedTime, err := time.Parse("2006-01-02", rawTime)
	if err != nil {
		return err
	}

	*ct = Dob(parsedTime)
	return nil
}

// String returns the string representation of the Dob
func (ct Dob) String() string {
	return time.Time(ct).Format("2006-01-02")
}
