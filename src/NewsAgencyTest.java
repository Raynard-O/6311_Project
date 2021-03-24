import static org.junit.jupiter.api.Assertions.*;

import org.junit.jupiter.api.Test;


public class NewsAgencyTest {

    @Test
    public void whenChangingNewsAgencyState_thenNewsChannelNotified() {

        NewsAgency observable = new NewsAgency();
        NewsChannel observer = new NewsChannel();

        observable.addObserver(observer);

        observable.setNews("news");
        assertEquals(observer.getNews(), "news");
    }
    
}
