import javax.swing.*;
import java.awt.*;
import java.awt.event.MouseEvent;
import java.awt.event.MouseListener;

//由张跃然实现
/*
Sun表示阳光
是一个实现了MouseListener的JPanel
成员变量有
    myX myY/int 表示阳光的初始坐标
    endY/int 表示阳光最后到达的纵坐标 
    garden/Garden 表示草坪
    sunImage/Image 表示阳光的图片
    survivalTime/int 表示阳光落地后的存活时间
方法有
    paintComponent(Graphics g) 用来绘制组件
    move() 使阳光移动
*/
public class Sun extends JPanel implements MouseListener {

    private Garden garden;
    private Image sunImage;

    private int myX;
    private int myY;
    private int endY;

    private int survivalTime = 200;

    public Sun(Garden garden, int startX, int startY, int endY) {
        this.garden = garden;
        this.endY = endY;
        setSize(80, 80);
        setOpaque(false);
        myX = startX;
        myY = startY;
        setLocation(myX, myY);
        sunImage = new ImageIcon(this.getClass().getResource("images/sun.png")).getImage();
        addMouseListener(this);
    }

    @Override
    protected void paintComponent(Graphics g) {
        super.paintComponent(g);
        g.drawImage(sunImage, 0, 0, null);
    }

    public void move() 
    {
        if (myY < endY) 
            myY += 4;
        else 
        {
            survivalTime--;
            if (survivalTime < 0) 
            {
                garden.remove(this);
                garden.getSuns().remove(this);
            }
        }
        setLocation(myX, myY);
    }

    @Override
    public void mouseClicked(MouseEvent e) {

    }

    @Override
    public void mousePressed(MouseEvent e) {

    }

    @Override
    public void mouseReleased(MouseEvent e) {
        garden.setSunScore(garden.getSunScore() + 25);
        garden.remove(this);
        garden.getSuns().remove(this);
    }

    @Override
    public void mouseEntered(MouseEvent e) {

    }

    @Override
    public void mouseExited(MouseEvent e) {

    }
}
