package questao1;

import java.util.ArrayList;
import java.util.concurrent.locks.LockSupport;

public class Fifo implements Runnable {
	private final ArrayList<Thread> waiters = new ArrayList<Thread>();

	public void lock() {
		synchronized (this) {
			waiters.add(Thread.currentThread());
		}
		if (peek() != Thread.currentThread())
			LockSupport.park(this);
	}

	// Verificacao dupla se o arraylist está vazio eh necessaria porque o elemento
	// remove antes de fazer o unpark. Pode existir uma situacao em que ele tenta
	// fazer um unpark em um elemento que nao existe
	public void unlock() {
		synchronized (this) {
			if (isNotEmpty())
				waiters.remove(0);
			if (isNotEmpty())
				LockSupport.unpark(peek());
		}
	}

	public Thread peek() {
		synchronized (this) {
			return waiters.get(0);
		}
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