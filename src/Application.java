import java.util.ArrayList;
import java.util.List;
import java.util.Random;

public class Application {
    private String notification;
    private Integer Ax;
    private Integer Ay;
    private Integer Az;
    private Integer Gx;
    private Integer Gy;
    private Integer Gz;
    private int[] coordinates;


    private List<User> users = new ArrayList<>();
    int upperbound = users.size();

    // adding a new user to system
    public void addUser(User user) {
        Random rand  = new Random();
        user.setId(upperbound);
        this.users.add(user);
    }

    // remove a new user to system
    public void removeUser (User user) {
        this.users.remove(user);
    }

    public void systemShutdown(){
        this.users.removeAll(users);
    }

    public void setNotification(String message) {
        this.notification = message;
        for (User user : this.users) {
            user.notify(this.notification);
        }
    }
    public int[] getCoordinate(){
        int[] accelerometer = new int[3];
        int[] gyroscope = new int[3];
        System.arraycopy(accelerometer,0, coordinates, 0, accelerometer.length);
        System.arraycopy(gyroscope,0, coordinates, accelerometer.length, gyroscope.length);
        return coordinates;
    }

    public void sendCoordinates(User user){
        // if no intervention
        user.Track(coordinates);
    }
}

