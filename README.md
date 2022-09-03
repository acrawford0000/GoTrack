# GoTrack

A program to create graphs for any osu players and their stats based on the osutrack API.

This project uses Fyne as a library for the GUI. Go-echarts for the graphs.

TODO:
Figure out how to get data for 1 user at a time to the chart without overwriting any previous
Need to finish Fyne GUI so that stats can be filtered to/from chart based on menu items (kinda sort of)
Set up settings.go so that Fyne GUI can be used to select which stats I want shown in the charts (dont know how to filter the struct fields when I go to chart them)
Make everything in chart.go dynamic so everything can be set from GUI and settings.go (cant make dynamic settings if I dont know how to even filter)
Fix the errors in chart.go (Done)

QUESTIONS:
Does modelstats need to be slices? The API response looks separated but how do you get the hraph to iterate over the entire response? Is it better to put them all together into one?