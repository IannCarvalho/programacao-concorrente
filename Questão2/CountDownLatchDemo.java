package questao2;

public class CountDownLatchDemo {
	public static void main(String args[]) throws InterruptedException {
		MyCountDownLatch latch = new MyCountDownLatch(10);

		Worker[] array = new Worker[10];
		for (int i = 0; i < 10; i++) {
			Worker da_vez = new Worker(latch, Integer.toString(i));
			array[i] = da_vez;
		}

		for (Worker worker : array) {
			worker.start();
		}

		latch.await();

		System.out.println(Thread.currentThread().getName() + " has finished");
	}
}
