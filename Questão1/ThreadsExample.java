package concorrente;

import java.util.ArrayList;

import concorrente.MyThread;

class ThreadsExample {
	public static void main(String[] args) {
		ArrayList<MyThread> array = new ArrayList<MyThread>();
		LockObject globalLock = new LockObject();
		
		for(int i = 1; i<= 10; i++) {
			array.add(new MyThread(Integer.toString(i), globalLock));
		}

		for (MyThread myThread : array) {
			myThread.runThread();
		}
	}
}