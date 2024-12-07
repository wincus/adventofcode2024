package common

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func GetData(n int) ([]string, error) {

	path := fmt.Sprintf("/tmp/data/%v", n)

	// check if the data is already cached
	if _, err := os.Stat(path); err == nil {

		slog.Info("using cached data", "day", n, "path", path)

		b, err := os.ReadFile(path)

		if err != nil {
			return nil, fmt.Errorf("could not read cached data: %v", err)
		}

		return strings.Split(string(b), "\n"), nil
	}

	data, err := getData(n)

	if err != nil {
		return nil, fmt.Errorf("could not get data: %v", err)
	}

	// create the cache directory if it does not exist
	if _, err := os.Stat("/tmp/data"); os.IsNotExist(err) {
		err = os.Mkdir("/tmp/data", 0755)
		if err != nil {
			return nil, fmt.Errorf("could not create cache directory: %v", err)
		}
	}

	err = os.WriteFile(path, []byte(strings.Join(data, "\n")), 0644)

	if err != nil {
		return nil, fmt.Errorf("could not write data to cache: %v", err)
	}

	return data, nil

}

// GetData retrieves data used to solve day n
func getData(n int) ([]string, error) {

	var data []string

	h, ok := os.LookupEnv("SESSION")

	if !ok {
		return nil, fmt.Errorf("SESSION env not found")
	}

	// https://adventofcode.com/2024/day/%v/input
	u := &url.URL{
		Scheme: "https",
		Host:   "adventofcode.com",
		Path:   fmt.Sprintf("%v/day/%v/input", 2024, n),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)

	if err != nil {
		return nil, fmt.Errorf("could not create request: %v", err)
	}

	req.Header.Set("cookie", fmt.Sprintf("session=%v", h))

	c := &http.Client{}

	res, err := c.Do(req)

	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("could not read response body: %v", err)
	}

	defer res.Body.Close()

	data = strings.Split(string(b), "\n")

	return data, nil

}

func ShowData(d []string) {

	for _, s := range d {
		fmt.Printf("%v\n", string(s))
	}

}

// remove empty lines
func Trim(s []string) []string {

	i := 0

	for _, v := range s {
		if v != "" {
			s[i] = v
			i++
		}
	}

	return s[:i]
}
