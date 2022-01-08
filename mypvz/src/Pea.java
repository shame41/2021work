import java.awt.*;
import javax.swing.*;

//由王墨然实现
/*
Pea是子弹
成员变量有
    x/int 表示子弹的横坐标
    garden/Garden 草坪
    myLane/int 表示子弹所处在的行
方法有
    move() 表示子弹的移动与判断是否撞上僵尸
    还有构造函数/getter setter
*/
public class Pea {

    private int x;
    protected Garden garden;
    private int myLane;
    private Image peaIMG;

    public Pea(Garden garden, int lane, int startX) {
        this.garden = garden;
        this.myLane = lane;
        x = startX;
        peaIMG = new ImageIcon(this.getClass().getResource("images/pea.png")).getImage();
    }

    public void move() 
    {
        Rectangle pRect = new Rectangle(x, 130 + myLane * 120, 28, 28);//子弹所处的范围
        for (int i = 0; i < garden.getZombies().get(myLane).size(); i++) 
        {
            Zombie zombie = garden.getZombies().get(myLane).get(i);
            Rectangle zRect = new Rectangle(zombie.getX(), 109 + myLane * 120, 400, 120);//僵尸所处的范围
            if (pRect.intersects(zRect)) 
            {//如果两范围相交
                zombie.setHealth(zombie.getHealth() - 300);
                boolean exit = false;
                System.out.println(zombie.getHealth());
                if (zombie.getHealth() < 0) {
                    System.out.println("ZOMBIE DIED");

                    garden.getZombies().get(myLane).remove(i);
                    // Garden.setProgress(10);
                    exit = true;
                }
                garden.getPeas().get(myLane).remove(this);//子弹消失
                if (exit) 
                    break;
            }
        }

        x += 15;//子弹向右移
    }

    public Image getImage()
    {
        return peaIMG;
    }

    public int getX() {
        return x;
    }

    public void setX(int posX) {
        this.x = posX;
    }

    public int getMyLane() {
        return myLane;
    }

    public void setMyLane(int myLane) {
        this.myLane = myLane;
    }
}
