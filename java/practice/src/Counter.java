import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

/**
 * Restrictions:
 * 1. Do not change data type of "count" and its initial value.
 * 2. Do not change "main" method.
 * 3. Do not change "loopTill" value
 * <p>
 * Problem:
 * 1. Every time you run the program, "Actual Count" output varies.
 * <p>
 * Success Criteria:
 * 1. "Actual Count" should always match to "Expected Count".
 */
public class Counter implements Runnable {
	private static int count = 0;
	private static final int loopTill = 1000;

	@Override
	public void run() {
		synchronized (this) {
			for (int i = 0; i < loopTill; i++) {
				count++;
			}
		}
	}

	public static void main(String[] args) {
		int noOfThreads = 10;
		ExecutorService executorService = Executors.newFixedThreadPool(noOfThreads);
		for (int i = 0; i < noOfThreads; i++) {
			executorService.execute(new Counter());
		}
		executorService.shutdown();
		System.out.println("Expected count: " + (noOfThreads * loopTill));
		System.out.println("Actual count: " + count);
	}
}