# GoTrack

A program to create graphs for any osu players and their stats based on the osutrack API.

This project uses Fyne as a library for the GUI. 

TODO:
Need to finish Fyne GUI so I can use that to pick which stats I want shown in the charts.
Set up settings.go so that Fyne GUI can be used to select which stats I want shown in the charts.
Make everything in chart.go dynamic so everything can be set from GUI and settings.go
Fix the errors in chart.go


I FIGURED IT OUT
Take the ID range loop out of GetStats. Make the loop its own function so that I can call GetStats and CreateGraph for each user before 