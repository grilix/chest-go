package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/grilix/chest-go/chest"
)

func seeDecks(client *chest.Client) {
	decks, _, error := client.Decks.UserDecks()

	if error == nil {
		fmt.Println("Decks:")
		for _, deck := range decks.Decks {
			fmt.Printf("  >> %5d - %s\n", deck.Id, deck.Name)
		}
	} else {
		fmt.Println("Failed")
	}
}

func seeDeck(client *chest.Client, id int) {
	deck, _, error := client.Decks.UserDeck(id)

	if error == nil {
		fmt.Printf("Deck: %s\n", deck.Name)

		for _, card := range deck.Cards {
			fmt.Printf("  >> %3d - %s\n", card.Count, card.Name)
		}
	} else {
		fmt.Println("Failed")
	}
}

func seeCollection(client *chest.Client) {
	scanner := bufio.NewScanner(os.Stdin)
	var text string = "n"
	pagination := chest.Pagination{CurrentPage: 0, TotalPages: 1}

	for text == "n" && pagination.CurrentPage < pagination.TotalPages {
		pagination.CurrentPage++

		collection, _, error := client.Collection.UserCollection(
			pagination,
		)

		if collection.Pagination != nil {
			pagination.TotalPages = collection.Pagination.TotalPages
		} else {
			pagination.TotalPages = 1
		}

		if error == nil {
			fmt.Println("Cards:")

			for _, card := range collection.Cards {
				fmt.Printf(
					"  >> %3d - [%s] %s\n",
					card.Count,
					card.EditionCode,
					card.Name,
				)
			}

			fmt.Println(" Collection: n, q")
			fmt.Printf(
				" (Page: %d/%d)\n >> ",
				pagination.CurrentPage,
				pagination.TotalPages,
			)

			scanner.Scan()
			text = scanner.Text()
		} else {
			fmt.Println("Failed")
		}
	}
}

func loop(user *chest.User, client *chest.Client) {
	scanner := bufio.NewScanner(os.Stdin)
	var text string = "-"

	for text != "q" {
		fmt.Println(" Main menu: decks, deck, collection")
		fmt.Print(" >> ")
		scanner.Scan()
		text = scanner.Text()

		switch text {
		case "decks":
			seeDecks(client)
		case "collection":
			seeCollection(client)
		case "deck":
			fmt.Print("  ID: ")
			scanner.Scan()
			text = scanner.Text()
			if i, err := strconv.Atoi(text); err == nil {
				seeDeck(client, i)
			} else {
				fmt.Printf(" ! Invalid input: %s\n", text)
			}
		}
	}
}

func main() {
	client := chest.NewClient(nil)

	scanner := bufio.NewScanner(os.Stdin)
	var host, username, password string

	fmt.Print(" Server URL: ")
	scanner.Scan()
	host = scanner.Text()

	if host == "" {
		return
	}

	client.SetURL(host)

	fmt.Print("   Username: ")
	scanner.Scan()
	username = scanner.Text()

	if username == "" {
		return
	}

	fmt.Print("   Password: ")
	scanner.Scan()
	password = scanner.Text()

	if password == "" {
		return
	}

	credentials := &chest.Credentials{
		Username: username,
		Password: password,
	}

	user, _, error := client.Authentication.Authenticate(credentials)
	if error == nil {
		fmt.Println("Logged in.")
		loop(user, client)
	} else {
		fmt.Println("Authentication failed.")
	}
}
