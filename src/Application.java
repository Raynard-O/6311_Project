import java.util.ArrayList;
import java.util.List;
import java.util.Random;

public class Application {
    private String notification;
    private Integer[][] coordinates;


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
    public Integer[][] getCoordinate(){

//        System.arraycopy(accelerometer,0, coordinates, 0, accelerometer.length);
//        System.arraycopy(gyroscope,0, coordinates, accelerometer.length, gyroscope.length);
        return coordinates;
    }

    public void setCoordinates(Integer[][] cord) {
        this.coordinates = cord;
    }

    public void sendCoordinates(User user){
        // if no intervention
        user.Track(coordinates);
    }




}

