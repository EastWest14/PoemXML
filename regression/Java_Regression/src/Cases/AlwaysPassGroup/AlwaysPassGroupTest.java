package Cases.AlwaysPassGroup;

import static org.junit.Assert.*;
import Cases.*;

import org.junit.Test;

public class AlwaysPassGroupTest {

	@Test
	public void test() {
		AlwaysPassGroup apGroup = new AlwaysPassGroup();
		GroupRunResult result = apGroup.execute();
		
		assertEquals(result.passes(), true);
	}

}
