package concorrente;

public class LockObject {
	boolean locked;

	public LockObject() {
		this.locked = false;
	}
	
	public boolean TAS() {
		boolean initial = locked;
		locked = true;
		return initial;
	}
	
	public void exitCritial() {
		locked = false;
	}
	
	public boolean getLock() {
		return locked;
	}

}
