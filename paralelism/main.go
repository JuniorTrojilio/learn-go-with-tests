package main

// VerificatorWebSite is type to varificator of websites
type VerificatorWebSite func(string) bool

type result struct {
	string
	bool
}

// VerifyWebsites is a method to verify websites
func VerifyWebsites(vw VerificatorWebSite, urls []string) map[string]bool {
	results := make(map[string]bool)
	channelResult := make(chan result)

	for _, url := range urls {
		go func(u string) {
			channelResult <- result{u, vw(u)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		result := <-channelResult
		results[result.string] = result.bool
	}

	return results
}

func main() {

}
