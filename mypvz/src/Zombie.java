import javax.swing.*;
import java.awt.*;

//由王墨然实现
/*
Zombie表示草坪上的僵尸
成员变量有
    isMoving/boolean 表示僵尸是否在移动
    health/int 表示僵尸的生命值
    velocity/int 表示僵尸的移动速度
    garden/Garden 表示草坪
    x/int 表示僵尸所处的横坐标
    myLane/int 表示僵尸所在的行
    slowInt/int 用于减速的一个参数
方法有
    move() 用于让僵尸移动
    slow() 用于让僵尸减速
    还有各种构造函数/getter setter
*/
public abstract class Zombie {
    private boolean isMoving = true;
    private int health = 1000;
    private int velocity = 2;

    private Garden garden;

    private int x = 1000;
    private int myLane;//僵尸所处的行
  

    public Zombie(Garden garden, int lane) {
        this.garden = garden;
        myLane = lane;
    }
    int slowInt = 0;
    public void move() //僵尸移动的操作
    {
        if (isMoving) 
        {
            boolean isInEmptyArce = false;
            Arce arce = null;
            for (int i = myLane * 9; i < (myLane + 1) * 9; i++) 
            {
                if (garden.getArces()[i].plant != null 
                    && garden.getArces()[i].isInArce(x))
                {
                    isInEmptyArce = true;
                    arce = garden.getArces()[i];
                }
            }
            if (!isInEmptyArce) 
            {
                if (slowInt > 0) 
                {
                    if (slowInt % 2 == 0) 
                        x--;
                    slowInt--;
                } else 
                    x -= velocity;
            } 
            else 
            {
                arce.plant.setHealth(arce.plant.getHealth() - 10);
                if (arce.plant.getHealth() < 0) 
                    arce.removePlant();
            }
            if (x < 0) 
            {
                isMoving = false;
                JOptionPane.showMessageDialog(garden, "ZOMBIES ATE YOUR BRAIN !" + '\n' + "Starting the level again");
                Game.game.dispose();
                Game.game = new Game();
            }
        }
    }



    public void slow() {
        slowInt = 1000;
    }



    public Image getImage()
    {
        return null;
    }

    public int getHealth() {
        return health;
    }

    public void setHealth(int health) {
        this.health = health;
    }

    public int getVelocity() {
        return velocity;
    }

    public void setVelocity(int velocity) {
        this.velocity = velocity;
    }

    public Garden getGarden() {
        return garden;
    }

    public void setGarden(Garden garden) {
        this.garden = garden;
    }

    public int getX() {
        return x;
    }

    public void setX(int x) {
        this.x = x;
    }

    public int getMyLane() {
        return myLane;
    }

    public void setMyLane(int myLane) {
        this.myLane = myLane;
    }

    public boolean isMoving() {
        return isMoving;
    }

    public void setMoving(boolean moving) {
        isMoving = moving;
    }

    public int getSlowInt() {
        return slowInt;
    }

    public void setSlowInt(int slowInt) {
        this.slowInt = slowInt;
    }
}
