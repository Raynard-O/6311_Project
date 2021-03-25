public class Bike implements User{
    private String message;
    private Integer id;
    private Boolean security;
    private Integer[][] coordinates;

    @Override
    public void notify(Object message){
        setMessage((String) message);
        setSecurity(false);
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

    public Boolean getSecurity() {
        return security;
    }

    public Boolean intervention(){
        if (message != null) {
            setMessage(null);
            Cancel();
        }
        return true;
    }

    public void Cancel(){
        setSecurity(false);
    }
    @Override
    public void setCoordinates(Integer[][] coordinates) {
        this.coordinates = coordinates;
    }

    public Integer[][] getCoordinates() {
        return coordinates;
    }

    public void Track(Object o){
        setCoordinates((Integer[][]) o);
    }

}
