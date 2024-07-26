package pkg

import (
	"Novice/parts"
	"fmt"
	"os/exec"
	"strings"
)

func Comparison() {
	parts.TimeHMS()
	println(" Comparing...")
	packageNames := []string{"wqy-microhei", "wqy-microhei-lite", "wqy-zenhei", "ttf-arphic-ukai", "ttf-arphic-uming", "adobe-source-han-sans-cn-fonts", "adobe-source-han-serif-cn-fonts", "noto-fonts-cjk", "yay", "noto-fonts-cjk", "noto-fonts-emoji", "wget", "paru", "jdk21-openjdk", "vi", "nano", "go", "wechat-uos-qt", "linuxqq-appimage", "Dingtalk-bin", "nautilus", "hyprctl", "tofictl", "pseudo", "grim", "OBS-studio", "obs-cli", "togglefloating", "swayidle", "jdk17-openjdk", "dbus", "git", "tofi", "vlc", "alacritty", "fish", "hyprland", "fontconfig", "hyprlock", "hyprpaper", "kitty", "mako", "waybar", "dwindle", "fcitx5-im", "fcitx5-chinese-addons", "fcitx5-rime", "fcitx5-qt"}
	comparePackages(packageNames)
}

func comparePackages(packageList []string) {
	installedPackages := make(map[string]bool)
	var notInstalledPackages []string
	out, err := exec.Command("pacman", "-Q").Output()
	if err != nil {
		parts.Warning()
		fmt.Println("ERRORED when getting pacman pkg-list:", err)
		return
	}
	pacmanOutput := strings.Split(string(out), "\n")
	for _, pkgLine := range pacmanOutput {
		pkgName := strings.Split(pkgLine, " ")[0]
		pkgName = strings.ToLower(pkgName)
		for _, targetPkg := range packageList {
			targetPkg = strings.ToLower(targetPkg)
			if strings.Contains(pkgName, targetPkg) {
				installedPackages[targetPkg] = true
				break
			}
		}
	}
	for _, pkg := range packageList {
		pkg = strings.ToLower(pkg)
		if installedPackages[pkg] {
			continue
		}
		notInstalledPackages = append(notInstalledPackages, pkg)
	}
	if len(installedPackages) > 0 {
		parts.Done()
		fmt.Println("HAD INSTALLED :: ", strings.TrimSuffix(strings.Join(getKeys(installedPackages), " "), ", "))
	}
	if len(notInstalledPackages) > 0 {
		parts.Remind()
		fmt.Println("NOT INSTALLED :: ", strings.Join(notInstalledPackages, " "))
	}
}

func getKeys(m map[string]bool) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}
