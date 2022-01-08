import java.awt.*;
import javax.swing.*;
//由王墨然实现
/*
NormalZombie表示普通僵尸
生命值为1000
*/
public class NormalZombie extends Zombie {

    private Image normalZombieIMG;
    public NormalZombie(Garden parent, int lane) {
        super(parent, lane);
        normalZombieIMG = new ImageIcon(this.getClass().getResource("images/zombies/zombie1.png")).getImage();
    }

    public Image getImage()
    {
        return normalZombieIMG;
    }

}
