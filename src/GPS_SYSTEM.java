public class GPS_SYSTEM extends Application{
    private Boolean[] sensorChange;
    private Integer[][] cord;


    public void setSensorChange(int index, Boolean change ) {
        this.sensorChange[index] = change;
    }

    public void checkDisplacement(){
        if (sensorChange[0] == true || sensorChange[1] == true) {
            super.setNotification("Bike Has Experienced Displacement");
        }
    }

    public void setCord(int index, Integer[] cord){
        this.cord[index] = cord;
    }

    @Override
    public void setCoordinates(Integer[][] cord) {
        super.setCoordinates(cord);
    }
}
