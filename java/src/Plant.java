import javax.swing.JPanel;

import java.awt.*;
import java.awt.Image;


public class Plant extends JPanel{
    private JPanel card;
    public Plant(JPanel card)
    {
        this.card = card;
    }
    public String test()
    {
        return "plant";
    }

    @Override
    protected void paintComponent(Graphics g)
    {
        super.paintComponent(g);


    }

}
