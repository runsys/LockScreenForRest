package main

// LockScreenForRest project main.go
import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/therecipe/qt/gui"

	"io/ioutil"

	"linbo.ga/toolfunc"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

/*


 */
import "C"

var blockok bool

func procacquire(sema *core.QSystemSemaphore) {
	sema.Acquire()
	blockok = true
}

type TimeDef struct {
	def          string
	lock         string //lock or unlock;
	begin_hour   int
	begin_minute int
	end_hour     int
	end_minute   int
	beginsecs    int
	endsecs      int
}

func main() {

	sym := core.NewQSystemSemaphore("LockScreenForRest", 1, core.QSystemSemaphore__Open)
	go procacquire(sym)
	time.Sleep(1 * time.Second)
	if blockok == false {
		fmt.Println("LockScreenForRest already in running exist.")
		ioutil.WriteFile("ShowOnce", []byte{}, 0666)
		return
	}
	defer sym.Release(1)

	ff, ffe := os.OpenFile("定时强制锁屏休息程序设置.txt", os.O_RDONLY, 0666)
	if ffe != nil {
		ff, ffe = os.OpenFile("定时强制锁屏休息程序设置.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		ff.Write([]byte("#help\n#maxscreen=1\n#fullscreen=1\n#passowrd=\n#hotkey=Ctrl+Shift+Alt+X\n#keyboardallow=1\n#hotkeyallow=1\n#alpha=200\n#lock|unlock hour:minute-hour:minute\nlock 00:00-07:00\nunlock 07:00-07:50\nlock 07:50-08:00\nunlock 08:00-08:50\nlock 08:50-09:00\nunlock 09:00-09:50\nlock 09:50-10:00\nunlock 10:00-10:50\nlock 10:50-11:00\nunlock 11:00-11:50\nlock 11:50-13:00\nunlock 13:00-13:50\nlock 13:50-14:00\nunlock 14:00-14:50\nlock 14:50-15:00\nunlock 15:00-15:50\nlock 15:50-16:00\nunlock 16:00-16:50\nlock 16:50-17:00\nunlock 17:00-17:50\nlock 17:50-18:00\nunlock 18:00-18:50\nlock 18:50-19:00\nunlock 19:00-19:50\nlock 19:50-21:00\nunlock 21:00-21:50\nlock 21:50-22:00\nunlock 22:00-22:50\nlock 22:50-23:00\nunlock 23:00-23:50\nlock 23:50-24:00\n\n\n\n\n"))
		ff.Close()
	} else {
		ff.Close()
	}
	ffctt, _ := ioutil.ReadFile("定时强制锁屏休息程序设置.txt")
	fflis := strings.Split(string(ffctt), "\n")
	tdefs := []*TimeDef{}
	var xxminute string
	var alwaysshow = 0
	var password string
	var fullscreen = 0
	var maxscreen = 1
	var hotkey string
	var keyboardallow = 1
	var alpha int = 200
	var hotkeyallow int = 1
	var lockuntiluserclose bool
	var unlockuntilnextunlock bool
	for _, lin := range fflis {
		lin = strings.Trim(lin, " \r\\n\t")
		if len(lin) > 0 {
			if len(lin) > 0 && lin[0] != '#' && strings.Index(lin, "lock ") != -1 && strings.Index(lin, "-") != -1 && strings.Index(lin, ":") != -1 && (len(lin) == 16 || len(lin) == 18) {
				//lin format:lock hour:minute-hour:minute
				td := &TimeDef{}
				td.def = lin[strings.Index(lin, " ")+1:]
				td.lock = lin[:strings.Index(lin, " ")]
				sp := strings.Index(lin, " ")
				slashp := strings.Index(lin, "-")
				slashppre := lin[sp+1 : slashp]
				slashpaft := lin[slashp+1:]
				slashpprep := strings.Index(slashppre, ":")
				slashpaftp := strings.Index(slashpaft, ":")
				td.begin_hour = toolfunc.Atoi(slashppre[:slashpprep])
				td.begin_minute = toolfunc.Atoi(slashppre[slashpprep+1:])
				td.beginsecs = td.begin_hour*3600 + td.begin_minute*60
				td.end_hour = toolfunc.Atoi(slashpaft[0:slashpaftp])
				td.end_minute = toolfunc.Atoi(slashpaft[slashpaftp+1:])
				td.endsecs = td.end_hour*3600 + td.end_minute*60
				tdefs = append(tdefs, td)
			} else if strings.Index(lin, "=") != -1 && lin[strings.Index(lin, "=")+1:] == "#hotkey=" {
				hotkey = lin[strings.Index(lin, "=")+1:]
			} else if strings.Index(lin, "=") != -1 && lin[strings.Index(lin, "=")+1:] == "#keyboardallow=" {
				keyboardallow = toolfunc.Atoi(lin[strings.Index(lin, "=")+1:])
			} else if strings.Index(lin, "=") != -1 && lin[strings.Index(lin, "=")+1:] == "#alpha=" {
				alpha = toolfunc.Atoi(lin[strings.Index(lin, "=")+1:])
			} else if strings.Index(lin, "=") != -1 && lin[strings.Index(lin, "=")+1:] == "#hotkeyallow=" {
				hotkeyallow = toolfunc.Atoi(lin[strings.Index(lin, "=")+1:])
			} else if strings.Index(lin, "=") != -1 && lin[strings.Index(lin, "=")+1:] == "#password=" {
				password = lin[strings.Index(lin, "=")+1:]
			} else if strings.Index(lin, "=") != -1 && lin[strings.Index(lin, "=")+1:] == "#fullscreen=" {
				fullscreen = toolfunc.Atoi(lin[strings.Index(lin, "=")+1:])
			} else if strings.Index(lin, "=") != -1 && lin[strings.Index(lin, "=")+1:] == "#maxscreen=" {
				maxscreen = toolfunc.Atoi(lin[strings.Index(lin, "=")+1:])
			}
		}
	}
	fmt.Println(
		"xxminute", xxminute,
		"alwaysshow", alwaysshow,
		"password", password,
		"fullscreen", fullscreen,
		"maxscreen", maxscreen,
		"hotkey", hotkey,
		"keyboardallow", keyboardallow,
		"alpha", alpha,
		"hotkeyallow", hotkeyallow,
	)

	widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("LockScreenForRest")
	window.SetMinimumSize2(200, 200)

	layout := widgets.NewQVBoxLayout()
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)

	restlabel := widgets.NewQLabel2("Take a rest", nil, 0)
	restlabel.SetAlignment(core.Qt__AlignHCenter)
	restlabel.SetStyleSheet("font-size:100pt")
	layout.AddWidget(restlabel, 0, 0)
	go func() {
		for true {
			time.Sleep(5 * time.Second)
			if len(tdefs) > 0 {
				now := time.Now()
				cursecs := now.Hour()*3600 + now.Minute()*60
				for i := 0; i < len(tdefs); i += 1 {
					if cursecs >= tdefs[i].beginsecs && cursecs < tdefs[i].endsecs {
						if tdefs[i].lock == "lock" {
							if toolfunc.FileExists("ShowOnce") {
								os.Remove("ShowOnce")
								lockuntiluserclose = true
							}
							if lockuntiluserclose {
								window.ShowMaximized()
								window.ActivateWindow()
								restlabel.SetText("Rest Time " + tdefs[i].def)
								continue
							}
							if unlockuntilnextunlock {
								window.Hide()
								restlabel.SetText("!Rest Time " + tdefs[i].def)
								continue
							} else {
								window.ShowMaximized()
								window.ActivateWindow()
								restlabel.SetText("Rest Time " + tdefs[i].def)
								continue
							}
						} else if tdefs[i].lock == "unlock" {
							unlockuntilnextunlock = false
							if toolfunc.FileExists("ShowOnce") || lockuntiluserclose {
								os.Remove("ShowOnce")
								lockuntiluserclose = true
								window.ShowMaximized()
								window.ActivateWindow()
								restlabel.SetText("！Working Time " + tdefs[i].def)
								continue
							} else {
								window.Hide()
								restlabel.SetText("Working Time " + tdefs[i].def)
								continue
							}
						}
					}
				}
			}

		}
	}()

	window.ConnectCloseEvent(func(evt *gui.QCloseEvent) {
		if password != "" {
			idlg := widgets.NewQInputDialog(nil, core.Qt__Window)
			idlg.SetLabelText("Password")
			ipass := idlg.GetText(nil, "Input Close Password", "Password:", widgets.QLineEdit__Password, "", nil, 0, 0)
			if ipass != password {
				evt.Ignore()
				return
			}
		}
		lockuntiluserclose = false
		unlockuntilnextunlock = true
		window.Hide()
		evt.Ignore()
	})

	window.SetCentralWidget(widget)
	window.ShowMaximized()
	widgets.QApplication_Exec()
}
