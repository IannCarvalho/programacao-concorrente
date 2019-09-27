package concorrente;

public class MyThread implements Runnable {
	private String id;
	Thread t;

	public MyThread(String id) {
		this.id = id;
		this.t = new Thread(this, id);
	}

	@Override
	public void run() {
		try {
			for (int i = 0; i < 3; i++) {
				System.out.println("Thread " + id + " executando");
				Thread.sleep(1000);
			}

		} catch (InterruptedException e) {
			System.out.println(id + "interrompida");
		}
		System.out.println("Thread " + id + " terminando");
	}

	public void runThread() {
		t.start();
	}
}