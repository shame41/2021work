public class ConditionOfMouse {

    private static Plant condition = new NonePlant();

    public void setPlant(Plant p) 
    {
        condition = p;
    }

    public Plant getCondition()
    {
        return condition;
    }

}
