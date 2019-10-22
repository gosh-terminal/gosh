import terminal
import osproc
setForeGroundColor(fgCyan)
let dir = execProcess("pwd")
stdout.write(dir)
stdout.write("gosh> ")
setForeGroundColor(fgDefault)