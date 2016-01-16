package muls

type Server interface {
	Run() bool;
	Stop() bool;
}