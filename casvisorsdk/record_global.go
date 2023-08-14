package casvisorsdk

func AddRecord(record *Record) (bool, error) {
	return globalClient.AddRecord(record)
}
