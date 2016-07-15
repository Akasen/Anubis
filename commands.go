// Commands Go file
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
		bot.Message("For help, please make the bot.")
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
	} else if message == "!time" {
		temp := time.Now()
		bot.Message(temp.String())
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
