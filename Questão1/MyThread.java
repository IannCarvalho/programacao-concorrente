package concorrente;

public class MyThread implements Runnable {
	private String id;
	Thread t;
	private LockObject lock;

	public MyThread(String id, LockObject lock) {
		this.id = id;
		this.t = new Thread(this, id);
		this.lock = lock;
	}

	@Override
	public void run() {
		while (!lock.TAS()) {
			try {
				for (int i = 0; i < 3; i++) {
					System.out.println("Thread " + id + " executando");
					Thread.sleep(1000);
				}

			} catch (InterruptedException e) {
				System.out.println(id + "interrompida");
			}
			System.out.println("Thread " + id + " terminando");
			lock.exitCritial();
			break;
		}
	}

	public void runThread() {
		t.start();
	}
}