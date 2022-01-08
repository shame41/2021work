import javax.swing.*;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.event.MouseEvent;
import java.awt.event.MouseMotionListener;
import java.util.ArrayList;
import java.util.Random;


//由何洛昌实现
/*
Garden表示整个草坪，包括上面的植物和僵尸
是一个实现了MouseMotionListener的JLayeredPane
成员变量有
    bgImage /Image 是各个元素的图片
    arces/Arce[] 是草坪上地块的集合
    zombies/ArrayList<ArrayList<Zombie>> 是草坪上僵尸的集合
    peas/ArrayList<ArrayList<Pea>> 是草坪上子弹的集合
    suns/ArrayList<Sun> 是掉落和生成的阳光的集合
    redrawTimer moveTimer sunProducer zombieProducer/Timer 是各种元素生成器的计时器
    sunScoreboard/JLable 是阳光的数量
    condition/Game.ConditionOfMouse 是鼠标上的植物类型
方法有
    move() 用于控制僵尸与子弹的移动
    paintComponent(Graphics g) 用于绘制JLayeredPane中的各个组件
    还有各种构造函数/getter setter
内部类
    PlantActionListener 用来充当ActionListener，表示植物的种植
*/
public class Garden extends JLayeredPane implements MouseMotionListener {

    private Image bgImage;
    
    
    
    // private Image peaIMG;
    // private Image freezePeaIMG;

    // private Image normalZombieIMG;
    // private Image headZombieIMG;
    private Arce[] arces;

    private ArrayList<ArrayList<Zombie>> zombies;
    private ArrayList<ArrayList<Pea>> peas;
    private ArrayList<Sun> suns;

    private Timer redrawTimer;
    private Timer moveTimer;
    private Timer sunProducer;
    private Timer zombieProducer;

    private JLabel sunScoreboard;

    private Game.ConditionOfMouse condition = Game.ConditionOfMouse.None;

    private int mouseX, mouseY;

    private int sunScore;

    public int getSunScore() {
        return sunScore;
    }

    public void setSunScore(int sunScore) {
        this.sunScore = sunScore;
        sunScoreboard.setText(String.valueOf(sunScore));
    }

    public Garden(JLabel sunScoreboard) 
    {
        setSize(1000, 752);
        setLayout(null);
        addMouseMotionListener(this);
        this.sunScoreboard = sunScoreboard;
        setSunScore(150);  //pool avalie

        initImg();
        initZombies();
        initPeas();
        initArces();
        initTimer();
        
        suns = new ArrayList<>();
        
    }


    private void move() //控制阳光、子弹与僵尸的移动
    {
        for (int i = 0; i < 5; i++) 
        {
            for (Zombie z : zombies.get(i)) 
                z.move();

            for (int j = 0; j < peas.get(i).size(); j++) 
            {
                Pea p = peas.get(i).get(j);
                p.move();
            }

        }

        for (int i = 0; i < suns.size(); i++) 
            suns.get(i).move();

    }

    @Override
    protected void paintComponent(Graphics g) 
    {
        super.paintComponent(g);
        g.drawImage(bgImage, 0, 0, null);

        for (int i = 0; i < 45; i++) //画植物
        {
            Arce arce = arces[i];
            if (arce.plant != null) 
            {
                Plant p = arce.plant;
                g.drawImage(p.getImage(), 60 + (i % 9) * 100, 129 + (i / 9) * 120, null);
            }
        }

        for (int i = 0; i < 5; i++) 
        {
            for (Zombie zombie : zombies.get(i)) //画僵尸
                g.drawImage(zombie.getImage(), zombie.getX(),109 + (i * 120), null);

            for (int j = 0; j < peas.get(i).size(); j++) //画子弹
            {
                Pea pea = peas.get(i).get(j);
                g.drawImage(pea.getImage(), pea.getX(), 130 + (i * 120), null);
            }

        }


    }

    private class PlantActionListener implements ActionListener {//用来初始化Arce

        int x, y;

        public PlantActionListener(int x, int y) {
            this.x = x;
            this.y = y;
        }

