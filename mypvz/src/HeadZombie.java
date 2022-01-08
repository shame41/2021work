import java.awt.*;
import javax.swing.*;
//由王墨然实现
/*
HeadZombie表示高级僵尸
生命值为1800
*/
public class HeadZombie extends Zombie {

    private Image HeadZombieIMG;
    public HeadZombie(Garden parent, int lane) 
    {
        super(parent, lane);
        setHealth(3000);
        HeadZombieIMG = new ImageIcon(this.getClass().getResource("images/zombies/zombie2.png")).getImage();
    }
    
    public Image getImage()
    {
        return HeadZombieIMG;
    }
}
