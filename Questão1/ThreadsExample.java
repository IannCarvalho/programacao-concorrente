package concorrente;

import concorrente.FIFOMutex;


public class ThreadsExample{
    public static void main(String[] args) {
    	FIFOMutex fifo = new FIFOMutex();
    	fifo.runThread();
    	fifo.runThread();
    	fifo.runThread();	
    }
}