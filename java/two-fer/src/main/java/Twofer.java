public class Twofer {

    public String twofer(String name) {
	if(name == null || name.length() == 0) name = "you";
	return "One for " + name + ", one for me.";
    }

}
