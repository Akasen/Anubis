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
public class AnubisController {
    private IRC_Client model;
    private Anubis_GUI view;
    
    public AnubisController(Anubis_GUI newView, IRC_Client newModel) {
        this.model = newModel;
        this.view = newView;
        
        view.setVisible(true);
        
        model.client.getEventManager().registerEventListener(new IRC_Client.eventListener());
    }
    
    public void updateView(){
        if(model.getChat() != null) {
            view.txtChat.append(model.getChat() + "\n");
            
        }
        
    }
}
