import java.util.concurrent.atomic.AtomicInteger;

public class Concurrent4T implements Runnable{
    private static AtomicInteger ai;

        @Override
        public void run() {

            System.out.println("run："+ ai.addAndGet(1));
        }

    public static void main(String[] args) {
        Concurrent4T cr = new Concurrent4T();
        new Thread(cr).start();
        System.out.println("main："+ ai.addAndGet(1));

    }
}
