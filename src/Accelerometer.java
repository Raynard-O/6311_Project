public class Accelerometer extends Microcontroller{
    private final Integer[] accelerometerInitValue = new Integer[3];
    int AcceIndex = 0;

    public void setAccelerometerInitValue(){
        accelerometerInitValue[0] = 0;
        accelerometerInitValue[1] = 0;
        accelerometerInitValue[2] = 0;
    }

    public Integer[] getAccelerometersValue(){
        //get accelerometer value from sensor
        Integer[] sensor = new Integer[3];
        sensor[0] = 0;
        sensor[1] = 0;
        sensor[2] = 0;
        setCord(AcceIndex, sensor);
        return sensor;
    }
    //check for sensor change
    public Boolean checkSensorDisplacement(){
        // there is a change in value
        Boolean change = getAccelerometersValue() != accelerometerInitValue;
        this.setSensorChange(AcceIndex, change);
        return change;
    }
}
