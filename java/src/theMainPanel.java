import javax.swing.*;
import java.awt.*;
import java.awt.event.*;

public class theMainPanel extends JLayeredPane
{
    private JPanel garden;

    private JPanel card_cherrybomb;
    private JPanel card_freezepeashooter;
    private JPanel card_peashooter;
    private JPanel card_sunflower;
    private JPanel card_wallnut;

    private Acre[][] acre;

    private Zombie[][] zombies;

    private Timer zombieMover;
    private Timer redrawTimer;

    private Image image;
    public theMainPanel()
    {  
        image = new ImageIcon(
                        this.getClass()
                        .getResource("images/cards/card_cherrybomb.png")).getImage();   
        zombieMover = new Timer(50, (ActionEvent e) -> {
            zombies[2][2].move();
        });
        redrawTimer = new Timer(25, (ActionEvent e) -> {
            repaint();
        });
        zombieMover.start();
        redrawTimer.start();
        setBounds(0, 0, 1500, 1000);
        setLayout(null);
 
        acre = new Acre[5][9];
        zombies = new Zombie[5][10];

        garden = new JPanel();
        card_cherrybomb = new JPanel();
        card_freezepeashooter = new JPanel();
        card_peashooter = new JPanel();
        card_sunflower = new JPanel();
        card_wallnut = new JPanel();

        for (Acre[] line : acre) 
            for(int i = 0; i < 9; i++)
                line[i] = new Acre();

        for (Zombie[] line : zombies) 
        {
            for (int i = 0; i < line.length; i++) 
            {
                int randY = (int)(Math.random()*5 + 1)*120;
                int randX = 950;
                int velocity = 1;
                line[i] = new Zombie(randX, randY, velocity, 
                    new ImageIcon("images/zombies/zombie1.png")); 
            }
        }
        // for (Acre[] line : acre) 
        //     for (Acre acre : line)
        //         acre.setOpaque(false);

    }
    private void initCardList()
    {
        card_sunflower.setBounds(110,10,64,95);
        card_peashooter.setBounds(174, 10, 64, 95);
        card_wallnut.setBounds(238, 10, 64, 95);
        card_cherrybomb.setBounds(302, 10, 64, 95);
        card_freezepeashooter.setBounds(366, 10, 64, 95);

        card_sunflower.add(
            new JLabel(
                new ImageIcon(
                    this.getClass()
                    .getResource("images/cards/card_sunflower.png"))));
        card_peashooter.add(
            new JLabel(
                new ImageIcon(
                    this.getClass()
                    .getResource("images/cards/card_peashooter.png"))));
        card_wallnut.add(
            new JLabel(
                new ImageIcon(
                    this.getClass()
                    .getResource("images/cards/card_wallnut.png"))));
        card_cherrybomb.add(
            new JLabel(
                new ImageIcon(
                    this.getClass()
                    .getResource("images/cards/card_cherrybomb.png"))));
        card_freezepeashooter.add(
            new JLabel(
                new ImageIcon(
                    this.getClass()
                    .getResource("images/cards/card_freezepeashooter.png"))));

        add(card_sunflower, 0);
        add(card_peashooter, 0);
        add(card_wallnut, 0);
        add(card_cherrybomb, 0);
        add(card_freezepeashooter, 0);

        
        card_sunflower.addMouseListener(
            new selectCard(
                new Sunflower(card_sunflower)));
        card_peashooter.addMouseListener(
            new selectCard(
                new Peashooter(card_peashooter)));
        card_cherrybomb.addMouseListener(
            new selectCard(
                new Cherrybomb(card_cherrybomb)));
        card_wallnut.addMouseListener(
            new selectCard(
                new Wallnut(card_wallnut)));
        card_freezepeashooter.addMouseListener(
            new selectCard(
                new Freezepeashooter(card_freezepeashooter)));
    }
    private void initGarden()
    {
        garden.setBounds(0, 0, 1000, 755);
        for (int j = 0; j < 5; j++) {
            for(int i = 0; i < 9; i++)
                acre[j][i].setBounds(40 + i*100, 120 + 120*j, 110, 120);

        }
       
        garden.add(
            new JLabel(
                new ImageIcon(
                    this.getClass()
                    .getResource("images/mainBG.png"))));
        // for (Acre[] line : acre) {
        //     for(int i = 0; i < 9; i++)
        //     line[i].add(
        //         new JLabel(
        //             new ImageIcon(
        //                 this.getClass()
        //                 .getResource("images/cards/card_cherrybomb.png"))));   
        // }        
        
        add(garden, 1);
        for (Acre[] line : acre) {
            for(int i = 0; i < 9; i++)
                add(line[i], 0);
        }

        for (Acre[] line : acre) {
            for(int i = 0; i < 9; i++)
                line[i].addMouseListener(new selectAcre(line[i]));
        }
    }
    private void initZombie()
    {

        for (Zombie[] line : zombies) 
        {
            for (Zombie zombie : line) 
            {   
                zombie.setBounds(zombie.getX(), zombie.getY(), 60, 90);
            }
        }
        for (Zombie[] line : zombies) 
        {
            for (Zombie zombie : line) 
            {   
                add(zombie, 3);
            }
        }
        
    }

    @Override
    protected void paintComponent(Graphics g)
    {
        super.paintComponent(g);
        g.drawImage(image, 200, 200, null);
    }
    public void init()
    {
        initGarden();
        initCardList();
        initZombie();
        zombieMover.start();
        redrawTimer.start();
    }

    

}
