package concorrente;

import java.util.ArrayList;

import concorrente.MyThread;

class ThreadsExample {
	public static void main(String[] args) {
		MyThread t1 = new MyThread("1");
		MyThread t2 = new MyThread("2");
		MyThread t3 = new MyThread("3");

		ArrayList<MyThread> array = new ArrayList<MyThread>();
		array.add(t1);
		array.add(t2);
		array.add(t3);
		
		for(MyThread myThread : array) {
			myThread.runThread();
			}
		}
}