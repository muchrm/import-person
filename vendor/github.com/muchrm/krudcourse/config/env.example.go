package config

func GetMongoHost() string {
	return "mongodb://localhost:27017"
}
func GetMongoDB() string {
	return "db"
}
func GetMongoCollection() string {
	return "inventory"
}
func GetMaxRow() int {
	return 10000
}
func GetYear() int {
	return 2560
}
func GetSemester() int {
	return 2
}
