import javax.swing.ImageIcon;
import javax.swing.JPanel;
import java.awt.*;

public class Acre extends JPanel{
    private Plant plant = new NonePlant();
    
    public void setPlant(Plant plant)
    {
        this.plant = plant;
    }

    public Plant getPlant()
    {
        return plant;
    }

    public void paintComponent(Graphics g)
    {   
        Image image = new ImageIcon(
                this.getClass()
                .getResource("images/cards/card_cherrybomb.png")).getImage();   
                super.paintComponent(g);
        g.drawImage(image, 200, 200, null);
    }
}
