/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.mycompany.anubis;

import org.kitteh.irc.client.library.event.channel.ChannelJoinEvent;
import org.kitteh.irc.client.library.event.channel.ChannelMessageEvent;
import org.kitteh.irc.lib.net.engio.mbassy.listener.Handler;

/**
 *
 * @author Joshua
 */
public class Bot {
    
    
    public static class eventListener {
        
        //Greet users
        @Handler
        public void onUserJoinChannel(ChannelJoinEvent event) {
            if (event.getClient().isUser(event.getUser())) {
                event.getChannel().sendMessage("This bot is active!");
                return;
            }
            event.getChannel().sendMessage("Welcome, " + event.getUser().getNick() + "! :3");
        }
        
        //Retreive chat messages
        @Handler
        public void getChat(ChannelMessageEvent event) {
            if (event.getMessage().isEmpty() == false) {
                //gui.chatText.append(event.getActor().getNick() + ": " + event.getMessage() + "\r\n");
            }
        }
        
        @Handler
        public void replyGreetings(ChannelMessageEvent event) {
            if(event.getMessage().contentEquals("Hello bot"))
            {
                
                event.getChannel().sendMessage("Hello there!");
            }
        }
        
        @Handler
        public void stopRiot(ChannelMessageEvent event) {
            if(event.getMessage().equalsIgnoreCase("RIOT")) {
                event.getChannel().sendMessage("STOP THE RIOTING");
            }
        }
            
        }
    
}
