/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.mycompany.anubis;

/**
 *
 * @author Joshua
 */

import java.util.function.Consumer;
import org.kitteh.irc.client.library.Client;
import org.kitteh.irc.client.library.event.channel.ChannelMessageEvent;
import org.kitteh.irc.client.library.event.channel.ChannelJoinEvent;
import org.kitteh.irc.lib.net.engio.mbassy.listener.Handler;

public class IRC_Client {
    
    Client client;
    
    private String host, nick, password;
    private int port = 6667;
    private String output = "";
    
    public IRC_Client() {
        Consumer<Exception> warning = (y) -> System.out.println(y);
        //Consumer<String> output = (x) -> listenChat(x);
        Consumer<String> input = (x) -> listenChat(x);
        client = Client.builder()
            //.listenOutput(output)
            .listenInput(input)
            .listenException(warning)
            .serverHost("irc.twitch.tv")
            .serverPort(port)
            .serverPassword("oauth:[code]")
            .nick("name")
            .secure(false)
            .build();
        //client.getEventManager().registerEventListener(new eventListener());
    }
    
    void listenChat(String X) {
        //System.out.print(X + "\n");
        this.output = (X + "\n");
    }
    
    String getChat() {

        return output;
    }

    void setPort(int newPort) {
        port = newPort;
    }
    
    void setNick(String newNick) {
        nick = newNick;
    }
    
    
    void setPassword(String newPassword) {
        password = newPassword;
    }
    
    void setHost(String newHost) {
        host = newHost;
    }
    
    
    void connect(String channel) {
        this.client.addChannel(channel);
        
    }
    
    //Initialize the IRC_Client
    
    /*
    public void initializeIRC() {
        //Consumer<String> ear = (x) -> gui.errorMessage.append(x + "\r\n");
        Consumer<Exception> warning = (y) -> System.out.println(y);
        client = Client.builder()
            //.listenOutput(ear)
            //.listenInput(ear)
            .listenException(warning)
            .serverHost("irc.twitch.tv")
            .serverPort(6667)
            .serverPassword("oauth:[code]")
            .nick("name")
            .secure(false)
            .build();
        client.getEventManager().registerEventListener(new eventListener());
        
        //return client;
    }
    */
    
    
    public static class eventListener {
        
        public String line = "";
        
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
            if(event.getMessage().equalsIgnoreCase("Hello bot"))
            {
                event.getChannel().sendMessage("Hello there!");
            }
        }
        
        //Fun
        @Handler
        public void stopRiot(ChannelMessageEvent event) {
            if(event.getMessage().equalsIgnoreCase("RIOT")) {
                event.getChannel().sendMessage("STOP THE RIOTING");
            }
        }
            
        }
    
}


