/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.mycompany.anubis;

import java.util.function.Consumer;
import org.kitteh.irc.client.library.Client;
import org.kitteh.irc.client.library.event.channel.ChannelJoinEvent;
import org.kitteh.irc.client.library.event.channel.ChannelMessageEvent;
import org.kitteh.irc.lib.net.engio.mbassy.listener.Handler;
import org.kitteh.irc.client.library.event.helper.*;

/**
 *
 * @author Joshua
 */
public class Anubis {
    
    public static void main(String[] args) {
       
        
        Anubis_GUI gui = new Anubis_GUI();
        IRC_Client Anubis = new IRC_Client();
        AnubisController controller = new AnubisController(gui, Anubis);
        
        //Anubis.initializeIRC();
        
        Anubis.connect("#akasen1226");
        
        //while(true)
        controller.updateView();
        
    }
}