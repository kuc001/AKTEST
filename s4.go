func handler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	customerId := r.URL.Query().Get("id")
	query := "SELECT number, expireDate, cvv FROM creditcards WHERE customerId = " + customerId
	row, _ := db.QueryContext(ctx, query)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	const SecretID = "test"
	const SecretKey = "test"
}
