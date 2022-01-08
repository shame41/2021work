import java.awt.*;
import javax.swing.*;

//由王墨然实现
/*
FreezePea是寒冰子弹，能让僵尸减速
是Pea的一个子类
方法有
    move() 表示子弹的移动与判断是否撞上僵尸
    还有构造函数
*/
public class FreezePea extends Pea {

    private Image freezePeaIMG;

    public FreezePea(Garden parent, int lane, int startX) {
        super(parent, lane, startX);
        freezePeaIMG = new ImageIcon(this.getClass().getResource("images/freezepea.png")).getImage();

    }

    public Image getImage()
    {
        return freezePeaIMG;
    }
    @Override
    public void move() 
    {
        Rectangle pRect = new Rectangle(getX(), 130 + getMyLane() * 120, 28, 28);
        for (int i = 0; i < garden.getZombies().get(getMyLane()).size(); i++) 
        {
            Zombie z = garden.getZombies().get(getMyLane()).get(i);
            Rectangle zRect = new Rectangle(z.getX(), 109 + getMyLane() * 120, 400, 120);
            if (pRect.intersects(zRect)) 
            {
                z.setHealth(z.getHealth() - 300);
                z.slow();
                boolean exit = false;
                if (z.getHealth() < 0) 
                {
                    System.out.println("ZOMBIE DIE");
                    // Garden.setProgress(10);
                    garden.getZombies().get(getMyLane()).remove(i);
                    exit = true;
                }
                garden.getPeas().get(getMyLane()).remove(this);
                if (exit) 
                    break;
            }
        }

        setX(getX() + 15);
    }

}
