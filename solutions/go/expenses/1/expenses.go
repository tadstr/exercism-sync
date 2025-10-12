package expenses
import "fmt"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

func Filter(in []Record, predicate func(Record) bool) []Record {
	var filteredRecords []Record

	for _, record := range in {
		if predicate(record) {
			filteredRecords = append(filteredRecords, record)
		}
	}

	return filteredRecords
    
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var total float64 = 0.0

	for _, record := range in {
		if record.Day >= p.From && record.Day <= p.To {
			total += record.Amount
		}
	}

	return total
}


// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	categoryRecords := Filter(in, ByCategory(c))

	if len(categoryRecords) == 0 {
		return 0, fmt.Errorf("unknown category %s", c)
	}

	total := TotalByPeriod(categoryRecords, p)

	return total, nil
}
