public class Bike implements User{
    private String message;
    private Integer id;
    private Boolean security;
    private Integer[] coordinates;

    @Override
    public void notify(Object message){
        setMessage((String) message);
    }
    @Override
    public void setId(Object id) {
        this.id = (Integer) id;
    }
    // get user id
    public Integer getId() {
        return id;
    }
    public String getMessage() {
        return message;
    }
    public void setMessage(String message){
        this.message = message;
    }

    public void setSecurity(Boolean secure){
        this.security = secure;
    }

    public void intervention(){
        if (message != null) {
            setMessage(null);
            Cancel();
        }
    }

    public void Cancel(){
        setSecurity(false);
    }

    public void setCoordinates(int[] coordinates) {
        this.coordinates = this.coordinates;
    }

    public Integer[] getCoordinates() {
        return coordinates;
    }

    public void Track(Object o){
        setCoordinates((int[]) o);
    }

}
