package concorrente;

import java.util.ArrayList;
import java.util.concurrent.locks.LockSupport;

public class FIFOMutex implements Runnable {
	private boolean locked;
	private final ArrayList<Thread> waiters = new ArrayList<Thread>();

	public FIFOMutex() {
		this.locked = false;
	}

	public void lock() {
		boolean wasInterrupted = false;
		synchronized (this) {
			Thread current = Thread.currentThread();
			waiters.add(current);
			
			while (peek() != current) {
				LockSupport.park(this);
				if (Thread.interrupted())
					wasInterrupted = true;
			}
			
			remove();			
			
			if (wasInterrupted)
				current.interrupt();
		}
	}

	public void unlock() {
		locked = false;
		if (!isEmpty())
			LockSupport.unpark(peek());
	}

	public Thread peek() {
		return waiters.get(0);
	}

	public boolean isEmpty() {
		return waiters.size() > 0;
	}

	public void remove() {
		if (!isEmpty())
			waiters.remove(0);
	}

	public boolean compareAndSet() {
		boolean initial = locked;
		locked = true;
		return initial;
	}

	@Override
	public void run() {
		lock();
		Thread current = Thread.currentThread();
		try {
			for (int i = 0; i < 3; i++) {
				System.out.println("Thread " + current.getId() + " executando");
				Thread.sleep(1000);
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
