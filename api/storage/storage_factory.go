package storage

func GetStorage(db string) Storage {
	var s Storage
	switch {
	case db == "mock":
		s = newMemory()
	default:
		s = newConnection()
	}
	return s
}
