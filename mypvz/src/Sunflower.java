import javax.swing.*;
import java.awt.*;
import java.awt.event.ActionEvent;

//由张跃然实现
/*
Sunflower表示向日葵
是Plant的一个子类
成员变量有
    sunProduceTimer/Timer 用来生产阳光的计时器
方法有
    构造函数
*/
public class Sunflower extends Plant {

    private Timer sunProduceTimer;
    private Image sunflowerIMG;

    public Sunflower(Garden garden, int x, int y) {
        super(garden, x, y);
        sunflowerIMG = new ImageIcon(this.getClass().getResource("images/plants/sunflower.gif")).getImage();
        sunProduceTimer = new Timer(15000, (ActionEvent e) -> {
            Sun sun = new Sun(getGarden(), 60 + x * 100, 110 + y * 120, 130 + y * 120);
            getGarden().getSuns().add(sun);
            getGarden().add(sun, 0);
        });
        sunProduceTimer.start();
    }

    public Image getImage()
    {
        return sunflowerIMG;
    }

}
