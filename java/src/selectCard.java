
import java.awt.event.*;

import javax.swing.JPanel;


public class selectCard extends MouseAdapter {
    Plant plant;
    public selectCard(Plant jp)
    {   super();
        this.plant = jp;
    }

    @Override
    public void mouseClicked(MouseEvent e) {
        super.mouseClicked(e);
        ConditionOfMouse condition = new ConditionOfMouse();
        if (condition.getCondition().test().equals(new NonePlant().test())) {
            condition.setPlant(plant);
        }
        else
        {
            condition.setPlant(new NonePlant());
        }
        // new ConditionOfMouse().setPlant(plant);
        // System.out.println(plant.getX());
    }


}
