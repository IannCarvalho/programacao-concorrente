public class MyCountDownLatch {
	private int count;

	public MyCountDownLatch(int count) {
		this.count = count;
	}

	public synchronized boolean isEmpty() {
		return this.count < 1;
	}

	public synchronized void countDown() throws InterruptedException {
		this.count = this.count - 1;
		notifyAll();
	}

	public synchronized void await() throws InterruptedException {
		while (!isEmpty()) {
			wait();
		}
	}
}
