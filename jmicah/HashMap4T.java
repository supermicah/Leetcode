import java.util.HashMap;
import java.util.concurrent.ConcurrentHashMap;

public class HashMap4T {
    HashMap<String,String> hm = new HashMap<>();
    ConcurrentHashMap<String, String>  chm = new ConcurrentHashMap<>();


    public synchronized String get(){
        return "";
    }
}
