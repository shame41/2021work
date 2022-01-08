import javax.swing.JPanel;
import java.awt.event.*;

public class plantCard extends JPanel{
    private ActionListener al;

    public void addActionListener(ActionListener al)
    {
        this.al = al;
    }
}
