package fyneapp

import "os/exec"

func (app *eflApp) OpenURL(url string) {
	exec.Command("xdg-open", url).Run()
}
