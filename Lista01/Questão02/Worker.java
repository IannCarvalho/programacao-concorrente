class Worker extends Thread {
	private MyCountDownLatch latch;

	public Worker(MyCountDownLatch latch, String name) {
		super(name);
		this.latch = latch;
	}

	@Override
	public void run() {
		Thread current = Thread.currentThread();
		try {
			for (int i = 0; i < 3; i++) {
				System.out.println("Thread " + current.getName() + " executando");
				Thread.sleep(1);
			}
			System.out.println("Thread " + current.getName() + " terminando");
			latch.countDown();
		} catch (InterruptedException e) {
			System.out.println(current.getName() + "interrompida");
		}
	}
}