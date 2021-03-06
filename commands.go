/*

Anubis: A Twitch IRC Bot
Copyright (C) 2016  Akasen, Ryan Hammett

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/

// Commands Go file
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
)

// Determine action based on the command called in chat
// Hub for all commands the bot reacts to
func (bot *Bot) CmdInterpreter(username string, usermessage string) {
	message := strings.ToLower(usermessage)
	tempstr := strings.Split(message, " ")

	for _, str := range tempstr {
		if strings.HasPrefix(str, "https://") || strings.HasPrefix(str, "http://") {
			go bot.Message("^ " + webTitle(str))
		} else if isWebsite(str) {
			go bot.Message("^ " + webTitle("http://"+str))
		}
	}

	if strings.HasPrefix(message, "!help") {
		bot.Message("For help, please make the bot. dreaSMUG")
	} else if strings.HasPrefix(message, "!quote") {
		bot.Message(bot.getQuote())
	} else if strings.HasPrefix(message, "!addquote ") {
		stringpls := strings.Replace(message, "!addquote ", "", 1)
		if bot.isMod(username) {
			bot.quotes[stringpls] = username
			bot.writeQuoteDB()
			bot.Message("Quote added!")
		} else {
			bot.Message(username + " you are not a mod!")
		}
	} else if strings.HasPrefix(message, "!timeout ") {
		stringpls := strings.Replace(message, "!timeout ", "", 1)
		temp1 := strings.Split(stringpls, " ")
		temp2 := strings.Replace(stringpls, temp1[0], "", 1)
		if temp2 == "" {
			temp2 = "no reason"
		}
		if bot.isMod(username) {
			bot.timeout(temp1[0], temp2)
		} else {
			bot.Message(username + " you are not a mod!")
		}
	} else if strings.HasPrefix(message, "!ban ") {
		stringpls := strings.Replace(message, "!ban ", "", 1)
		temp1 := strings.Split(stringpls, " ")
		temp2 := strings.Replace(stringpls, temp1[0], "", 1)
		if temp2 == "" {
			temp2 = "no reason"
		}
		if bot.isMod(username) {
			bot.ban(temp1[0], temp2)
		} else {
			bot.Message(username + " you are not a mod!")
		}
	} else if strings.HasPrefix(message, "!uptime") {
		uptime := bot.getUptime(bot.channel)
		bot.Message(uptime)
	} else if message == "!time" {
		temp := time.Now()
		bot.Message(temp.String())
	} else if message == "!welcome" {
		bot.Message("Hi! Welcome to the stream, friend!")
	} else if message == "!whereami" {
		bot.Message("I am at " + bot.channel)
	}
}

// Begin website section 

func webTitle(website string) string {
	response, err := http.Get(website)
	if err != nil {
		return "Error reading website"
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "Error reading website"
		}
		if strings.Contains(string(contents), "<title>") && strings.Contains(string(contents), "</title>") {
			derp := strings.Split(string(contents), "<title>")
			derpz := strings.Split(derp[1], "</title>")
			return derpz[0]
		}
		return "No title"
	}
}

func isWebsite(website string) bool {
	domains := []string{".com", ".net", ".org", ".info", ".fm", ".gg", ".tv"}
	for _, domain := range domains {
		if strings.Contains(website, domain) {
			return true
		}
	}
	return false
}

// End website section 

// Begin Twitch API section

func (bot *Bot) getUptime(username string) string {

	// Passed channel as username
	channel := strings.TrimPrefix(username, "#")
	url := "https://api.twitch.tv/kraken/streams/" + channel
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err)

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var f interface{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	m := f.(map[string]interface{})

	// Better way to check for null?
	if stream, ok := m["stream"]; ok {
		streamMap := stream.(map[string]interface{})
		created_at := reflect.ValueOf(streamMap["created_at"]).String()
		fmt.Println("Created at: " + created_at)

		t, err := time.Parse("2006-01-02T15:04:05Z07:00", created_at)
		fmt.Println(t, err)
		fmt.Println(time.Since(t))
		str := time.Since(t).String()

		// Check for hours, because they may not exist (yet)
		hours := ""
		if  strings.Contains(str, "h") {
			hours = str[0:strings.Index(str, "h")] 
		} else {
			hours = "0"
		}

		mins := str[strings.Index(str, "h")+1 : strings.Index(str, "m")]

		return "Stream has been up for " + hours + " hours and " + mins + " mins."
	} else {
		return "Stream is offline."
	}
}

// End Twitch API section

// Begin mod section

func (bot *Bot) isMod(username string) bool {
	temp := strings.Replace(bot.channel, "#", "", 1)
	if bot.mods[username] == true || temp == username { 
		return true
	}
	return false
}

// If user is mod, timeout user and display the given reasoning
func (bot *Bot) timeout(username string, reason string) {
	if bot.isMod(username) {
		return
	}
	fmt.Fprintf(bot.conn, "PRIVMSG "+bot.channel+" :/timeout "+username+"\r\n")
	bot.Message(username + " was timed out(" + reason + ")!")
}

// If user is mod, ban user and display the given reasoning 
func (bot *Bot) ban(username string, reason string) {
	if bot.isMod(username) {
		return
	}
	fmt.Fprintf(bot.conn, "PRIVMSG "+bot.channel+" :/ban "+username+"\r\n")
	bot.Message(username + " was banned(" + reason + ")!")
}

// End mod section 

