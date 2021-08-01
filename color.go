package main

type Color string

var (
	Pink      Color = "\033[38;5;198m"
	Red       Color = "\033[38;5;196m"
	Purple    Color = "\033[38;5;135m"
	LightBlue Color = "\033[38;5;45m"
	Blue      Color = "\033[38;5;132m"
	Orange    Color = "\033[38;5;215m"
	Green     Color = "\033[38;5;154m"
	Grey      Color = "\033[38;5;241m"
	DarkGrey  Color = "\033[38;5;239m"
	Black     Color = "\033[38;5;237m"
	White     Color = "\033[38;5;255m"

	Clear Color = "\033[0m"
)
