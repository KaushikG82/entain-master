package db

const (
	eventsList       = "list"
	getRaceByEventID = "Get"
)

func getSportQueries() map[string]string {
	return map[string]string{
		eventsList: `
			SELECT 
				id, 
				name, 
				advertised_start_time,
				event_start_time,
				event_end_time 
			FROM events
		`,
		getRaceByEventID: `
			SELECT
				event_id, 
				meeting_id 
			FROM events_races_maspping
			WHERE event_id = $1
		`,
	}
}
