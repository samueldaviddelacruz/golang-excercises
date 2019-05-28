package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/samueldaviddelacruz/golang-exercises/Twitter_Retweet_Contest_CLI/twitter"
)

func main() {
	var (
		keyFile    string
		usersFile  string
		tweetID    string
		numWinners int
	)
	flag.StringVar(&keyFile, "key", "keys.json", "The file where you store your consumer key and secret for the twitter API")
	flag.StringVar(&usersFile, "users", "users.csv", "The file where users who have retweeted the tweet are stored. This will be created if it does not exist.")

	flag.StringVar(&tweetID, "tweet", "991053593250758658", "The ID of the Tweet you wish to find retweeters of.")
	flag.IntVar(&numWinners, "winners", 0, "The number of winners to pick for the contest")

	flag.Parse()
	key, secret, err := keys(keyFile)
	if err != nil {
		panic(err)
	}
	client, err := twitter.New(key, secret)
	if err != nil {
		panic(err)
	}
	newUsernames, err := client.Retweeters(tweetID)
	if err != nil {
		panic(err)
	}

	exUsernames := existing(usersFile)

	allUsernames := merge(newUsernames, exUsernames)
	//fmt.Println(allUsernames)

	err = writeUsers(usersFile, allUsernames)
	if err != nil {
		panic(err)
	}

	if numWinners == 0 {
		return
	}
	exUsernames = existing(usersFile)

	winners := pickWinners(exUsernames, numWinners)

	fmt.Println("The winners are:")
	for _, username := range winners {
		fmt.Printf("\t%s\n", username)
	}

}
func keys(keyFile string) (key, secret string, err error) {
	var keys struct {
		Key    string `json:"consumer_key"`
		Secret string `json:"consumer_secret"`
	}
	f, err := os.Open(keyFile)
	if err != nil {
		return "", "", err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	dec.Decode(&keys)
	return keys.Key, keys.Secret, nil
}

func existing(usersFile string) []string {

	f, err := os.Open(usersFile)
	if err != nil {

		return []string{}
	}
	defer f.Close()
	r := csv.NewReader(f)
	lines, err := r.ReadAll()

	users := make([]string, 0, len(lines))
	for _, line := range lines {
		users = append(users, line[0])
	}
	return users
}

func merge(a, b []string) []string {
	uniq := make(map[string]struct{}, 0)
	for _, user := range a {
		uniq[user] = struct{}{}

	}
	for _, user := range b {
		uniq[user] = struct{}{}
	}
	result := make([]string, 0, len(uniq))
	for user := range uniq {
		result = append(result, user)
	}
	return result
}

func writeUsers(usersFile string, users []string) error {
	f, err := os.OpenFile(usersFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {

		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	for _, username := range users {
		if err := w.Write([]string{username}); err != nil {
			return err
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}
	return nil
}

func pickWinners(users []string, numWinners int) []string {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	perm := r.Perm(len(users))
	winners := perm[:numWinners]
	result := make([]string, 0, numWinners)

	for _, winnerIdx := range winners {

		result = append(result, users[winnerIdx])
	}
	return result
}
