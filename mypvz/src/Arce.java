import javax.swing.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.event.MouseEvent;
import java.awt.event.MouseListener;

//由张跃然实现
/*
Arce表示植物的地块，用以种植植物
是一个实现了MouseListener的JPanel
成员变量有 
    al/ActionListener 用于监听地块是否被点击
    plant/Plant 表示这个地块上所种植的植物
方法有
    isInArce(int tx) 用于表示某物是否处在这个地块上
    mouseReleased(MouseEvent e) 用于触发事件
    还有各类getter/setter/构造器
*/

public class Arce extends JPanel implements MouseListener {

    private ActionListener al;

    public Arce() {
        setOpaque(false);
        addMouseListener(this);
        setSize(100, 120);
    }

    public Plant plant;

    public void setPlant(Plant p) {
        plant = p;
    }

    public void removePlant() {
        plant.stop();
        plant = null;
    }

    public boolean isInArce(int tx) {
        return (tx > getLocation().x) && (tx < getLocation().x + 100);
    }

    public void setAction(ActionListener al) {
        this.al = al;
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
