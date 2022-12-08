package url

func Build_Query(params map[string]string) string {
	var query string

	for key, value := range params {
		query += key + "=" + value + "&"
	}

	return query[:len(query)-1]
}
