import org.junit.Test;

import java.security.PublicKey;

import static org.junit.Assert.assertEquals;

public class BikeTest {

    @Test
    public void WhenApplicationChangesState_thenNotifyUser(){
        Application observable = new Application();
        User        observer = new Bike();

        observable.addUser(observer);
        observable.setNotification("Change In State");
        assertEquals(observer.getMessage(), "news");
    }

    @Test
    public void WhenApplicationChangesState_thenNotifyUser(){
        Application observable = new Application();
        User        observer = new Bike();

        observable.addUser(observer);
        observable.setNotification("Change In State");
        assertEquals(observer.getMessage(), "news");
    }

}

public class BikeTest {

    @Test
    public void WhenApplicationChangesState_thenNotifyUser(){
        Application observable = new Application();
        User        observer = new Bike();

        observable.addUser(observer);
        observable.setNotification("Change In State");
        assertEquals(observer.getMessage(), "news");
    }

}



