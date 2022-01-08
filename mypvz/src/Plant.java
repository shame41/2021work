import java.awt.*;
//由何洛昌实现
/*
Plant表示植物
是一个抽象类
成员变量有
    health/int 表示植物的生命值
    x y/int 表示植物的横坐标纵坐标
    garden/Garden 表示草坪
方法有
    stop() 用来终结当前植物的生命
    各种构造函数/getter setter
*/
public abstract class Plant {

    private int health = 200;

    private int x;
    private int y;

    private Garden garden;


    public Plant(Garden garden, int x, int y) {
        this.x = x;
        this.y = y;
        this.garden = garden;
    }

    public Image getImage()
    {
        return null;
    }

    public void stop() {
    }

    public int getHealth() {
        return health;
    }

    public void setHealth(int health) {
        this.health = health;
    }

    public int getX() {
        return x;
    }

    public void setX(int x) {
        this.x = x;
    }

    public int getY() {
        return y;
    }

    public void setY(int y) {
        this.y = y;
    }

    public Garden getGarden() {
        return garden;
    }

    public void setGarden(Garden garden) {
        this.garden = garden;
    }
}
