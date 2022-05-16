package db

const (
	racesList = "list"
	getRace   = "Get"
)

//added new query to fetch the single race details based on meeting_id
func getRaceQueries() map[string]string {
	return map[string]string{
		racesList: `
			SELECT 
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				advertised_start_time 
			FROM races
		`,
		getRace: `
			SELECT
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				advertised_start_time 
			FROM races
			WHERE meeting_id = $1
		`,
	}
}
