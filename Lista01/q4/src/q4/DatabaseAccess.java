package q4;

import java.util.HashMap;
import java.util.Map;

public class DatabaseAccess {
	
	private String filename;
	
	public DatabaseAccess(String filename) {
		this.filename = filename;
	}
	
	public synchronized void updateDatabase(Map<Integer, Integer> entries) {
		
	}
	
	public synchronized Map<Integer, Integer> getDatabase() {
		return new HashMap<Integer, Integer>();
	}
	
	public synchronized void reset() {
		
	}

	public synchronized Integer getDbValue(Integer key) {
		return null;
	}
	
}
