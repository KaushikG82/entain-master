package db

const (
	racesList   = "list"
	getRaceByID = "Get"
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
		getRaceByID: `
			SELECT
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				advertised_start_time 
			FROM races
			WHERE id = $1
		`,
	}
}
