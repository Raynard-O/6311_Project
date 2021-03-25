public class gyroscope extends Microcontroller {
    private final Integer[] gyroscopeInitValue = new Integer[3];

    int GyscIndex = 1;
    public void setGyroscopeInitValue() {
        gyroscopeInitValue[0] = 0;
        gyroscopeInitValue[1] = 0;
        gyroscopeInitValue[2] = 0;
    }

    public Integer[] getGyroscopesValue() {
        //get accelerometer value from sensor
        Integer[] sensor = new Integer[3];
        sensor[0] = 0;
        sensor[1] = 0;
        sensor[2] = 0;
        setCord(GyscIndex, sensor);
        return sensor;
    }

    //check for sensor change
    public Boolean checkSensorDisplacement() {
        // there is a change in value
        Boolean change = getGyroscopesValue() != gyroscopeInitValue;
        // there is a change in value
        this.setSensorChange(GyscIndex, change);
        return change;
    }
}

