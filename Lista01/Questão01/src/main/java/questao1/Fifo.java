package questao1;

import java.util.ArrayList;
import java.util.concurrent.locks.LockSupport;

public class Fifo implements Runnable {
	private final ArrayList<Thread> waiters = new ArrayList<Thread>();

	public void lock() {
		synchronized (this) {
			waiters.add(Thread.currentThread());
		}
		while (peek() != Thread.currentThread()) {
			LockSupport.park(this);
		}

	}

	public void unlock() {
		synchronized (waiters) {
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
		System.out.println("Thread " + current.getName() + " terminando");
		unlock();
	}

	public void runThread(String name) {
		Thread t = new Thread(this);
		t.setName(name);
		t.start();
	}
}