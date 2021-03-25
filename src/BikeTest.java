import org.junit.Test;

import java.security.PublicKey;

import static org.junit.Assert.assertEquals;

public class BikeTest {

    @Test
    public void WhenApplicationChangesState_thenNotifyUser(){
        Microcontroller observable = new Microcontroller();
        User        observer = new Bike();

        observable.addUser(observer);
        //check for changes
        Boolean changes = observable.checkDisplacement();
        if (changes = true) {
            observable.setNotification("change in state");
            assertEquals(observer.getMessage(), "change in state");
            assertEquals(observer.getSecurity(), false);
        }
        Boolean intervention = observer.intervention();
        if (intervention = true) {
            // assert that system is still secure if intervention occurs
            assertEquals(observer.getSecurity(), true);
            assertEquals(observer.getMessage(), null);
        }else {
            // track device
            while (observer.intervention()) {
                observer.Track(observer.getCoordinates());
            }

        }


    }

//    @Test
//    public void WhenApplicationChangesState_thenNotifyUser(){
//        Application observable = new Application();
//        User        observer = new Bike();
//
//        observable.addUser(observer);
//        observable.setNotification("Change In State");
//        assertEquals(observer., "news");
//    }

}





