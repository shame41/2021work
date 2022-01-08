
import java.awt.event.*;

import javax.swing.ImageIcon;
import javax.swing.JLabel;

public class selectAcre extends MouseAdapter {
    private Acre acre;
    public selectAcre(Acre acre)
    {   super();
        this.acre = acre;
    }

    @Override
    public void mouseClicked(MouseEvent e) {
        super.mouseClicked(e);
        ConditionOfMouse condition = new ConditionOfMouse();
        if (acre.getPlant().test().equals(new NonePlant().test())) {
            acre.setPlant(condition.getCondition());
            System.out.println(
                "X IS " + acre.getX() +
                " AND Y IS " + acre.getY() + 
                " PLANT IS " + acre.getPlant().test());
            condition.setPlant(new NonePlant());
            acre.add(
                new JLabel(
                    new ImageIcon(
                        this.getClass()
                        .getResource("images/cards/card_peashooter.png"))));
        }
        else
        {
            if (condition.getCondition().test().equals(new NonePlant().test())) {
                System.out.println(
                    "X IS " + acre.getX() +
                    " AND Y IS " + acre.getY() + 
                    " PLANT IS " + acre.getPlant().test());
            }
            else
            {
                System.out.println("this acre has been occupied");
            }
        }

    }
    
}
