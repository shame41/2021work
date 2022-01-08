import javax.swing.*;



public class game extends JFrame{


    public void init()
    {
        setSize(1500, 1000);
        setLocationRelativeTo(null);
        setDefaultCloseOperation(WindowConstants.EXIT_ON_CLOSE);
        theMainPanel layeredPane = new theMainPanel();
        layeredPane.init();
        setContentPane(layeredPane);
        setVisible(true);
        
    }





    public static void main(String[] args) {
        game game = new game();
        game.init();
    }
}