        @Override
        public void actionPerformed(ActionEvent e) 
        {
            if (condition == Game.ConditionOfMouse.Sunflower)
                if (getSunScore() >= 50) 
                {
                    arces[x + y * 9].setPlant(new Sunflower(Garden.this, x, y));
                    setSunScore(getSunScore() - 50);
                }
            if (condition == Game.ConditionOfMouse.Peashooter) 
                if (getSunScore() >= 100) 
                {
                    arces[x + y * 9].setPlant(new Peashooter(Garden.this, x, y));
                    setSunScore(getSunScore() - 100);
                }

            if (condition == Game.ConditionOfMouse.FreezePeashooter) 
                if (getSunScore() >= 175) {
                    arces[x + y * 9].setPlant(new FreezePeashooter(Garden.this, x, y));
                    setSunScore(getSunScore() - 175);
                }
            condition = Game.ConditionOfMouse.None;
        }
    }
    private void initPeas()
    {
        peas = new ArrayList<>();
        peas.add(new ArrayList<>()); //line 1
        peas.add(new ArrayList<>()); //line 2
        peas.add(new ArrayList<>()); //line 3
        peas.add(new ArrayList<>()); //line 4
        peas.add(new ArrayList<>()); //line 5
    }
    private void initTimer()
    {
        redrawTimer = new Timer(25, (ActionEvent e) -> {
            repaint();
        });
        redrawTimer.start();

        moveTimer = new Timer(60, (ActionEvent e) -> move());
        moveTimer.start();

        sunProducer = new Timer(5000, (ActionEvent e) -> {
            Random rnd = new Random();
            Sun sta = new Sun(this, rnd.nextInt(800) + 100, 0, rnd.nextInt(300) + 200);
            suns.add(sta);
            add(sta, 1);
        });
        sunProducer.start();

        zombieProducer = new Timer(7000, (ActionEvent e) -> {
            Random rnd = new Random();

            int l = rnd.nextInt(5);//行数
            int t = rnd.nextInt(100);//随机权重
            Zombie z = null;
            for (int i = 0; i < 50; i++) 
            {
                if (t < 25) 
                    z = new HeadZombie(this, l);
                else
                    z = new NormalZombie(this, l);
            }
            zombies.get(l).add(z);
        });
        zombieProducer.start();

    }

    private void initZombies()
    {
        zombies = new ArrayList<>();
        zombies.add(new ArrayList<>()); //line 1
        zombies.add(new ArrayList<>()); //line 2
        zombies.add(new ArrayList<>()); //line 3
        zombies.add(new ArrayList<>()); //line 4
        zombies.add(new ArrayList<>()); //line 5
    }

    private void initArces()
    {
        arces = new Arce[45];
        for (int i = 0; i < 45; i++) 
        {
            Arce a = new Arce();
            a.setLocation(44 + (i % 9) * 100, 109 + (i / 9) * 120);
            a.setAction(new PlantActionListener((i % 9), (i / 9)));
            arces[i] = a;
            add(a, 0);
        }

    }

    private void initImg()
    {
        bgImage = new ImageIcon(
            this.getClass()
            .getResource("images/mainBG.png"))
            .getImage();
    }

    @Override
    public void mouseDragged(MouseEvent e) {

    }

    @Override
    public void mouseMoved(MouseEvent e) {
        mouseX = e.getX();
        mouseY = e.getY();
    }

    static int progress = 0;


    public Game.ConditionOfMouse getCondition() {
        return condition;
    }

    public void setCondition(Game.ConditionOfMouse condition) {
        this.condition = condition;
    }

    public ArrayList<ArrayList<Zombie>> getZombies() {
        return zombies;
    }

    public void setZombies(ArrayList<ArrayList<Zombie>> zombies) {
        this.zombies = zombies;
    }

    public ArrayList<ArrayList<Pea>> getPeas() {
        return peas;
    }

    public void setPeas(ArrayList<ArrayList<Pea>> peas) {
        this.peas = peas;
    }

    public ArrayList<Sun> getSuns() {
        return suns;
    }

    public void setSuns(ArrayList<Sun> suns) {
        this.suns = suns;
    }

    public Arce[] getArces() {
        return arces;
    }

    public void setArces(Arce[] arces) {
        this.arces = arces;
    }
}
