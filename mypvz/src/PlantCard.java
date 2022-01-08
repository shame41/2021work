import javax.swing.*;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.event.MouseEvent;
import java.awt.event.MouseListener;

//由何洛昌实现
/*
PlantCard是可选择的植物卡片，阳光足够时，可以选择植物卡片来种植植物
是JPanel的子类，实现了MouseListener
成员变量有
    img/Image 表示卡片的图片
    al/ActionListener 用于监听卡片是否被点击
方法有
    paintComponent(Graphics g) 用来绘制组件
    还有构造函数/getter setter
*/
public class PlantCard extends JPanel implements MouseListener {

    private Image img;
    private ActionListener al;

    public PlantCard(Image img) {
        setSize(64, 90);
        this.img = img;
        addMouseListener(this);
    }

    public void setAction(ActionListener al) {
        this.al = al;
    }

    @Override
    protected void paintComponent(Graphics g) {
        super.paintComponent(g);
        g.drawImage(img, 0, 0, null);
    }

    @Override
    public void mouseClicked(MouseEvent e) {

    }

    @Override
    public void mousePressed(MouseEvent e) {

    }

    @Override
    public void mouseReleased(MouseEvent e) {
        if (al != null) {
            al.actionPerformed(new ActionEvent(this, ActionEvent.RESERVED_ID_MAX + 1, ""));
        }
    }

    @Override
    public void mouseEntered(MouseEvent e) {

    }

    @Override
    public void mouseExited(MouseEvent e) {

    }
}
