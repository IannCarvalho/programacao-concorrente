import java.util.ArrayList;
import java.util.concurrent.locks.LockSupport;

public class Questao01 implements Runnable {
	private final ArrayList<Thread> waiters = new ArrayList<Thread>();

	public void lock() {
		boolean wasInterrupted = false;
		Thread current = Thread.currentThread();
		
		synchronized (this) {
			waiters.add(current);
		}

		while (peek() != current) {
			LockSupport.park(this);
			if (Thread.interrupted())
				wasInterrupted = true;
		}

		if (wasInterrupted)
			current.interrupt();

	}

	public void unlock() {

		synchronized (this) {
			if (isNotEmpty())
				waiters.remove(0);
		}

		if (isNotEmpty())
			LockSupport.unpark(peek());

	}

	public Thread peek() {
		return waiters.get(0);
	}

	public boolean isNotEmpty() {
		return waiters.size() > 0;
	}

	@Override
	public void run() {
		lock();
		Thread current = Thread.currentThread();
		try {
			for (int i = 0; i < 3; i++) {
				System.out.println("Thread " + current.getId() + " executando");
				Thread.sleep(1);
			}
		} catch (InterruptedException e) {
			System.out.println(current.getId() + "interrompida");
		}
		System.out.println("Thread " + current.getId() + " terminando");
		unlock();
	}

	public void runThread() {
		Thread t = new Thread(this);
		t.start();
	}
}
