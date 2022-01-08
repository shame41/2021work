import javax.swing.ImageIcon;
import javax.swing.JLabel;

public class Zombie extends JLabel{
    private int x;
    private int y;
    private int velocity;

    public Zombie(int x, int y, int velocity, ImageIcon imageIcon)
    {
        super(imageIcon);
        this.x = x;
        this.y = y;
        this.velocity = velocity;
    }

    public void move()
    {
        x = x - velocity;
    }

    public int getX()
    {

        return x;
    }

    public int getY()
    {

        return y;
    }
}
