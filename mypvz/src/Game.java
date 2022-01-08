import javax.swing.*;
import java.awt.event.ActionEvent;

//由何洛昌实现
/*
游戏的主界面/main函数所处的位置
成员变量有
    sun/JLabel 表示阳光值的标签
    garden/Garden 表示整个草坪
    sunflower peashooter freezepeashooter/PlantCard 表示可选择的植物卡片
    ConditionOfMouse/enum 表示鼠标选到了哪个植物
方法有
    initGarden initWindows initSun initPlantCards/四个init构造器
*/
public class Game extends JFrame {

    private JLabel sun;
    
    private Garden garden;

    private PlantCard sunflower;
    private PlantCard peashooter;
    private PlantCard freezepeashooter;
    

    enum ConditionOfMouse {
        None,
        Sunflower,
        Peashooter,
        FreezePeashooter
    }
    static Game game;

    public Game() 
    {        
        setSize(1012, 785);
        setDefaultCloseOperation(WindowConstants.EXIT_ON_CLOSE);
        setLayout(null);

        // initWindows();
        initSun();
        initGarden();
        initPlantCards();

        getLayeredPane().add(sun, 0);
        setResizable(false);
        setVisible(true);
    }

    private void initGarden()
    {
        garden = new Garden(sun);
        garden.setLocation(0, 0);
        getLayeredPane().add(garden, 0);
    }



    private void initSun()
    {
        sun = new JLabel("SUN");
        sun.setLocation(37, 80);
        sun.setSize(60, 20);
        

    }

    private void initPlantCards()
    {
        sunflower = new PlantCard(
            new ImageIcon(
                this.getClass()
                .getResource("images/cards/card_sunflower.png"))
                .getImage());
        sunflower.setLocation(110, 8);
        sunflower.setAction((ActionEvent e) -> {
            garden.setCondition(Game.ConditionOfMouse.Sunflower);
        });
        getLayeredPane().add(sunflower, 0);//放到最顶层
        
        peashooter = new PlantCard(
            new ImageIcon(
                this.getClass()
                .getResource("images/cards/card_peashooter.png"))
                .getImage());
        peashooter.setLocation(175, 8);
        peashooter.setAction((ActionEvent e) -> {
            garden.setCondition(Game.ConditionOfMouse.Peashooter);
        });
        getLayeredPane().add(peashooter, 0);
        
        freezepeashooter = new PlantCard(
            new ImageIcon(this.getClass()
            .getResource("images/cards/card_freezepeashooter.png"))
            .getImage());
        freezepeashooter.setLocation(240, 8);
        freezepeashooter.setAction((ActionEvent e) -> {
            garden.setCondition(Game.ConditionOfMouse.FreezePeashooter);
        });
        getLayeredPane().add(freezepeashooter, 0);

    }
    
    

    public static void main(String[] args) {
        // game.dispose();
        game = new Game();
    }

}
