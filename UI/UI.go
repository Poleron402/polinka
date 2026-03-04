package ui

import "fmt"

func AppUI() {
	color1 := "\x1b[34m"
	color2 := "\x1b[31m"
	color3 := "\x1b[32m"
	color4 := "\x1b[33m"
	color5 := "\x1b[35m"
	color6 := "\x1b[94m"
	color7 := "\x1b[96m"

	resetColor := "\x1b[0m"
	name := fmt.Sprintf(`%s      ___   %s%s      ___      %s%s                               ___     %s%s      ___       %s%s    ___    %s 
%s     /  /\   %s%s    /  /\    %s%s               %s%s   ___      %s%s    /__/\    %s%s     /__/|     %s%s    /  /\    %s
%s    /  /::\  %s%s   /  /::\   %s%s               %s%s  /  /\     %s%s    \  \:\   %s%s    |  |:|     %s%s   /  /::\   %s
%s   /  /:/\:\ %s%s  /  /:/\:\  %s%s  ___     ___  %s%s /  /:/     %s%s     \  \:\   %s%s   |  |:|     %s%s  /  /:/\:\  %s
%s  /  /:/~/:/ %s%s /  /:/  \:\ %s%s /__/\   /  /\ %s%s/__/::\     %s%s _____\__\:\ %s%s  __|  |:|     %s%s /  /:/~/::\ %s
%s /__/:/ /:/  %s%s/__/:/ \__\:\ %s%s\  \:\ /  /:/ %s%s\__\/\:\__  %s%s/__/::::::::\ %s%s/__/\_|:|____ %s%s/__/:/ /:/\:\%s
%s \  \:\/:/  %s%s \  \:\ /  /:/ %s%s \  \:\  /:/  %s%s   \  \:\/\ %s%s\  \:\~~\~~\/ %s%s\  \:\/:::::/ %s%s\  \:\/:/__\/%s
%s  \  \::/   %s%s  \  \:\  /:/  %s%s  \  \:\/:/   %s%s    \__\::/ %s%s \  \:\  ~~~  %s%s \  \::/~~~~  %s%s \  \::/     %s
%s   \  \:\  %s%s    \  \:\/:/   %s%s   \  \::/    %s%s    /__/:/  %s%s  \  \:\      %s%s  \  \:\      %s%s  \  \:\     %s
%s    \  \:\ %s%s     \  \::/    %s%s    \__\/     %s%s    \__\/   %s%s   \  \:\     %s%s   \  \:\     %s%s   \  \:\  %s  
%s     \__\/ %s%s      \__\/     %s%s              %s%s            %s%s    \__\/     %s%s    \__\/     %s%s    \__\/  %s`, 
color1, resetColor, color2, resetColor, color5, resetColor, color6, resetColor, color7, resetColor,
color1, resetColor, color2, resetColor, color3, resetColor, color4, resetColor, color5, resetColor, color6, resetColor, color7, resetColor,
color1, resetColor, color2, resetColor, color3, resetColor, color4, resetColor, color5, resetColor, color6, resetColor, color7, resetColor,
color1, resetColor, color2, resetColor, color3, resetColor, color4, resetColor, color5, resetColor, color6, resetColor, color7, resetColor,
color1, resetColor, color2, resetColor, color3, resetColor, color4, resetColor, color5, resetColor, color6, resetColor, color7, resetColor,
color1, resetColor, color2, resetColor, color3, resetColor, color4, resetColor, color5, resetColor, color6, resetColor, color7, resetColor,
color1, resetColor, color2, resetColor, color3, resetColor, color4, resetColor, color5, resetColor, color6, resetColor, color7, resetColor,
color1, resetColor, color2, resetColor, color3, resetColor, color4, resetColor, color5, resetColor, color6, resetColor, color7, resetColor,
color1, resetColor, color2, resetColor, color3, resetColor, color4, resetColor, color5, resetColor, color6, resetColor, color7, resetColor,
color1, resetColor, color2, resetColor, color3, resetColor, color4, resetColor, color5, resetColor, color6, resetColor, color7, resetColor,
color1, resetColor, color2, resetColor, color3, resetColor, color4, resetColor, color5, resetColor, color6, resetColor, color7, resetColor,)


	fmt.Println(name)
}