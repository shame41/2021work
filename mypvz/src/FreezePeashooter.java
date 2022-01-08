import javax.swing.*;
import java.awt.*;
import java.awt.event.ActionEvent;

//由张跃然实现
/*
FreezePeashooter是寒冰射手，射出的子弹能让僵尸减速
是Plant的一个子类
它的成员变量有
    shootTimer/Timer 子弹发射计时器
方法有
    stop() 表示植物的死亡
    还有构造函数
*/
public class FreezePeashooter extends Plant 
{

    private Timer shootTimer;
    private Image freezePeashooterIMG;

    public FreezePeashooter(Garden garden, int x, int y) 
    {
        super(garden, x, y);
        freezePeashooterIMG = new ImageIcon(this.getClass().getResource("images/plants/freezepeashooter.gif")).getImage();
        shootTimer = new Timer(2000, (ActionEvent e) -> 
        {

            if (getGarden().getZombies().get(y).size() > 0) 
            {
                getGarden().getPeas().get(y)
                .add(
                    new FreezePea(getGarden(), 
                    y, 
                    103 + this.getX() * 100));
            }
        });
        shootTimer.start();
    }

    public Image getImage()
    {
        return freezePeashooterIMG;
    }

    @Override
    public void stop() 
    {
        shootTimer.stop();
    }

}