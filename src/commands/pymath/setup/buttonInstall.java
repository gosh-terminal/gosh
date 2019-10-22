import javax.swing.*;
import java.awt.*;
class buttonInstall {
    public static void main(String[] args) {
        JFrame frame = new JFrame("pymath installer");
        JLabel label = new JLabel("Click here to install");
        JButton install = new JButton("Install");
        JButton cancel = new JButton("Cancel");
        frame.setVisible(true);
    }
}