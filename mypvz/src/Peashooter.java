import javax.swing.*;
import java.awt.*;
import java.awt.event.ActionEvent;

//由张跃然实现
/*
Peashooter是豌豆射手
是Plant的一个子类
它的成员变量有
    shootTimer/Timer 子弹发射计时器
方法有
    stop() 表示植物的死亡
    还有构造函数
*/
public class Peashooter extends Plant {

    public Timer shootTimer;
    private Image peashooterIMG;

    public Peashooter(Garden garden, int x, int y) 
    {
        super(garden, x, y);
        peashooterIMG = new ImageIcon(this.getClass().getResource("images/plants/peashooter.gif")).getImage();
        shootTimer = new Timer(2000, (ActionEvent e) -> {
            //System.out.println("SHOOT");
            if (getGarden().getZombies().get(y).size() > 0) {
                getGarden().getPeas().get(y).add(new Pea(getGarden(), y, 103 + this.getX() * 100));
            }
        });
        shootTimer.start();
    }

    public Image getImage()
    {
        return peashooterIMG;
    }

    @Override
    public void stop() {
        shootTimer.stop();
    }

}
